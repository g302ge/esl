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

// Channel implementors ServerChannel ClientChannel
// see details https://freeswitch.org/confluence/display/FREESWITCH/mod_commands
type Channel interface {
	// Execute the command async
	Execute(application, args, uuid string) (response *Event, err error)

	// Execute the command aync will return the Job-UUID event
	ExecuteAsync(application, args, uuid string) (response *Event, err error)

	// Api sync execute api
	Api(command, args string) (response *Event, err error)

	// Bgapi async execute api
	Bgapi(command, args, uuid string) (response *Event, err error) // should return the Job-UUID

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

// OutboundChannel when accept a new connection from FS will derive a new channel
// Outbound pattern
type OutboundChannel struct {
	connection
}

// Connect send the connect command to FS
// Sync method
func (channel *OutboundChannel) Connect() (err error) {
	return
}

// Linger send the linger command to FS
// Sync method
func (channel *OutboundChannel) Linger() (err error) {
	return
}

// Nolinger send the noliner command to FS
// Sync method
func (channel *OutboundChannel) Nolinger() (err error) {
	return
}

// ClientChannel is the simple wrapper of the
// Inbound pattern
type InboundChannel struct {
	connection
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
