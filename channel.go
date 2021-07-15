package esl

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net"
	"net/textproto"
	"strings"
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
	sync.Once
	ctx      cx
	reply    chan *Event
	response chan *Event
	Events   chan *Event
	Signal   <-chan struct{}
	ch       chan struct{}
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
		sync.Once{},
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
func (channel *Channel) Run(stop <-chan struct{}) error {
	channel.Once.Do(func() {
		for {
			if channel.ctx.Err() != nil {
				if channel.ctx.Context.Err() == nil {
					// check the inner logic
					close(channel.ch)
				}
				errorf("erros: %v", channel.ctx.Err().Error())
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
	})

	go func() {
		<-stop
		channel.connection.Close()
	}()

	return nil
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
	reply := <-channel.reply // reply could change the SC FUCK!
	channel.Unlock()
	if strings.Contains(reply.Body, "-ERR") {
		debugf("channel send command failed %v", err)
		return errors.New(reply.Body)
	}
	return
}

// execute the sendmsg logic with application
func (channel *Channel) execute() (response *Event, err error){
	return
}

// TODO: here need a batch pattern to execute for one connection
// TODO: here also need abtach pattern for execute method

// execute unload a module mod_event_socket
func (channel *Channel) unload() (err error) {
	return channel.command("api bgapi unload mod_event_socket")
}

// execute reload a module mod_event_socket
func (channel *Channel) reload() (err error) {
	return channel.command("api bgapi reload mod_event_socket")
}

// execute the filter command
func (channel *Channel) filter(action string, events ...string) (err error) {
	// filter delete commands...
	// filter add commands...
	es := strings.Join(events, " ")
	return channel.command(fmt.Sprintf("filter %s %s", action, es))
}

// execute the resume command
func (channel *Channel) resume() (err error) {
	// in FS could set this session as LFLAG_RESUME
	// maybe useful for Inbound mode
	return channel.command("resume")
}

// TODO: except these methods other methods should check the auth logic

// execute auth command
func (channel *Channel) auth(password string) (err error) {
	return channel.command(fmt.Sprintf("auth %s", password))
}

// execute userauth command
func (channel *Channel) userauth(username, password string) (err error) {
	return channel.command(fmt.Sprintf("userauth %s %s", username, password))
}

// bellow methods should be authorizated or outbound method

// execute connect command
func (channel *Channel) connect() (event *Event, err error) {
	// connect with connect event 
	// should check the result should return result event response 
	return //channel.command("connect")
}

// execute answer command
func (channel *Channel) answer() (err error) {
	// this should use ececute command 
	return
}

// execute linger command
func (channel *Channel) linger() (err error) {

	return
}

// execute nolinger command
func (channel *Channel) nolinger() (err error) {

	return
}

// execute the getvar command to get the variable of current channel
func (channel *Channel) getvar(key string) (err error) {

	return
}

// execute sendevent command
func (channel *Channel) sendevent() (err error) {

	return
}

// execute command sendmsg
func (channel *Channel) sendmsg() (response *Event, err error) {
	return
}

// execute api command sync
func (channel *Channel) api() (err error) {

	return
}

// execute bgapi command async
func (channel *Channel) bgapi() (err error) {

	return
}

// execute event command to subcribe the events specificed
func (channel *Channel) event(format string, events ...string) (err error) {

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
