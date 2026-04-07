package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ewpt3ch/gator-bootdev/internal/database"
	"github.com/google/uuid"
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
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
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

	fmt.Printf("Updating %s feed.\n", feed.Name)

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
		id := uuid.New()
		timeNow := time.Now()
		title := entry.Title
		url := entry.Link
		description := entry.Description
		pubDate, err := parsePubDate(entry.PubDate)
		if err != nil {
			return err
		}
		publishedAt := pubDate
		feed_id := feed.ID

		postParams := database.CreatePostParams{
			ID:          id,
			CreatedAt:   timeNow,
			UpdatedAt:   timeNow,
			Title:       title,
			Url:         url,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feed_id,
		}

		err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				fmt.Printf("Adding post %v\n", entry.Link)
				fmt.Printf("%v\n", err)
			}
		}
	}

	return nil
}

func parsePubDate(date string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
	}

	var parsed time.Time
	var err error
	for _, layout := range layouts {
		parsed, err = time.Parse(layout, date)
		if err == nil {
			break
		}
	}
	if err != nil {
		return parsed, err
	}

	return parsed, nil
}
