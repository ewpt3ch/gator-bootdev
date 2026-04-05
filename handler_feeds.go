package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ewpt3ch/gator-bootdev/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Expected 2 arguments: <feed name> <url>")
	}

	userName := s.cfg.Current_user_name

	user, err := s.db.GetUserByName(context.Background(), userName)
	if err != nil {
		return err
	}

	id := uuid.New()
	timeNow := time.Now()
	feedName := cmd.arguments[0]
	url := cmd.arguments[1]
	user_id := user.ID
	dbParams := database.CreateFeedParams{
		ID:        id,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      feedName,
		Url:       url,
		UserID:    user_id,
	}

	feed, err := s.db.CreateFeed(context.Background(), dbParams)
	if err != nil {
		return err
	}

	fmt.Printf("Created feed: %s\n", feedName)
	printFeed(feed)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {

		userid := feed.UserID
		user, err := s.db.GetUserByID(context.Background(), userid)
		if err != nil {
			return err
		}

		fmt.Printf("* Feed name:				%s\n", feed.Name)
		fmt.Printf("* Feed URL:					%s\n", feed.Url)
		fmt.Printf("* User Name:				%s\n", user.Name)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:							%s\n", feed.ID)
	fmt.Printf("* Created:				%v\n", feed.CreatedAt)
	fmt.Printf("* Updated					%v\n", feed.UpdatedAt)
	fmt.Printf("* Name:						%s\n", feed.Name)
	fmt.Printf("* URL:						%s\n", feed.Url)
	fmt.Printf("* UserID:					%s\n", feed.UserID)
}
