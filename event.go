package esl

import "errors"

// FreeSwitch EventSocket Protocol
//

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

// Content-Type of Event
const (
	EslEventContentPlain         = "text/event-plain"
	EslEventContentJson          = "text/event-json"
	EslEventContentApiResponse   = "api/response"
	EslEventContentCommandReplay = "command/reply"
)

// Event Error defines
var (
	ErrHeaderNotFound  = errors.New("header not found")
	ErrJsonBodyParsing = errors.New("parsing JSON from body failed")
)

// Header of the Event in FS
type Header struct {
	Name  string
	Value []string // FIXME: in C lib is the idx to specific the Header more than one
}

// Event of FS
type Event struct {
	Name     string   // Event Name e.g. CHANNEL_CREATE
	Owner    string   // Event owner
	Subclass string   // Event subclass name
	Headers  []Header // Event Headers
	Body     []byte   // Event Body
	BindData []byte   // BindData from the subclass provider
	UserData []byte   // UserData of user, but now don't know how to use this field
	Key      int64    // Key like UUID
}

// GetHeader from current event
// event keep that the header name is unique in the total event
func (event *Event) GetHeader(name string) (header *Header, err error) {
	//TODO: foreach to get current header
	err = ErrHeaderNotFound
	return
}

// Json serialize the event to json
func (event *Event) Json() (json string, err error) {
	err = ErrJsonBodyParsing
	return
}

// Merge other event to current event
func (event *Event) Merge(rhs *Event) (err error) {

	return
}
