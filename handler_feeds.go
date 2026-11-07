package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feedList, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching feeds: %w\n", err)
	}

	for _, feed := range feedList {
		fmt.Printf("* name: %v, url: %v, user: %v\n", feed.Name, feed.Url, feed.Username)
	}
	return nil
}
