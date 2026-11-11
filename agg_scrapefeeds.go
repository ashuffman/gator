package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/ashuffman/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkLastFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for idx := range rssFeed.Channel.Item {
		item := rssFeed.Channel.Item[idx]
		err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: time.Now(), Valid: true},
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
