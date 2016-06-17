package cachet

// Message Cachet data model
type Message struct {
	Incident int    `json:"incident"`
	Status   int    `json:"status"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

// Message of current incident
func (cfg *CachetMonitor) Message(incident int, status int) *Message {
	for _, message := range cfg.Messages {
		if message.Incident == incident && message.Status == status {
			return message
		}
	}
	return &Message{}
}
