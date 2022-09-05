package packet

import "time"

type Header struct {
	Timestamp time.Time
	Sender    string
	Recipient string
}

func NewHeader(from, to string) Header {
	return Header{
		Timestamp: time.Now(),
		Sender:    from,
		Recipient: to,
	}
}
