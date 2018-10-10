// Slackaws: integrando AWS e Slack
package main
import (
   "fmt"
   "os"
   "github.com/jmhal/slackaws/slack"
//   "github.com/jmhal/slackaws/aws"
)

func main() {
   // Na primeira versão, vamos começar passando dois parâmetros:
   // - a url do grupo no slack (por exemplo, orientadosjm.slack.com)
   // - o token de workspace da API para o mesmo grupo
   // slackWorkspaceURL := os.Args[1]
   workspaceToken := os.Args[1]

   // Recupera uma slice com todos os usuários
   users := slack.UsersList(workspaceToken)
   fmt.Println(users)

   slack.SendMessageToUser("joao.marcelo", "Teste direto do projeto", workspaceToken)

   // Voltaremos a ativar a sequência abaixo quando finalizarmos o código AWS.
   // Recupera a URL de acesso público da instância e um map[nomedousuario->chave]
   // publicUrl := aws.CreateInstance()
   /// sshKeys := aws.CreateUsers(users)

   // Manda uma mensagem para cada usuário com o conteúdo da chave
   // for user, key := range sshKeys {
   //   slack.SendMessage(user, key)
   //}
}
