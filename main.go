package main

import (
	"article-reminder/domain"
	"article-reminder/handlers"
	"article-reminder/postgres"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"
)

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "article-reminder",
	})

	defer DB.Close()

	domainDB := domain.DB{UserRepo: postgres.NewUserRepo(DB)}

	d := &domain.Domain{DB: domainDB}

	r := handlers.SetupRouter(d)

	port := os.Getenv("POST")
	if port == "" {
		port = "8081"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}
}
