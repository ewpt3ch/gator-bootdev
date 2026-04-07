# Feed aggregator
Requires:
- postgres
- go

Usage: gator <command> <arguments>

 Create a config file called .gatorconfig.json in your home directory. It contains
```
# ~/.gatorconfig.json
{
    "bd_url":postgres://username:password@localost:5432/gator?sslmode=disable",
}
```


Install and initial setup with:
```
go install gator
go install goose
find sql/schema and run goose <postgres database string>
```

Commands:
- register: registers a user <user>
- login: change user <user>
- reset: deletes all users from the database and all feeds and posts
- users: lists all registered users
- agg: loop that checks feeds and stores new posts <refresh time ex: 5m>
- addfeed: addsa feed <name> <url>
- follow: follow a feed already in the database
- browse: show all posts followed by current user
