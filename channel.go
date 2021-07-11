package esl

import (
	"net"
)

// Channel should be thread safe
// and every client could use this Channel abstraction in multi-tenant system building
// if the connection need the queue ?

// sendmsg
// sendevent
// connect
// auth ClueCon
// userauth username ClueCon

// simple wrapper of the net.Conn
type connection struct {
	net.Conn
}

// send plain text command over net.Conn
func (c *connection) send(command string) (err error) {
	// ioutil.
	// should implement the write implementation
	return
}

// TODO: implement the splite package from TCP stream

// Channel implementors ServerChannel ClientChannel
type Channel interface {
	// FIXME: every execute shoulde be asyn if there some reasons about waiting ?
	// Execute the application in dialplan and waiting retun respon as Event format
	Execute(application, args, uuid string) (response *Event, err error)

	// Api sync execute api
	Api(command, args string)

	// Bgapi async execute api
	Bgapi(command, args, uuid string) // should return the Job-UUID

	// Filter apply a filter on this socket to receive the specific event
	Filter(header, value string) (err error)

	// Close the inner scoket and release some resources
	Close()

	// these pair functions are used in the UNIX socket transfering
	// you could keep the application long link without disconnected from parent to child proceess
	// it's useful the time you have to update your application whenever need
	// FIXME: if these functions should be stateless or stateful?
	IntoRaw() net.Conn
	FromRaw(conn net.Conn)
}

// both ServerChannel and ClientChannel implement the Channel interface

// ServerChannel when accept a new connection from FS will derive a new channel
// Outbound pattern
type ServerChannel struct {
}

// Connect send the connect command to FS
// Sync method
func (channel *ServerChannel) Connect() (err error) {
	return
}

// Linger send the linger command to FS
// Sync method
func (channel *ServerChannel) Linger() (err error) {
	return
}

// Nolinger send the noliner command to FS
// Sync method
func (channel *ServerChannel) Nolinger() (err error) {
	return
}

// ClientChannel is the simple wrapper of the
// Inbound pattern
type ClientChannel struct {
}

// Auth send the auth command with password to FS
// Sync method
func (channel *ClientChannel) Auth(password string) (err error) {
	return
}

// Userauth send the userauth command with username and password to FS
// Sync method
func (channel *ClientChannel) Userauth(username, password string) (err error) {

	return
}

// Events send the events command to FS
// Sync method
func (channel *ClientChannel) Events(category string, evetns []string) (err error) {

	return
}