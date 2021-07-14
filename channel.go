package esl

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/textproto"
	"sync"
	"time"
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

// inner context implment the Context
type cx struct {
	context.Context
	err error
}

func (c *cx) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

func (c *cx) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *cx) Err() error {
	if c.err == nil {
		return c.Context.Err()
	}
	return c.err
}

func (c *cx) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}

// Channel wrapper of the connection with specific methods
// parent -> child -> current
// cancel the current is correct ?
type Channel struct {
	connection
	sync.Mutex
	ctx      cx
	reply    chan *Event
	response chan *Event
	Events   chan *Event
	Signal <-chan struct{}
	ch chan struct{}
}

// channel builder used in inner package
func newChannel(ctx context.Context, conn net.Conn) *Channel {
	ch := make(chan struct{})
	c := &Channel{
		connection{
			conn,
			textproto.NewReader(bufio.NewReader(conn)),
		},
		sync.Mutex{},
		cx{ctx, nil},
		make(chan *Event),
		make(chan *Event),
		make(chan *Event),
		ch,
		ch,
	}
	return c
}

// Run the Channel
// in your application should create a new goroutine to loop run this function
func (channel *Channel) Run() {
	for {
		if channel.ctx.Err() != nil {
			if channel.ctx.Context.Err() == nil {
				// check the inner logic
				close(channel.ch)
			}
			break
		}
		event, err := channel.recv()
		if err != nil {
			errorm(err)
			channel.ctx.err = err
			continue
		}
		if event.Type == EslEvent {
			channel.Events <- event
		}
		if event.Type == EslReply {
			channel.reply <- event
		}
		if event.Type == EslResponse {
			channel.response <- event
		} else {
			break
		}
	}
	channel.connection.Close()
}

// Alive return the Channel  state
func (channel *Channel) Alive() bool {
	return channel.ctx.Err() == nil
}

// execute command sync
func (channel *Channel) command(cmd string) (err error) {
	cmd = fmt.Sprintf("%s\r\n\r\n", cmd)
	channel.Lock()
	err = channel.send(cmd)
	if err != nil {
		errorf("channel execute command failed %v", err)
		channel.ctx.err = err
		return
	}
	channel.Unlock()
	return
}

// execute unload a module
func (channel *Channel) unload(module string) (err error) {
	return
}

// execute reload a module
func (channel *Channel) reload(module string) (err error) {

	return
}

// execute command sendmsg
func (channel *Channel) sendmsg() (response *Event, err error) {
	return
}

// execute connect command
func (channel *Channel) connect() (err error){

	return
}

// execute answer command 
func (channel *Channel) answer() (err error){
	return
}

// execute linger command 
func (channel *Channel) linger() (err error){

	return
}

// execute nolinger command 
func (channel *Channel) nolinger() (err error){

	return
}

// execute the exit command
func (channel *Channel) exit() {

}

// sendmsg
// sendevent
// filter
// connect
// answer
// getvar
// myevents
// divert_events
// api
// bgapi
// log
// linger
// nolinger
// nolog
// event
// nixevent
// noevents
// resume



// Close the channel normally
func (channel *Channel) Close() {

}


// FS strange event model
// Generaly speaking the Events on Network should be Sequential consistency
//
// SendApi -> ReceiveApiResponse
// SendApi -> ReceiveEvent -> ReceiveApiResponse
//
// because that the if the same time FS received the SendApi and prepare to execute, also receive INVITE of SIP generated a CHANNEL_EVENT
// so any sync all should be blocked before the SendApi finish
//
// but there is impossible in this ordering
// ReceiveApiResponse2 -> ReceiveApiResponse1
//
// In Outbound accepted channel there is nothing to worry about becuase that all the channel is loop in their own goroutine
// But in Inbound pattern everything could be complex
// if we use the Global lock to sync to keep the Sequential consistency
// the application which using this library could be inefficiencily
// so how to handle this situation ?
// FIXME: should use two inbound connection one to receive the event and one execute command ?

/// TODO: async method use the UUID to handle the fuck model ?

// ClientChannel is the simple wrapper of the
// Inbound pattern
type InboundChannel struct {
	Channel
}

// Auth send the auth command with password to FS
// Sync method
func (channel *InboundChannel) Auth(password string) (err error) {
	return
}

// Userauth send the userauth command with username and password to FS
// Sync method
func (channel *InboundChannel) Userauth(username, password string) (err error) {

	return
}

// Events send the event command to FS
// Sync method
func (channel *InboundChannel) Events(category string, evetns []string) (err error) {

	return
}

// Noevents send noevents command to FS
func (channel *InboundChannel) Noevents() (err error) {
	return
}
