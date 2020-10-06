package main

import (
	dao2 "app05/dao"
	"app05/server"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	gotenv.Load()
	port := os.Getenv("PUBLIC_PORT")
	if port == "" {
		log.Fatal("empty PUBLIC_PORT")
		return
	}

	db, err := sqlx.Open("postgres", "postgres:///app05?sslmode=disable")
	if err != nil {
		log.Fatalf("can't connect: %s", err)
		return
	}

	m := chi.NewMux()
	dao := dao2.New(db)
	s := server.NewServer(dao)

	m.Get("/users", s.GetAllUsersHandler)
	m.Post("/users", s.CreateUserHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), m); err != nil {
		log.Fatalf("can't start server: %s", err)
		return
	}
}
