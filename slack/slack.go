package slack
//Alterei
func UsersList(slackURL string) ([]string) {
   users := []string{slackURL}
   return users
}

func SendMessageToUser(user string, message string) (string) {
    return user + ":" + message
}
