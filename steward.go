package steward

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

var (
	statusEndpoint = "/status"
	resp           *Response
)

func init() {
	appStr := strings.Split(os.Args[0], "/")
	app := appStr[len(appStr)-1]
	resp = &Response{app}
	http.Handle(statusEndpoint, http.HandlerFunc(Status))
}

type Response struct {
	Application string `json:"app"`
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
