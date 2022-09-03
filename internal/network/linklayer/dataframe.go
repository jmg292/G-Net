package linklayer

import (
	"encoding/json"
	"time"

	"github.com/jmg292/G-Net/internal/datagram"
	gnet "github.com/jmg292/G-Net/pkg/gneterrs"
)

// DataFrame performs LinkLayer datagram framing, similar to
// Ethernet's wire protocol framing.  A DataFrame contains a
// signed and sealed Datagram, along with all metadata required
// to deliver the frame to its intended destination, authenticate
// its origin upon receipt, authorize the sender, unseal the datagram,
// and safely unmarshal the datagram's content before allowing the
// datagram content to be passed to the receiving service.
type DataFrame struct {
	Header
	datagram.Sealed
	contentType      datagram.ContentType
	AuthorizationTag []byte
}

// Frame instantiates and returns a *DataFrame without
func Frame(from, to string, payload datagram.ContentType) *DataFrame {
	return &DataFrame{
		Header:      NewHeader(from, to),
		contentType: payload,
	}
}

// UUID implements LinkLayer for DataFrame
func (packet *DataFrame) UUID() string {
	return string(packet.EphemeralKey)
}

// ParentID implements LinkLayer for DataFrame
func (packet *DataFrame) ParentID() string {
	return packet.Sender
}

// ReplyTo implements LinkLayer for DataFrame
func (packet *DataFrame) ReplyTo() string {
	return packet.Recipient
}

// SetReplyTo implements LinkLayer for DataFrame
func (packet *DataFrame) SetReplyTo() time.Time {
	return packet.Timestamp
}

// Type implements datagram.Opaque for DataFrame
func (packet *DataFrame) Type() string {
	return string(packet.contentType)
}

// Data implements datagram.Opaque for DataFrame
func (packet *DataFrame) Data() (sealed []byte) {
	sealed = append(sealed, packet.EphemeralKey...)
	sealed = append(sealed, packet.Payload...)
	sealed = append(sealed, packet.Signature...)
	return
}

// Marshal implements datagram.Opaque for DataFrame
func (packet *DataFrame) Marshal() ([]byte, error) {
	return json.Marshal(packet)
}

// Unmarshal implements datagram.Opaque for DataFrame
func (packet *DataFrame) Unmarshal(data []byte) (err error) {
	return json.Unmarshal(data, packet)
}

// UnmarshalData implements datagram.Content for DataFrame
func (packet *DataFrame) UnmarshalData(data any) (err error) {
	if sealed, ok := data.(datagram.Sealed); !ok {
		err = gnet.ErrorInvalidDatagram
	} else {
		packet.EphemeralKey = sealed.EphemeralKey
		packet.Payload = sealed.Payload
		packet.Signature = sealed.Signature
	}
	return
}
