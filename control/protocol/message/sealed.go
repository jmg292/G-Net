package message

type SealedMessage struct {
	Sender    string
	Header    string
	Payload   string
	Signature string
}
