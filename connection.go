package esl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/textproto"
	"strconv"
	"strings"
)

// error types
var (
	ErrParsingFailed = errors.New("parse the header failed")
)

// simple wrapper of the net.Conn
type connection struct {
	net.Conn
	*textproto.Reader
}

// send plain text command over net.Conn
// simple write to socket
func (c *connection) send(command string) (err error) {
	wait := []byte(command)
	n := 0
	s := len(wait)
	for n < s {
		wait = wait[n:] // next slice to send
		n, err = c.Conn.Write(wait)
		if err != nil {
			warnf("I/O error when write to socket %v", err)
			return
		}
	}
	return
}

// recv a new event object
// using the textproto to parse the protocol of FS
func (c *connection) recv() (event *Event, err error) {
	headers, err := c.ReadMIMEHeader()
	if err != nil && err != io.EOF {
		errorf("parse event failed %v", err)
		err = ErrParsingFailed
		return
	}
	event = new(Event)
	event.Body = ""
	event.Headers = headers

	var body []byte
	if contentLength := headers.Get(EslContentLength); contentLength != "" {
		length, _ := strconv.Atoi(contentLength) // ignore the error please check the logic your own module issued
		body = make([]byte, length)
		if _, err = io.ReadFull(c.Reader.R, body); err != nil {
			errorf("parse body failed %v", err)
			return
		}
	}

	// switch the content type to handle the body
	redo:
	if contentType := headers.Get(EslContentType); contentType != "" {
		switch contentType {
		case EslEventContentJson:
			{
				// handle json
				err = json.Unmarshal(body, &event.Headers)
				if err != nil {
					err = ErrJsonBodyParsing
					return
				}
				//FS JSON format the event body is emplaced in the fuck _body field name ? what's Fuck Design these idiots done :)
				if length := len(event.Headers["_body"]); length > 0 {
					event.Body = event.Headers["_body"][0]
				}
				delete(event.Headers, "_body")
				event.Type = EslEvent
			}
		case EslEventContentPlain:
			{
				// handle plain
				// in plain the body is also Header: value pattern
				// TODO: should check the textproto splitor is \r\n but the body is \n
				r := textproto.NewReader(
					bufio.NewReader(
						bytes.NewReader(body)))
				content, e := r.ReadMIMEHeader()
				if e != nil && e != io.EOF {
					errorf("parse event failed %v", e)
					err = ErrParsingFailed
					return
				}
				if contentLength := content.Get(EslContentLength); contentLength != "" {
					length, _ := strconv.Atoi(contentLength)
					body = make([]byte, length)
					if _, e := io.ReadFull(c.Reader.R, body); e != nil {
						errorf("parse body failed %v", e)
						err = e
						return
					}
					event.Body = string(body)
				}
				event.Headers = content
				event.Type = EslEvent
				// handle the ARRAY:: situation have to care about trim to array
				// WTF design they have done LAMO
				for ok, ov := range event.Headers {
					for _, v := range ov {
						if strings.Contains(v, "ARRAY::") {
							rv := strings.Split(v, "ARRAY::")
							event.Headers[ok] = rv
						}
					}
				}
			}
		case EslEventContentCommandReplay:
			{
				// command reply
				event.Type = EslReply
				event.Body = headers.Get(EslReplyText)
			}
		case EslEventContentApiResponse:
			{
				// api response
				event.Type = EslResponse
				event.Body = string(body)
			}
		case EslEventContentDiscoonectNotice:
			{
				// disconnect-notice
				event.Type = EslDisconnectedNotice
				debug("receive the disconnected notice")
				uuid := headers.Get("Controlled-Session-UUID")
				// FIXME: maybe there need an error to identify this event
				debugf("Session UUID socket Disconnected Notice %s", uuid)
			}
		case EslEventAuthRequest: goto redo // because the auth request will be none should ignore
		}
	}

	return
}
