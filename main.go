package main

import (
	"fmt"
	"github.com/ewpt3ch/gator-bootdev/internal/config"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error getting config: %s", err)
	}
	s := state{cfg: &cfg}

	cmds := commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) == 0 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	userCommand := command{os.Args[1], os.Args[2:]}
	err = cmds.run(&s, userCommand)
	if err != nil {
		fmt.Printf("Failed to run command %s: %s", userCommand.name, err)
		os.Exit(1)
	}
}
