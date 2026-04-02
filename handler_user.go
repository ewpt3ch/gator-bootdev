package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No username given")
	}

	username := cmd.arguments[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Set user to %s", username)
	return nil
}
