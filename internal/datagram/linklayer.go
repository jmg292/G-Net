package datagram

import "time"

// Compatibility with suborbital/grav/grav.Message
type LinkLayer interface {
	Opaque
	Content
	UUID() string
	ParentID() string
	ReplyTo() string
	SetReplyTo(string)
	Timestamp() time.Time
}
