package main

import (
	"context"
	"fmt"

	"github.com/ashuffman/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedparams := database.FeedUnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.FeedUnfollow(context.Background(), feedparams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed %v successfully unfollowed\n", feed.Name)
	return nil
}
