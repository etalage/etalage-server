package main

import (
	"net/http"
)

type Router struct {
}

func (*Router) BuildRoute(httpHandler *http.ServeMux) {
	// todo add route url and functions
	httpHandler.HandleFunc("/", process)
}
