package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Must include time between requests argument")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
