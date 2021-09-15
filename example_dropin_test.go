package steward_test

import (
	"net"
	"net/http"

	_ "github.com/gomicro/steward"
)

var (
	host = "0.0.0.0"
	port = "8000"
)

// ExampleDropin demonstrates how to start a service with steward automatically injecting a status endpoint
func Example_dropin() {
	go http.ListenAndServe(net.JoinHostPort(host, port), nil) //nolint:errcheck
	select {}
}
