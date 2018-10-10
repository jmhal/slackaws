package slack

// Aqui estou importando o framework com o nome "bluele" para não entrar em conflito 
// com o nosso próprio pacote que já tem o nome de "slack".
import (
   bluele "github.com/bluele/slack"
)

func UsersList(token string) ([]string) {
   api := bluele.New(token)
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

func SendMessageToUser(name string, message string, token string) (string) {
   api := bluele.New(token)
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

