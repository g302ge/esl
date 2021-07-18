package esl

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

// channel and the protocol wrappers

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

// execute command sync without reply
func (channel *channel) noreplycmd(cmd string) (err error) {
	_, err = channel.replycmd(cmd)
	return
}

// execute command sync with reply
func (channel *channel) replycmd(cmd string) (reply *Event, err error) {
	cmd = fmt.Sprintf("%s\n\n", cmd)
	channel.Lock()
	err = channel.send(cmd)
	if err != nil {
		//FIXME: in FS this situation will cause the socket close FUCK!
		errorf("channel execute command failed %v", err)
		channel.err = err
		return
	}
	reply = <-channel.reply
	channel.Unlock()
	if strings.Contains(reply.Body, "-ERR") {
		//FIXME: in FS this situation will cause the socket close FUCK!
		debugf("channel send command failed %v", err)
		return nil, errors.New(reply.Body)
	}
	return
}

// execute command sendmsg
func (channel *channel) sendmsg(body, uuid string) (response *Event, err error) {
	builder := strings.Builder{}
	builder.WriteString("sendmsg")
	if uuid != "" {
		builder.WriteString(fmt.Sprintf("%s\n", uuid))
	} else {
		builder.WriteString("\n")
	}
	builder.WriteString(fmt.Sprintf("%s\n", body)) //FIXME: maybe the \n is useless

	channel.Lock()
	err = channel.send(builder.String())
	if err != nil {
		//FIXME: in FS this situation will cause the socket close FUCK!
		errorf("channel execute command failed %v", err)
		channel.err = err
		return
	}
	response = <-channel.reply // TODO: check if really use the reply channel ?
	channel.Unlock()
	if strings.Contains(response.Body, "-ERR") {
		//FIXME: in FS this situation will cause the socket close FUCK!
		debugf("channel send command failed %v", err)
		return nil, errors.New(response.Body)
	}
	return
}

// execute the sendmsg logic with application
func (channel *channel) execute(application, arg, uuid string) (response *Event, err error) {
	builder := strings.Builder{}
	builder.WriteString("call-command: execute\n")
	builder.WriteString(fmt.Sprintf("execute-app-name: %s\n", application))
	if arg != "" {
		builder.WriteString(fmt.Sprintf("execute-app-arg: %s\n", arg))
	}
	//FIXME: arg need be urlencoded
	builder.WriteString("event-lock: true\n") // the sendmsg call will fill the final \n
	response, err = channel.sendmsg(builder.String(), uuid)
	return
}

// execute unload a module mod_event_socket
func (channel *channel) unload() (err error) {
	return channel.noreplycmd("api bgapi unload mod_event_socket")
}

// execute reload a module mod_event_socket
func (channel *channel) reload() (err error) {
	return channel.noreplycmd("api bgapi reload mod_event_socket")
}

// execute the filter command
func (channel *channel) filter(action string, events ...string) (err error) {
	// filter delete commands...
	// filter add commands...
	es := strings.Join(events, " ")
	return channel.noreplycmd(fmt.Sprintf("filter %s %s", action, es))
}

// execute the resume command
// TODO: should know when should recover the session in the listener strtuct 
// should review the FS source code to figure out this fuck thing
func (channel *channel) resume() (err error) {
	// in FS could set this session as LFLAG_RESUME
	// maybe useful for Inbound mode
	return channel.noreplycmd("resume")
}

// TODO: except these methods other methods should check the auth logic

// execute auth command
func (channel *channel) auth(password string) (err error) {
	return channel.noreplycmd(fmt.Sprintf("auth %s", password))
}

// execute userauth command
func (channel *channel) userauth(username, password string) (err error) {
	return channel.noreplycmd(fmt.Sprintf("userauth %s %s", username, password))
}

// bellow methods should be authorizated or outbound method

// execute connect command
func (channel *channel) connect() (response *Event, err error) {
	return channel.replycmd("connect")
}

// execute linger command
func (channel *channel) linger(second int) (err error) {
	return channel.noreplycmd(fmt.Sprintf("linger %d", second))
}

// execute nolinger command
func (channel *channel) nolinger() (err error) {
	return channel.noreplycmd("nolinger")
}

// execute the getvar command to get the variable of current channel
func (channel *channel) getvar(key string) (result string, err error) {
	r, err := channel.replycmd(fmt.Sprintf("getvar %s", key))
	if err != nil {
		return
	}
	result = r.Body // this command result will be emplaced in the Body
	return
}

// execute sendevent command
func (channel *channel) sendevent(event *Event) (err error) {

	return
}

// execute api command sync
func (channel *channel) api() (err error) {

	return
}

// execute bgapi command async
// FIXME: return response ? or Job-UUID ?
func (channel *channel) bgapi() (err error) {

	return
}

// execute event command to subcribe the events specificed
func (channel *channel) event(format string, events ...string) (err error) {
	es := strings.Join(events, " ")
	return channel.noreplycmd(fmt.Sprintf("event %s %s", format, es))
}

// execute nixevent cancel events subcribe
func (channel *channel) nixevent(events ...string) (err error) {
	es := strings.Join(events, " ")
	return channel.noreplycmd(fmt.Sprintf("nixevent %s", es))
}

// execute noevents command
func (channel *channel) noevents() (err error) {
	return channel.noreplycmd("noevents")
}

// execute divert_events command
func (channel *channel) divertEvents(open bool) (err error) {
	if open {
		return channel.noreplycmd("divert_events on")
	}
	return channel.noreplycmd("divert_events off")
}

// execute myevents command
func (channel *channel) myevents(uuid, format string) (err error) {
	if format != "" {
		channel.noreplycmd(fmt.Sprintf("myevents %s %s", uuid, format))
	}
	return channel.noreplycmd("myevents")
}

// execute the exit command
func (channel *channel) exit() {
	channel.noreplycmd("exit")
	// TODO: should change the channel state?
	// which could make sync closed ?
	// but in some concurrent situation this thing could make memory leak
	// see more about the SC
	channel.running.Store(false)
}
