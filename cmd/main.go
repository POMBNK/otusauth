package main

import (
	"context"
	authRepo "github.com/POMBNK/gateway/internal/repository/auth"
	authService "github.com/POMBNK/gateway/internal/service/auth"
	"github.com/POMBNK/gateway/internal/transport/http/auth"
	"github.com/POMBNK/gateway/pkg/client/postgres"
	"github.com/POMBNK/gateway/pkg/jwt"
	_ "github.com/POMBNK/gateway/statik"
	"github.com/rakyll/statik/fs"
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

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger/", http.FileServer(statikFS)).ServeHTTP(w, r)
	})

	storage := authRepo.NewRepo(dbConn)
	svc := authService.NewService(storage)
	authHandler := auth.NewServer(svc)
	mux.HandleFunc("/api/v1/sign-in", authHandler.SignIn)
	mux.HandleFunc("/api/v1/sign-up", authHandler.SignUp)
	mux.HandleFunc("/api/v1/validate", jwt.Middleware(authHandler.Validate, os.Getenv("SECRET")))
	log.Println("Listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
