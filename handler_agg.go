package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ewpt3ch/gator-bootdev/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("We need a time interval, 1s, 1m, usage: <time-interval>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %v", timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	markFeedFetchedParams := database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now(),
	}
	err = s.db.MarkFeedFetched(context.Background(), markFeedFetchedParams)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, entry := range rssFeed.Channel.Item {
		fmt.Printf("Feed: %v Title: %v\n", rssFeed.Channel.Title, entry.Title)
	}

	return nil
}
