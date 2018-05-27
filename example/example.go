package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skaji/exec-tcp-server/listener"
)

func main() {
	listeners, err := listener.ListenAll()
	if err != nil {
		log.Fatal(err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world\n")
	}
	server := &http.Server{Handler: http.HandlerFunc(handler)}

	if err := server.Serve(listeners[0]); err != nil {
		log.Fatal(err)
	}
}
