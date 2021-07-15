package esl

// Batch mode to execute multi commands and api for one call

// Batch operations waiting the batch response
type Batch struct {
	Result  <-chan *Event
	command []string
}

// TODO: need change the channel methods first to format command and api and then batch execute

// Connect execute the connect command
func (batch *Batch) Connect() *Batch {
	return batch
}
