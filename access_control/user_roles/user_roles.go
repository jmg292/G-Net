package userroles

type UserRole uint16

const (
	Owner UserRole = iota
	Administrator
	Operator
	Unprivileged
)
