package main

import (
	"net"
	"net/http"

	. "github.com/gomicro/steward"
	"github.com/gorilla/mux"
)

type Response struct {
	Application string `json:"app"`
	Version     string `json:"version"`
	BuildTime   string `json:"buildTime"`
}

func main() {
	SetStatusResponse(&Response{Application: "Foo", Version: "1.0.0", BuildTime: "Today"})
	go startServiceLoop()
	select {}
}

func registerEndpoints() http.Handler {
	r := mux.NewRouter()
	r.Path("/foo").HandlerFunc(FooHandler).Methods("GET")
	return r
}

func startServiceLoop() {
	handler := registerEndpoints()
	http.Handle("/", handler)
	http.ListenAndServe(net.JoinHostPort("0.0.0.0", "8000"), nil)
}

func FooHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Foobar"))
}
