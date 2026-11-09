package main

import (
	"context"
	"fmt"

	"github.com/ashuffman/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedList, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("GetFeedFollowsForUser error: %w", err)
	}

	for idx := range feedList {
		fmt.Printf("* %v\n", feedList[idx].FeedName)
	}

	return nil
}
