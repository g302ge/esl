package esl

// Client implement the Inbound mode
type Client struct{
	*channel
	address string
	port string
}

// should implement the reconnect logic 