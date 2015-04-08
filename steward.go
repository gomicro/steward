package steward

import (
	"encoding/json"
	"net/http"
)

var (
	statusEndpoint = "/status"
	resp           interface{}
)

func init() {
	http.Handle(statusEndpoint, http.HandlerFunc(Status))
}

func Status(w http.ResponseWriter, req *http.Request) {
	payload, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func StatusResponse(v interface{}) {
	resp = v
}
