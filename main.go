package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ewpt3ch/gator-bootdev/internal/config"
	"github.com/ewpt3ch/gator-bootdev/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error getting config: %s", err)
	}

	dbURL := cfg.DB_url
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	s := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
