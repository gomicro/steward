package main

import (
	"net"
	"net/http"

	_ "github.com/gomicro/steward"
	"github.com/gorilla/mux"
)

func main() {
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
