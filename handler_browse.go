package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ewpt3ch/gator-bootdev/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.arguments) == 1 {
		parsedNum, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return nil
		}
		limit = int32(parsedNum)
	}

	getPostParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), getPostParams)
	if err != nil {
		return err
	}

	for _, post := range posts {
		printPost(post)
		if err != nil {
			return err
		}
	}

	return nil
}

func printPost(post database.Post) {
	fmt.Printf("Title:	%v\n", post.Title)
	fmt.Printf("URL: %v\n", post.Url)
	fmt.Printf("Published on: %v\n", post.PublishedAt)
	fmt.Printf("Description: %v\n", post.Description)
}
