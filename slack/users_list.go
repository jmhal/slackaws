package main

import (
	"fmt"
	"github.com/bluele/slack"
)

const (
	token = "add your token"
)

func main() {
	api := slack.New(token)
	users, err := api.UsersList()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user.Id, user.Name)
	}
}
