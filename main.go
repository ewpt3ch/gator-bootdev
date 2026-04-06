package main

import (
	"database/sql"
	"log"
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
		log.Fatalf("Error getting config: %v", err)
	}

	dbURL := cfg.DB_url
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerResetUsers)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", middlewareLoggedIn(handlerGetFeeds))
	cmds.register("follow", middlewareLoggedIn(handlerCreateFeedFollow))
	cmds.register("following", middlewareLoggedIn(handlerGetFeedFollowsForUser))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	userCommand := command{os.Args[1], os.Args[2:]}
	err = cmds.run(&programState, userCommand)
	if err != nil {
		log.Fatal(err)
	}
}
