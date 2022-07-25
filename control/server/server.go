package server

import (
	"net/http"
)

func getIdentityCertificate() []byte {
	return make([]byte, 0)
}

func NewControlServeMux() *http.ServeMux {
	controlServeMux := http.NewServeMux()
	controlServeMux.HandleFunc("/identity", func(w http.ResponseWriter, r *http.Request) {
		w.Write(getIdentityCertificate())
	})
	controlServeMux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Registration
	})
	controlServeMux.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Authentication
	})
	controlServeMux.HandleFunc("/network/info", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Local network information
	})
	controlServeMux.HandleFunc("/network/configuration", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Network configuration retrieval
		// TODO: Network configuration updates
	})
	controlServeMux.HandleFunc("/network/connect", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Connection negotiation
	})
	return controlServeMux
}
