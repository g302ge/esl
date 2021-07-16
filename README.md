# esl

Yet another Go ESL for FreeSwitch, to build lightweight FS event socket application :)

## Inbound Client Example

```go
// TODO...
```

## Outbound Server Example

```go
func main() {
  server := NewServer(context.TODO(),"0.0.0.0:8080")
  server.Callback = func(channel *OutboundChannel) {
    channel.Connect()
    channel.Answer()
    // TODO: create your own state machine model
		for event := range channel.Events {
			switch event.Name {
        case EslEventCustom: 
        {// custom event logic
            channel.Hangup()
        }
        case EslEventChannelHangupComplete: // create a new cdr send to channel
      }
    }
    channel.Shutdown()
	}
}
```

## Event Socket concurrency pattern

What is SC?

TODO:

## How to build your FS control logic 

TODO:

## Clearfy FS Manual and Concepts

Here is a simple clearfy FS manual and introduce some confused concepts in FS not only Channel lifecycle and more.

TODO: