package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ashuffman/gator/internal/config"
	"github.com/ashuffman/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfg.DBurl = "postgres://angelahuffman:@localhost:5432/gator?sslmode=disable"
	db, err := sql.Open("postgres", cfg.DBurl)
	dbQueries := database.New(db)

	currentState := state{
		db:        dbQueries,
		configPtr: &cfg,
	}

	commands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	err = commands.register("login", handlerLogin)
	if err != nil {
		log.Fatal(err)
	}
	err = commands.register("register", handlerRegister)
	if err != nil {
		log.Fatal(err)
	}

	command := parseCommand(os.Args)
	err = commands.run(&currentState, command)
	if err != nil {
		log.Fatal(err)
	}
}

type state struct {
	db        *database.Queries
	configPtr *config.Config
}
