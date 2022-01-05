package main

import (
	"haberdasher/internal/haberdasherserver"
	"haberdasher/rpc/haberdasher"
	"net/http"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
