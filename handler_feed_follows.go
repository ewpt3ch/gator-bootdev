package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ewpt3ch/gator-bootdev/internal/database"
	"github.com/google/uuid"
)

func handlerCreateFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Expected a singer argument: <url>")
	}

	url := cmd.arguments[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	id := uuid.New()
	timeNow := time.Now()
	user_id := user.ID
	feed_id := feed.ID

	FeedFollowParams := database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    user_id,
		FeedID:    feed_id,
	}

	result, err := s.db.CreateFeedFollow(context.Background(), FeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println(result.UserName)
	fmt.Println(result.FeedName)
	return nil
}

func handlerGetFeedFollowsForUser(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}
