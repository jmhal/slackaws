// Pacote com as funções para o AWS.
package aws
import (
   "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func CreateInstance(Imagem string, Instancia string, NameInstancia string) {

   svc := ec2.New(session.New(&aws.Config{Region: aws.String("us-west-1")}))
   runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
      ImageId:      aws.String(Imagem),
      InstanceType: aws.String(Instancia),
      MinCount:     aws.Int64(1),
      MaxCount:     aws.Int64(1),
   })

   if err != nil {
      log.Println("Could not create instance", err)
      return
   }

   log.Println("Created instance", *runResult.Instances[0].InstanceId)
   _ , errtag := svc.CreateTags(&ec2.CreateTagsInput{
      Resources: []*string{runResult.Instances[0].InstanceId},
      Tags: []*ec2.Tag {
         {
            Key:   aws.String("Name"),
            Value: aws.String(NameInstancia),
         },
      },
   })

   if errtag != nil {
      log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
      return
   }

   log.Println("Successfully tagged instance")
}

func CreateUsers(users []string) (map[string]string) {
   // Com a instância já criada, cria os usuários
   // Talvez seja necessário uma variável global para guardar informações sobre a instância já criada.
}


