package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Printf("The Raw Feed: %v", rssFeed)
	fmt.Printf("User: %s and command %v", s.cfg.Current_user_name, cmd.name)
	return nil
}
