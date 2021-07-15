package esl

import (
	"context"
	"testing"
)

func TestSampleServer(t *testing.T) {
	server := NewServer(context.TODO(),"0.0.0.0:8080")

	server.Callback = func(channel *OutboundChannel) {
		// create a state machine
		// waiting the cancel signal ?
		// loop to receive
		for event := range channel.Events {
			t.Log(event.IntoPlain()) //
		}
	}

	err := server.Listen()
	if err != nil {
		t.Errorf("Listen failed %v", err)
		return
	}

	<-server.Signal // waiting done
}
