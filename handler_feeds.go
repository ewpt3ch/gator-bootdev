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

	user, err := s.db.GetUser(context.Background(), userName)
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

	fmt.Println("Created feed: %s", feedName)
	fmt.Println(feed)

	return nil
}
