package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ashuffman/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	// if the arg slice is empty, return an error
	if len(cmd.args) == 0 {
		return errors.New("Login Failed: No username provided")
	}

	username := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		return errors.New("Error: User already exists")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// create a new user in the database
	// pass `context.Background()` to the query to create an empty `Context` argument
	// use `uuid.New()` function to generate a new UUID for the user
	// `created_at` and `updated_at` should be the current time
	// use the provided name
	// exit with a code `1` if a user with that name already exists
	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return err
	}

	// Set the current user in the config to the given name
	err = s.configPtr.SetUser(username)
	if err != nil {
		return err
	}

	// Print a message that the user was created and log the user's data to the console for your own debugging
	fmt.Printf("User %v was created\n", username)
	fmt.Printf("%+v\n", newUser)

	return nil
}
