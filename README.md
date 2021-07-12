# esl

Yet another Go ESL for FreeSwitch :)

## Motivation

Because of the closed chinese FreeSwitch community and the bulshit document on FreeSwitch wiki, I have to view both source code of FreeSwitch and do too many useless experiments in these day to figure out the correct business implementation ways and clustering deploying scheme.The other motivation is that the all Go implmentations of ESL are not soundly.

At this last `Fuck` all the `Closed` developers with FreeSwitch who always `Zhuangbi`, meaning for One bottle is dissatisfied and half bottle is sloshing, in the fuck community on QQ the biggest IM in china, especially the nway.cn CEO who now maintain the China FreeSwitch comunity and always `Kaiche`, meaning for broadcasting the porn views in community.

So I developed this library help developer to control the behaviors of FreeSwitch both Inbound and Outbound pattern.


## Inbound Client Example

```go
// TODO...
```

## Outbound Server Example

```go
// TODO...
```

## Best practice in ESL

Well, the event socket protocol is very strange in morden network protocol design. In network application every event is async, so in normal every request in application should carry a unique sequence and the server resposne need the same unique sequence to implment the ACK pattern. But in FS ESL lib writen in C there is only a Fuck global lock.

In call center business, there are 

1. customer call -> IVR
2. customer call -> IVR -> ACD -> agent call
3. agent call -> ACD -> customer call
4. agent call -> agent call

In this library the `OutboundChannel` is that accepted from the Server socket, and implemented using the lock to implement the pattern 1. 
