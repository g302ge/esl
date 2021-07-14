package esl

import (
	"context"
	"net"
)

// Callback user defined handler logic
type Callback = func(ctx context.Context, channel *OutboundChannel)

// TODO: need a connection manager to manage the long connection 

// Server wrapper to use the Outbound pattern of FS
type Server struct {
	net.Listener
	channel  chan struct{}
	ctx      context.Context
	Error    error
	Callback Callback
	Signal   <-chan struct{}
	cancels []context.CancelFunc
}

// NewServer create a new server 
func NewServer() (server *Server) {
	server = &Server{}
	server.channel = make(chan struct{})
	server.Signal = server.channel
	return
}

// Listen on specific port
func (server *Server) Listen(ctx context.Context, address string) (err error) {
	server.ctx = ctx
	server.Listener, err = net.Listen("TCP", address)
	if err != nil {
		return
	}
	go func() {
		for {
			if server.ctx.Err() != nil {
				close(server.channel)
				break
			}
			if conn, e := server.Accept(); e != nil {
				server.Error = e
				close(server.channel) // channel cancel ctx 
				for _, c := range server.cancels{
					c() // all canceled
				}
				break
			} else {
				// create and call the user callback
				c, cancel := context.WithCancel(ctx)
				server.cancels = append(server.cancels, cancel)
				channel := newChannel(c, conn)
				go channel.Run()
				go server.Callback(c, &OutboundChannel{Channel: channel})
			}
		}
	}()
	return
}
