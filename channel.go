package esl

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

// Channel should be thread safe
// and every client could use this Channel abstraction in multi-tenant system building
// if the connection need the queue ?

// sendmsg
// sendevent
// connect
// auth ClueCon
// userauth username ClueCon

// Execute an application sync method
// e.g.
//
// api status
//
// Content-Type: api/response
// Content-Length: 367
//
// UP 0 years, 0 days, 16 hours, 51 minutes, 5 seconds, 534 milliseconds, 583 microseconds
// FreeSWITCH (Version 1.10.7-dev git 81fff85 2021-06-14 16:46:28Z 64bit) is ready
// 16 session(s) since startup
// 0 session(s) - peak 2, last 5min 0
// 0 session(s) per Sec out of max 30, peak 1, last 5min 0
// 1000 session(s) max
// min idle cpu 0.00/99.63
// Current Stack Size/Max 240K/8192K

// when the channel closed will callback this function
type closeFunc = func()

// inner channel
type channel struct {
	connection               // inner connection body
	sync.Mutex               // sync mutex when execute the command
	reply      chan *Event   // reply channel
	response   chan *Event   // response channel
	Events     chan *Event   // events channel
	signal     chan struct{} // signal useful for outbound pattern
	close      closeFunc     // inner connection close callback
	err        error
	running    atomic.Value
}

// loop for recv
func (c *channel) loop() {
	c.running.Store(true)
	for {
		running, _ := c.running.Load().(bool)
		if !running {
			break
		}
		if c.err != nil {
			debug("channel meet a fuck IO error")
			break
		}
		if event, err := c.recv(); err != nil {
			debugf("recv error %v", err)
			break
		} else {
			if event.Type == EslEvent {
				c.Events <- event
			}
			if event.Type == EslReply {
				c.reply <- event
			}
			if event.Type == EslResponse {
				c.response <- event
			} else {
				m, _ := event.IntoPlain()
				errorf("unsupport content type %v, and the result is %s", event.Type, m)
				break
			}
		}

	}
	c.shutdown()
	c.clear()
}

func (c *channel) clear() {
	if c.signal != nil {
		close(c.signal)
	}
	if c.close != nil {
		c.close()
	}
}

// shutdown the channel
func (c *channel) shutdown() {
	c.connection.Close()
	c.clear()
}

// execute command sync
func (channel *channel) command(cmd string) (err error) {
	cmd = fmt.Sprintf("%s\r\n\r\n", cmd)
	channel.Lock()
	err = channel.send(cmd)
	if err != nil {
		errorf("channel execute command failed %v", err)
		channel.err = err
		return
	}
	reply := <-channel.reply // reply could change the SC FUCK!
	channel.Unlock()
	if strings.Contains(reply.Body, "-ERR") {
		debugf("channel send command failed %v", err)
		return errors.New(reply.Body)
	}
	return
}

// execute unload a module mod_event_socket
func (channel *channel) unload() (err error) {
	return channel.command("api bgapi unload mod_event_socket")
}

// execute reload a module mod_event_socket
func (channel *channel) reload() (err error) {
	return channel.command("api bgapi reload mod_event_socket")
}

// execute the filter command
func (channel *channel) filter(action string, events ...string) (err error) {
	// filter delete commands...
	// filter add commands...
	es := strings.Join(events, " ")
	return channel.command(fmt.Sprintf("filter %s %s", action, es))
}

// execute the resume command
func (channel *channel) resume() (err error) {
	// in FS could set this session as LFLAG_RESUME
	// maybe useful for Inbound mode
	return channel.command("resume")
}

// TODO: except these methods other methods should check the auth logic

// execute auth command
func (channel *channel) auth(password string) (err error) {
	return channel.command(fmt.Sprintf("auth %s", password))
}

// execute userauth command
func (channel *channel) userauth(username, password string) (err error) {
	return channel.command(fmt.Sprintf("userauth %s %s", username, password))
}

// execute the sendmsg logic with application
func (channel *channel) execute() (response *Event, err error) {
	return
}

// TODO: here need a batch pattern to execute for one connection
// TODO: here also need abtach pattern for execute method

// bellow methods should be authorizated or outbound method

// execute connect command
func (channel *channel) connect() (event *Event, err error) {
	// connect with connect event
	// should check the result should return result event response
	return //channel.command("connect")
}

// execute answer command
func (channel *channel) answer() (err error) {
	// this should use ececute command
	return
}

// execute linger command
func (channel *channel) linger() (err error) {

	return
}

// execute nolinger command
func (channel *channel) nolinger() (err error) {

	return
}

// execute the getvar command to get the variable of current channel
func (channel *channel) getvar(key string) (err error) {

	return
}

// execute sendevent command
func (channel *channel) sendevent() (err error) {

	return
}

// execute command sendmsg
func (channel *channel) sendmsg() (response *Event, err error) {
	return
}

// execute api command sync
func (channel *channel) api() (err error) {

	return
}

// execute bgapi command async
func (channel *channel) bgapi() (err error) {

	return
}

// execute event command to subcribe the events specificed
func (channel *channel) event(format string, events ...string) (err error) {

	return
}

// execute the exit command
func (channel *channel) exit() {

}
