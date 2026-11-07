package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ashuffman/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("wrong number of arguments provided")
	}

	user, err := s.db.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	newFeed, err := s.db.CreateFeed(context.Background(), feedParams)
	fmt.Printf("%+v", newFeed)

	return nil
}
