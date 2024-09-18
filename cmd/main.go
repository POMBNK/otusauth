package main

import (
	"context"
	authRepo "github.com/POMBNK/gateway/internal/repository/auth"
	authService "github.com/POMBNK/gateway/internal/service/auth"
	"github.com/POMBNK/gateway/internal/transport/http/auth"
	"github.com/POMBNK/gateway/pkg/client/postgres"
	_ "github.com/POMBNK/gateway/statik"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("SECRET", "secret")
	ctx := context.Background()
	dbConn, err := postgres.NewClient(ctx, postgres.Cfg{
		Login:       "pombnk",
		Password:    "postgres",
		Host:        "localhost",
		Port:        "5432",
		Database:    "authdb",
		MaxAttempts: 3,
	})
	if err != nil {
		panic(err)
	}

	storage := authRepo.NewRepo(dbConn)
	svc := authService.NewService(storage)
	authHandler := auth.NewServer(svc)

	log.Println("Listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", authHandler.Register("/api/v1")))
}
