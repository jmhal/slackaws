// Pacote com as funções para o AWS.
package aws
import (
   "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func CreateInstance(instanceImage string, instanceType string, instanceName string, region string, keyName string) {

   svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))
   runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
      ImageId:      aws.String(instanceImage),
      InstanceType: aws.String(instanceType),
      MinCount:     aws.Int64(1),
      MaxCount:     aws.Int64(1),
      KeyName:      aws.String(keyName),
      NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification {&ec2.InstanceNetworkInterfaceSpecification{
         DeviceIndex:              aws.Int64(0),
	 AssociatePublicIpAddress: aws.Bool(true),
      }},
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
            Value: aws.String(instanceName),
         },
      },
   })

   if errtag != nil {
      log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
      return
   }
   log.Println("Successfully tagged instance")

   instanceId := *runResult.Instances[0].InstanceId
   state := "pending"
   publicDns := ""
   for state != "running" {
      result, err := svc.DescribeInstances(nil)
      if err != nil {
         log.Println(err)
      }
      for _, reservation := range result.Reservations {
         if (instanceId == *reservation.Instances[0].InstanceId) && (*reservation.Instances[0].State.Name == "running") {
	    state = "running"
	    publicDns = *reservation.Instances[0].PublicDnsName
	 }
      }
   }
   log.Println(publicDns)
}
/*
func CreateUsers(users []string) (map[string]string) {
   // Com a instância já criada, cria os usuários
   // Talvez seja necessário uma variável global para guardar informações sobre a instância já criada.
  
}*/


