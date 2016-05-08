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

func SetStatusResponse(v interface{}) {
	resp = v
}

func SetStatusEndpoint(e string) {
	endpoint = e
	http.Handle(endpoint, http.HandlerFunc(HandleStatus))
}
