package main

import (
	"log"
	"os"

	"github.com/ashuffman/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	currentState := state{
		configPtr: &cfg,
	}

	commands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}
	err = commands.register("login", handlerLogin)
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
	configPtr *config.Config
}
