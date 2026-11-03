package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// if the arg slice is empty, return an error
	if len(cmd.args) == 0 {
		return errors.New("Login Failed: No username provided")
	}

	// use the state's access to the config struct to set the user to the given username
	username := cmd.args[0]
	err := s.configPtr.SetUser(username)
	// remember to return any errors
	if err != nil {
		return err
	}

	fmt.Printf("Current Username has been set to %v\n", username)

	return nil
}
