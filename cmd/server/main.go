package main

import (
	"context"
	"net/http"

	"github.com/learngofarsi/go-basics-project/internal/db"
	"github.com/learngofarsi/go-basics-project/internal/db/postgres"
	"github.com/learngofarsi/go-basics-project/pkg/config"
	"github.com/learngofarsi/go-basics-project/pkg/server"
)

func main() {

	cnf := config.LoadConfigOrPanic()
	pg, err := postgres.NewPostgres(cnf.Postgres)
	if err != nil {
		panic(err)
	}

	if err := db.Migrate(context.Background(), pg); err != nil {
		panic(err)
	}

	server := server.NewHttpServer(cnf.Server)
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	server.Start()
}
