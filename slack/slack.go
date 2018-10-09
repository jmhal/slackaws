package slack

import (
	"github.com/bluele/slack"
)

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

func SendMessageToUser(name string, message string, workspaceToken string) (string) {
	api := slack.New(token)
    users, err := api.UsersList()
	if err != nil {
		panic(err)
	}
    for _, user := range users {
		if(user.Name == name) {
			api.ChatPostMessage(user.Id, message, nil)
		}
	}
   return name + ":" + message
}
