package esl

import (
	"errors"
	"fmt"
	"sync"
)

// OutboundChannel when accept a new connection from FS will derive a new channel
// Outbound pattern
// in this pattern anything could be sync because this object will be maintained by the business goroutine
// there is not thread-safe problem
type OutboundChannel struct {
	connection
	sync.Mutex
	reply    chan *Event
	response chan *Event
	Events   chan *Event
}

// TODO: Execute method execute on the outbound channel

// Loop to receive the events
func (channel *OutboundChannel) Loop() {
	for {
		event, err := channel.recv()
		if err != nil {
			//TODO: should handle the logic with the EOF
			errorm(err)
			break
		}
		if event.Type == EslEvent {
			channel.Events <- event
		}
		if event.Type == EslReplyErr || event.Type == EslReplyOk {
			channel.reply <- event
		}
		if event.Type == EslResponseErr || event.Type == EslResponseOk {
			channel.response <- event
		} else {
			break
		}
	}
}

func (channel *OutboundChannel) syncCommand(command string) (err error) {
	command = fmt.Sprintf("%s\r\n\r\n", command)
	channel.Lock()
	err = channel.send(command)
	if err != nil {
		channel.Unlock()
		errorf("%s execute on outbound connection failed %v", command, err)
		return
	}
	defer channel.Unlock()
	reply := <-channel.reply
	if reply.Type == EslReplyErr {
		err = errors.New(reply.Body)
		return
	}
	return
}

// Connect send the connect command to FS
// Sync method
func (channel *OutboundChannel) Connect() (err error) {
	return channel.syncCommand("connect")
}

// Answer the inbound call from outbound channel
func (channel *OutboundChannel) Answer() (err error) {
	return channel.syncCommand("answer")
}

// Linger send the linger command to FS
// Sync method
func (channel *OutboundChannel) Linger() (err error) {
	return channel.syncCommand("linger")
}

// Nolinger send the noliner command to FS
// Sync method
func (channel *OutboundChannel) Nolinger() (err error) {
	return channel.syncCommand("nolinger")
}
