package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ashuffman/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Wrong number of arguments")
	}

	feedURL := cmd.args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("GetFeedByURL error: %w", err)
	}

	username := s.configPtr.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("GetUser error: %w", err)
	}

	feedFollowParams := database.FeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollowSlice, err := s.db.FeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("FeedFollow error: %w", err)
	}

	fmt.Printf("user: %v, feed: %v\n", feedFollowSlice[0].Username, feedFollowSlice[0].FeedName)

	return nil
}
