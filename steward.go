// Package steward provides a drop in way of adding a status endpoint to your
// services
package steward

import (
	"encoding/json"
	"net/http"
)

var (
	endpoint = "/status"
	resp     interface{}
)

func init() {
	http.Handle(endpoint, http.HandlerFunc(HandleStatus))
}

// HandleStatus implements the go http handler interface to
// display specified status details
func HandleStatus(w http.ResponseWriter, req *http.Request) {
	payload, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

// SetStatusResponse overrides the default response payload
// for the status endpoint
func SetStatusResponse(v interface{}) {
	resp = v
}

// SetStatusEndpoint overrides the default endpoint with desired value
func SetStatusEndpoint(e string) {
	endpoint = e
	http.Handle(endpoint, http.HandlerFunc(HandleStatus))
}
