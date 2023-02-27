package main

import (
	"net/http"

	"github.com/learngofarsi/go-basics-project/pkg/server"
)

func main() {
	server := server.NewHttpServer("localhost", 9090)
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server.Start()
}
