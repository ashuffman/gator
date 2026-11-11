package main

import (
	"context"
	"fmt"
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
		fmt.Printf("* %v\n", rssFeed.Channel.Item[idx])
	}

	return nil
}
