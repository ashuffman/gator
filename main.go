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

	db, err := sql.Open("postgres", cfg.DBurl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	currentState := state{
		db:        dbQueries,
		configPtr: &cfg,
	}

	commands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)

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
