package main

import (
	"github.com/bluele/slack"
	"fmt"
)

const (
	token = "add your token"
)

func main() {

	UsersList(token);

	/*
	api := slack.New(token)
	users, err := api.UsersList()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user.Id, user.Name)
	}
	*/
}

func UsersList(workspaceToken string) {
	api := slack.New(workspaceToken)
	users, err := api.UsersList()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		if(user.Name != "slackbot" && user.Name != "slackaws") {
			fmt.Println(user.Name)
		}
	}
}
