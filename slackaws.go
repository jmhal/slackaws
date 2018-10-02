package main
import (
   "fmt"
   "os"
   "github.com/jmhal/slackaws/slack"
)

func main() {
   slackURL := os.Args[1]
   fmt.Println(slackURL)
   users := slack.UsersList(slackURL)
   fmt.Println(users)
   status := slack.SendMessageToUser("joaomarcelo", "teste")
   fmt.Println(status)
}
