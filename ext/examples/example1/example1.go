package main

import (
	"net"
	"net/http"

	_ "github.com/gomicro/steward"
)

func main() {
	go http.ListenAndServe(net.JoinHostPort("0.0.0.0", "8000"), nil)
	select {}
}
