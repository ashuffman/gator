package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ashuffman/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	limit = 2
	if len(cmd.args[0]) > 0 {
		argint, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Error parsing limit parameter: %w\n", err)
		}
		limit = int32(argint)
	}

	params := database.GetPostsForUserParams{
		Name:  user.Name,
		Limit: limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error fetching posts: %w\n", err)
	}

	for i := range posts {
		p := posts[i]
		fmt.Printf(`
Title: %v
URL: %v
Description: %v
Published At: %v
Created At: %v
Updated At: %v
From Feed: %v
		`, p.Title, p.Url, p.Description.String, p.PublishedAt.Time, p.CreatedAt, p.UpdatedAt, p.FeedName)
	}

	return nil
}
