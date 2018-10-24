// Slackaws: integrando AWS e Slack
package main
import (
   "log"
   "net"
   "time"
   "os"
   "github.com/jmhal/slackaws/slack"
   "github.com/jmhal/slackaws/aws"
)

func main() {
   // Descrição dos parâmetros:
   // - o token de workspace da API para o mesmo grupo
   // - identificador da imagem a ser usada (depende da Região da Amazon)
   // - tipo da instância (t2.micro, t1.small, etc)
   // - nome da instância (ser usado como tag)
   // - região da AWS a ser usada
   // - nome da chave a ser usada (supõe-se que existe como arquivo .pem no diretório da execução)
   workspaceToken := os.Args[1]
   imageId := os.Args[2]
   instanceType := os.Args[3]
   instanceName := os.Args[4]
   region := os.Args[5]
   key := os.Args[6]

   // Cria a instância na nuvem.
   publicDns := aws.CreateInstance(imageId, instanceType, instanceName, region, key)
   log.Println("Instance Created with Address: " + publicDns)

   // Verifica se o servidor SSH já está acessível.
   conn, err := net.Dial("tcp", publicDns + ":22")
   for err != nil {
      log.Println("Esperando o estabelecimento do Servidor SSH.")
      time.Sleep(5 * time.Second)
      conn, err = net.Dial("tcp", publicDns + ":22")
   }
   log.Println("Servidor SSH Estabelecido.")
   log.Printf("%s\n", conn)

   // Recupera uma slice com todos os usuários
   users := slack.UsersList(workspaceToken)
   log.Printf("Available Users: %s\n", users)

   // Cria os usuários na instância
   aws.CreateUsers(publicDns, key + ".pem", users)

   // Manda uma mensagem para cada usuário com a chave
   slack.SendKeys(workspaceToken)

   // Envia mensagem para os usuários
   for _, usuario := range users {
      if usuario == "slackbot" {
         continue
      }
      slack.SendMessageToUser(usuario, "Para se conectar no servidor use o seguinte comando:", workspaceToken)
      nomeusuario := aws.RemoveSpecialChars(usuario)
      slack.SendMessageToUser(usuario, "ssh -i " + nomeusuario + ".pem " + nomeusuario + "@" + publicDns, workspaceToken)
   }

   return;
}
