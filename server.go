package esl

// listen
// accept from connection
// start channel loop
// connect
// answer

// Server of FS outbound event socket
type Server struct {
	Chans chan *Channel
}

