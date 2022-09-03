module github.com/jmg292/G-Net

go 1.18

require (
	github.com/awnumar/memguard v0.22.3
	github.com/go-piv/piv-go v1.10.0
	github.com/google/go-tpm v0.3.3
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
)

require github.com/awnumar/memcall v0.1.2 // indirect

require (
	github.com/VirusTotal/gyp v0.8.0 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/suborbital/grav v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/google/go-tpm => github.com/jmg292/go-tpm v0.3.3
