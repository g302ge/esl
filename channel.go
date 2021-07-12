package esl

import (
	"sync"
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


// Channel implementors ServerChannel ClientChannel
// see details https://freeswitch.org/confluence/display/FREESWITCH/mod_commands
type Channels interface {

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

// Channel the sync mode control the event socket 
type Channel struct{
	connection
	sync.Mutex
}




// OutboundChannel when accept a new connection from FS will derive a new channel
// Outbound pattern
// in this pattern anything could be sync because this object will be maintained by the business goroutine
// there is not thread-safe problem
type OutboundChannel struct {
	connection
	sync.Mutex
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
