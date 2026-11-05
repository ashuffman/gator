package main

import (
	"errors"
	"log"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	// run the given command with the provided state, if it exists
	handler, ok := c.commandMap[cmd.name]
	if !ok {
		return errors.New("Error: Command not registered.")
	}

	err := handler(s, cmd)
	return err
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func parseCommand(args []string) command {
	if len(args) == 1 {
		log.Fatal("Error: Not enough arguments.")
	}

	var command command
	command.name = args[1]
	command.args = args[2:]
	return command
}
