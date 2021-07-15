package esl

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"sync"
	"sync/atomic"
)

// Callback user defined handler logic
type Callback = func(channel *OutboundChannel)

// TODO: need a connection manager to manage the long connection

// Server wrapper to use the Outbound pattern of FS
type Server struct {
	net.Listener
	address  string
	Signal   <-chan struct{}
	ch       chan struct{}
	ctx      context.Context
	Error    error
	Callback Callback
	channels sync.Map //
	running  atomic.Value
}

// NewServer create a new server
func NewServer(ctx context.Context, address string) (server *Server) {
	server = &Server{}
	server.ch = make(chan struct{})
	server.Signal = server.ch
	server.address = address
	server.ctx = ctx
	return
}

// Shutdown the Outbound server
func (server *Server) Shutdown() {
	running, _ := server.running.Load().(bool)
	if !running {
		return
	}
	server.running.Store(false)
	server.Close()
	server.channels.Range(func(key, value interface{}) bool {
		server.channels.Delete(key)
		ch := value.(*channel)
		ch.shutdown()
		<-ch.signal
		return true
	})
	close(server.ch)
}

// Listen on specific port
func (server *Server) Listen() (err error) {
	server.Listener, err = net.Listen("TCP", server.address)
	if err != nil {
		return
	}
	server.running.Store(true)
	for {
		running, _ := server.running.Load().(bool)
		if !running {
			break
		}
		if server.ctx.Err() != nil {
			break
		}
		if conn, e := server.Accept(); e != nil {
			server.Error = e
			close(server.ch) // channel cancel ctx
			server.Shutdown()
			break
		} else {
			c, _ := conn.(*net.TCPConn)
			c.SetNoDelay(true)
			f, _ := c.File()
			fd := f.Fd()
			ch := &channel{
				connection: connection{
					conn,
					textproto.NewReader(bufio.NewReader(conn)),
				},
				reply:    make(chan *Event),
				response: make(chan *Event),
				Events:   make(chan *Event),
				close: func() {
					server.channels.Delete(fd) // unregister the channel
				},
				signal: make(chan struct{}),
			}
			server.channels.Store(fd, ch)
			go ch.loop()
			go server.Callback(&OutboundChannel{ch})
		}
	}

	// Clear function
	server.Shutdown()
	return
}
