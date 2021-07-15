package esl

// Client implement the Inbound mode
type Client struct {
	*channel
	address string
	port    string
}

// should implement the reconnect logic

// Dial the remote FS
func (client *Client) Dial() (err error) {
	return
}
