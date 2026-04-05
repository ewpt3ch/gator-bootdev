package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/ewpt3ch/gator-bootdev/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("No username given")
	}

	username := cmd.arguments[0]

	user, err := s.db.GetUserByName(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Set user to %s\n", username)
	printUser(user)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("No username given")
	}

	id := uuid.New()
	timeNow := time.Now()
	username := cmd.arguments[0]
	dbParams := database.CreateUserParams{
		ID:        id,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      username,
	}

	user, err := s.db.CreateUser(context.Background(), dbParams)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Created user: %s\n", username)
	printUser(user)
	return nil
}

func handlerResetUsers(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return nil
	}

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("* %s", user)
		if s.cfg.Current_user_name == user {
			fmt.Println(" (current)")
			continue
		}
		fmt.Println()
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:				%v\n", user.ID)
	fmt.Printf(" * Name:			%v\n", user.Name)
}
