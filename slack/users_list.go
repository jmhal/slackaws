package main

import (
	"github.com/bluele/slack"
	"fmt"
)

const (
	token = "xoxb-424210374210-425256236567-2ZNtzcbyZ4wNnvSuNucziZK5"
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
