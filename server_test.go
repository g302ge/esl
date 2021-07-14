package esl

import (
	"context"
	"testing"
)

func TestSampleServer(t *testing.T) {
	server := NewServer()

	server.Callback = func(ctx context.Context, channel *OutboundChannel) {
		// create a state machine
		// waiting the cancel signal ?
		// loop to receive
		for event := range channel.Events {
			t.Log(event.IntoPlain()) //
		}
	}

	err := server.Listen(context.TODO(), "0.0.0.0:8080")
	if err != nil {
		t.Errorf("Listen failed %v", err)
		return
	}

	<-server.Signal // waiting done
}
