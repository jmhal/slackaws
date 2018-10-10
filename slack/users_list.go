package main

import (
	"github.com/bluele/slack"
	"fmt"
)

const (
	token = "xoxb-424210374210-425256236567-2ZNtzcbyZ4wNnvSuNucziZK5"
)

func main() {
	use := Lista(token)
	for i := 0; i<len(use); i++ {
		fmt.Println(use[i])
	}
}
func Lista(workspaceToken string) ([]string) {
	api := slack.New(token)
	users, err := api.UsersList()
	if err != nil {
		panic(err)
	}
	usuarios := make([]string, len(users))
	for i := 0; i<len(users); i++ {
		usuarios[i] = users[i].Name
	}
	return usuarios
}
