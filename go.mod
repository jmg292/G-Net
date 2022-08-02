module github.com/jmg292/G-Net

go 1.18

require (
	github.com/cretz/bine v0.2.0
	github.com/go-piv/piv-go v1.9.0
	github.com/google/go-tpm v0.3.3
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
)

require (
	golang.org/x/net v0.0.0-20220728211354-c7608f3a8462 // indirect
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
)

replace github.com/google/go-tpm => github.com/jmg292/go-tpm v0.3.3
