package line

type message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func marshalTextMessage(text string) message {
	return message{
		"text",
		text,
	}
}
