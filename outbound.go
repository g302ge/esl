package esl

// OutboundChannel when accept a new connection from FS will derive a new channel
// Outbound pattern
// in this pattern anything could be sync because this object will be maintained by the business goroutine
// there is not thread-safe problem
type OutboundChannel struct {
	*Channel

	// here need to cancel the registerstion of the 
}

// // TODO: Execute method execute on the outbound channel

// // Loop to receive the events
// func (channel *OutboundChannel) Loop() {
// 	for {
// 		event, err := channel.recv()
// 		if err != nil {
// 			//TODO: should handle the logic with the EOF
// 			errorm(err)
// 			break
// 		}
// 		if event.Type == EslEvent {
// 			channel.Events <- event
// 		}
// 		if event.Type == EslReply{
// 			channel.reply <- event
// 		}
// 		if event.Type ==EslResponse{
// 			channel.response <- event
// 		} else {
// 			break
// 		}
// 	}
// }

// func (channel *OutboundChannel) syncCommand(command string) (err error) {
// 	command = fmt.Sprintf("%s\r\n\r\n", command)
// 	channel.Lock()
// 	err = channel.send(command)
// 	if err != nil {
// 		channel.Unlock()
// 		errorf("%s execute on outbound connection failed %v", command, err)
// 		return
// 	}
// 	channel.Unlock() // FIXME： send ordering is the same reply ordering this is the SC
// 	reply := <-channel.reply
// 	if strings.Contains(reply.Body, "-ERR"){
// 		return errors.New(string(reply.Body))
// 	}
// 	return
// }

// // Execute an application on this channel
// func (channel *OutboundChannel) Execute(application, arg, uuid string) (response *Event, err error) {
// 	var builder strings.Builder
// 	builder.WriteString("sendmsg")
// 	if uuid != "" {
// 		builder.WriteString(fmt.Sprintf(" %s\n", uuid))
// 	} else {
// 		builder.WriteString("\n")
// 	}
// 	builder.WriteString("call-command: execute\n")
// 	builder.WriteString(fmt.Sprintf("execute-app-name: %s\n", application))
// 	if arg != "" {
// 		builder.WriteString(fmt.Sprintf("execute-app-arg: %s\n", arg))
// 	}
// 	builder.WriteString("event-lock: true\n")

// 	channel.Lock()
// 	err = channel.send(builder.String())
// 	if err != nil {
// 		errorf("sendmsg failed %v", err)
// 		return
// 	}
// 	channel.Unlock()
// 	response = <-channel.reply // TODO: fix it ?
// 	// -ERR check logic should be done
// 	return
// }

// // func (channel *OutboundChannel) Execute(application, arg string) ()

// // Connect send the connect command to FS
// // Sync method
// func (channel *OutboundChannel) Connect() (err error) {
// 	return channel.send("command\r\n\r\n")
// }

// // Answer the inbound call from outbound channel
// func (channel *OutboundChannel) Answer() (err error) {
// 	respnse, err := channel.Execute("answer", "", "")
// 	if err != nil {
// 		return err
// 	}
// 	debug(respnse.Body)
// 	return
// }

// // Linger send the linger command to FS
// // Sync method
// // if seconds is 0 will use the default expired time
// func (channel *OutboundChannel) Linger(seconds int) (err error) {
// 	if seconds == 0 {
// 		return channel.syncCommand("linger")
// 	}
// 	return channel.syncCommand(fmt.Sprintf("linger %d", seconds))
// }

// // Nolinger send the noliner command to FS
// // Sync method
// func (channel *OutboundChannel) Nolinger() (err error) {
// 	return channel.syncCommand("nolinger")
// }
