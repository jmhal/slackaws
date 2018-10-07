package main

import (
	"fmt"
	"github.com/bluele/slack"
)

const (
	token = "xoxb-424210374210-425256236567-2ZNtzcbyZ4wNnvSuNucziZK5"
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
