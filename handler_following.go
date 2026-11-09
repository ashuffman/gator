package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	feedList, err := s.db.GetFeedFollowsForUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return fmt.Errorf("GetFeedFollowsForUser error: %w", err)
	}

	for idx := range feedList {
		fmt.Printf("* %v\n", feedList[idx].FeedName)
	}

	return nil
}
