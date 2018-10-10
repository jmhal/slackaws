package slack

import (
	"github.com/bluele/slack"
)

func UsersList(workspaceToken string) ([]string) {
	api := slack.New(token)
	users, err := api.UsersList()
	if err != nil {
		panic(err)
	}
	foundUsers := make([]string, len(users))
	for i := 0; i<len(users); i++ {
		foundUsers[i] = users[i].Name
	}
	return foundUsers
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

