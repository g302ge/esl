package esl

import "errors"

// FreeSwitch EventSocket Protocol
// all the events async sent by FS and responses event are formated in the Event fromat
// Event:
// Headers \n\n
// Body \n
//
// Header format
// HeaderName: HeaderValue\n
//
// The Event which carry the body always have the header
// Content-Length: BodyLength in bytewise
//
// Execute format
// sendmsg\n
// call-command: execute\n
// execute-app-name: app_name\n
// execute-app-arg: app_args_string\n
// async: true\n if ExecuteAsync
// event-lock: true\n Execute

// FS events constant variables
const (
	EslEventCustom                 = "CUSTOM"
	EslEventClone                  = "CLONE"
	EslEventChannelCreate          = "CHANNEL_CREATE"
	EslEventChannelDestory         = "CHANNEL_DESTORY"
	EslEventChannelState           = "CHANNEL_STATE"
	EslEventChannelCallState       = "CHANNEL_CALLSTATE"
	EslEventChannelAnswer          = "CHANNEL_ANSWER"
	EslEventChannelHangup          = "CHANNEL_HANGUP"
	EslEventChannelHangupComplete  = "CHANNEL_HANGUP_COMPLETE"
	EslEventChannelExecute         = "CHANNEL_EXECUTE"
	EslEventChannelExecuteComplete = "CHANNEL_EXECUTE_COMPLETE"
	EslEventChannelHold            = "CHANNEL_HOLD"
	EslEventChannelUnhold          = "CHANNEL_UNHOLD"
	EslEventChannelBridge          = "CHANNEL_BRIDGE"
	EslEventChannelUnbridge        = "CHANNEL_UNBRIDGE"
	EslEventChannelProgress        = "CHANNEL_PROGRESS"
	EslEventChannelProgressMedia   = "CHANNEL_PROGRESS_MEDIA"
	EslEventChannelOutgoing        = "CHANNEL_OUTGOING"
	EslEventChannelPark            = "CHANNEL_PARK"
	EslEventChannelUnpark          = "CHANNEL_UNPARK"
	EslEventChannelApplication     = "CHANNEL_APPLICATION"
	EslEventChannelOriginate       = "CHANNEL_ORIGINATE"
	EslEventChannelUUID            = "CHANNEL_UUID"
	EslEventAPI                    = "API"
	EslEventLog                    = "LOG"
	EslEventInboundChan            = "INBOUND_CHAN"
	EslEventOutboundChan           = "OUTBOUND_CHAN"
	EslEventStartup                = "STARTUP"
	EslEventShutdown               = "SHUTDOWN"
	EslEventPublish                = "PUBLISH"
	EslEventUnpublish              = "UNPUBLISH"
	EslEventTalk                   = "TALK"
	EslEventNoTalk                 = "NOTALK"
	EslEventSessionCrash           = "SESSION_CRASH"
	EslEventModuleLoad             = "MODULE_LOAD"
	EslEventModuleUnload           = "MODULE_UNLOAD"
	EslEventDTMF                   = "DTMF"
	EslEventMessage                = "MESSAGE"
	EslEventPresenceIn             = "PRESENCE_IN"
	EslEventNotifyIn               = "NOTIFY_IN"
	EslEventPresenceOut            = "PRESENCE_OUT"
	EslEventPresenceProbe          = "PRESENCE_PROBE"
	EslEventMessageWaiting         = "MESSAGE_WAITING"
	EslEventMessageQuery           = "MESSAGE_QUERY"
	EslEventRoster                 = "ROSTER"
	EslEventCodec                  = "CODEC"
	EslEventBackgroundJob          = "BACKGROUND_JOB"
	EslEventDetectedSpeech         = "DETECTED_SPEECH"
	EslEventDetectedTone           = "DETECTED_TONE"
	EslEventPrivateCommand         = "PRIVATE_COMMAND"
	EslEventHeartbeat              = "HEARTBEAT"
	EslEventTrap                   = "TRAP"
	EslEventAddSchedule            = "ADD_SCHEDULE"
	EslEventDelSchedule            = "DEL_SCHEDULE"
	EslEventExeSchedule            = "EXE_SCHEDULE"
	EslEventReSchedule             = "RE_SCHEDULE"
	EslEventReloadXML              = "RELOADXML"
	EslEventNotify                 = "NOTIFY"
	EslEventPhoneFeature           = "PHONE_FEATURE"
	EslEventPhoneFeatureSubscribe  = "PHONE_FEATURE_SUBSCRIBE"
	EslEventSendMessage            = "SEND_MESSAGE"
	EslEventRecvMessage            = "RECV_MESSAGE"
	EslEventRequestParams          = "REQUEST_PARAMS"
	EslEventChannelData            = "CHANNEL_DATA"
	EslEventGeneral                = "GENERAL"
	EslEventCommand                = "COMMAND"
	EslEventSessionHeartbeat       = "SESSION_HEARTBEAT"
	EslEventSessionDisconnected    = "CLIENT_DISCONNECTED"
	EslEventServerDisconnected     = "SERVER_DISCONNECTED"
	EslEventSendInfo               = "SEND_INFO"
	EslEventRecvInfo               = "RECV_INFO"
	EslEventRTCPMessage            = "RECV_RTCP_MESSAGE"
	EslEventCallSecure             = "CALL_SECURE"
	EslEventNat                    = "NAT"
	EslEventRecordStart            = "RECORD_START"
	EslEventRecordStop             = "RECORD_STOP"
	EslEventPlaybackStart          = "PLAYBACK_START"
	EslEventPlaybackStop           = "PLAYBACK_STOP"
	EslEventCallUpdate             = "CALL_UPDATE"
	EslEventFailure                = "FAILURE"
	EslEventSocketData             = "SOCKET_DATA"
	EslEventMediaBugStart          = "MEDIA_BUG_START"
	EslEventMediaBugStop           = "MEDIA_BUG_STOP"
	EslEventConferenceDataQuery    = "CONFERENCE_DATA_QUERY"
	EslEventConferenceData         = "CONFERENCE_DATA"
	EslEventCallSetupReq           = "CALL_SETUP_REQ"
	EslEventCallSetupResult        = "CALL_SETUP_RESULT"
	EslEventCallDetail             = "CALL_DETAIL"
	EslEventDeviceState            = "DEVICE_STATE"
	EslEventText                   = "TEXT"
	EslEventShutdownRequested      = "SHUTDOWN_REQUESTED"
	EslEventAll                    = "ALL"
)

// esl normal headers
const (
	EslContentType   = "Content-Type"
	EslContentLength = "Content-Length"
	EslEventName     = "Event-Name"
	EslReplyText     = "Reply-Text"

	// TODO: other logic
)

// Content-Type of Event
const (
	EslEventContentPlain            = "text/event-plain"
	EslEventContentJson             = "text/event-json"
	EslEventContentApiResponse      = "api/response"
	EslEventContentCommandReplay    = "command/reply"
	EslEventContentDiscoonectNotice = "text/disconnect-notice"
)

// Event Type
const (
	EslEvent = iota // text/event-plain and text/event-json
	EslReply // command/reply +OK -ERR
	EslResponse // api/response +OK -ERR
	EslDisconnectedNotice // text/disconnect-notice
)

// Event Error defines
var (
	ErrHeaderNotFound  = errors.New("header not found")
	ErrJsonBodyParsing = errors.New("parsing JSON from body failed")
)

// Event of FS
type Event struct {
	Name     string              // Event Name e.g. CHANNEL_CREATE
	Subclass string              // Event subclass name
	Type     int                 // Event content typetext/event-plain
	Headers  map[string][]string // Event Headers
	Body     string              // Event Body always string type
}

// GetHeader from current event
// event keep that the header name is unique in the total event
func (event *Event) GetHeader(name string) (value []string, err error) {
	//TODO: foreach to get current header
	err = ErrHeaderNotFound
	return
}

// AddHeader is the helper function to mainpulate the inner headers
func (event *Event) AddHeader(name, value string) {

}

// IntoJson serialize the event to json
func (event *Event) IntoJson() (json string, err error) {
	err = ErrJsonBodyParsing
	return
}

// IntoPlain serialize the event to plaintext
func (event *Event) IntoPlain() (plain string, err error) {
	return
}

// Merge other event to current event
func (event *Event) Merge(rhs *Event) (err error) {

	return
}
