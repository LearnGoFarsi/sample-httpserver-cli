package main

import (
	"context"

	"github.com/learngofarsi/go-basics-project/internal/db"
	"github.com/learngofarsi/go-basics-project/internal/db/postgres"
	"github.com/learngofarsi/go-basics-project/internal/handler"
	"github.com/learngofarsi/go-basics-project/internal/repo"
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

	dbRepo := repo.NewTrackRepo(pg)
	trackHandler := handler.NewTrackHandler(dbRepo)

	server := server.NewHttpServer(cnf.Server)
	server.HandleFunc("/track", trackHandler.Handle)

	server.Start()
}
