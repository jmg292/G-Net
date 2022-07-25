module github.com/jmg292/G-Net

go 1.18

require (
	github.com/cretz/bine v0.2.0
	github.com/go-piv/piv-go v1.9.0
	github.com/google/go-tpm v0.3.3
	github.com/google/uuid v1.3.0
	github.com/mattn/go-sqlite3 v1.14.14
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8
)

require golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect

replace github.com/google/go-tpm => github.com/jmg292/go-tpm v0.3.3
