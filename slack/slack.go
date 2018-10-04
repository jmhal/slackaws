package slack

func UsersList(slackURL string, workspaceToken string) ([]string) {
   // Esta função deve ser implementada. Deve acessar a API do Slack e retornar uma lista de usuários do grupo.
   users := []string{slackURL, workspaceToken}
   return users
}

func SendMessageToUser(user string, message string) (string) {
   // Envia uma mensagem para o usuário.
   return user + ":" + message
}
