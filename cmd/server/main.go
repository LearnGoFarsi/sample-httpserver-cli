package main

import (
	"net/http"

	"github.com/learngofarsi/go-basics-project/pkg/config"
	"github.com/learngofarsi/go-basics-project/pkg/server"
)

func main() {
	cnf := config.LoadConfigOrPanic()
	server := server.NewHttpServer(cnf)
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server.Start()
}
