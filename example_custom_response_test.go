package steward_test

import (
	"net"
	"net/http"

	"github.com/gomicro/steward"
)

// StatusResponse is the data for the status endpoint to display
type StatusResponse struct {
	Application string `json:"app"`
	Version     string `json:"version"`
	BuildTime   string `json:"buildTime"`
}

// ExampleCustomResponse demonstrates how to configure a custom response
func Example_customResponse() {
	steward.SetStatusResponse(&StatusResponse{Application: "Foo", Version: "1.0.0", BuildTime: "Today"})

	go http.ListenAndServe(net.JoinHostPort("0.0.0.0", "8000"), nil) //nolint:errcheck
	select {}
}
