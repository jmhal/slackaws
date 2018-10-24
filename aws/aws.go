// Pacote com as funções para o AWS.
package aws
import (
   "log"
   "fmt"
   "io"
   "io/ioutil"
   "net"
   "os"
   "strings"

   "github.com/aws/aws-sdk-go/aws"
   "github.com/aws/aws-sdk-go/aws/session"
   "github.com/aws/aws-sdk-go/service/ec2"
   "github.com/bramvdbogaerde/go-scp"
   "github.com/bramvdbogaerde/go-scp/auth"

   "golang.org/x/crypto/ssh"
   "golang.org/x/crypto/ssh/agent"
)

type SSHCommand struct {
   Path   string
   Env    []string
   Stdin  io.Reader
   Stdout io.Writer
   Stderr io.Writer
}

type SSHClient struct {
   Config *ssh.ClientConfig
   Host   string
   Port   int
}

func (client *SSHClient) RunCommand(cmd *SSHCommand) error {
   var (
      session *ssh.Session
      err     error
   )

   if session, err = client.newSession(); err != nil {
      return err
   }
   defer session.Close()

   if err = client.prepareCommand(session, cmd); err != nil {
      return err
   }

   err = session.Run(cmd.Path)
   return err
}

func (client *SSHClient) prepareCommand(session *ssh.Session, cmd *SSHCommand) error {
   for _, env := range cmd.Env {
      variable := strings.Split(env, "=")
      if len(variable) != 2 {
         continue
      }

      if err := session.Setenv(variable[0], variable[1]); err != nil {
	 return err
      }
   }

   if cmd.Stdin != nil {
      stdin, err := session.StdinPipe()
      if err != nil {
         return fmt.Errorf("Unable to setup stdin for session: %v", err)
      }
      go io.Copy(stdin, cmd.Stdin)
   }

   if cmd.Stdout != nil {
      stdout, err := session.StdoutPipe()
      if err != nil {
         return fmt.Errorf("Unable to setup stdout for session: %v", err)
      }
      go io.Copy(cmd.Stdout, stdout)
   }

   if cmd.Stderr != nil {
      stderr, err := session.StderrPipe()
      if err != nil {
         return fmt.Errorf("Unable to setup stderr for session: %v", err)
      }
      go io.Copy(cmd.Stderr, stderr)
   }
   return nil
}

func (client *SSHClient) newSession() (*ssh.Session, error) {
   connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
   if err != nil {
      return nil, fmt.Errorf("Failed to dial: %s", err)
   }

   session, err := connection.NewSession()
   if err != nil {
      return nil, fmt.Errorf("Failed to create session: %s", err)
   }

   modes := ssh.TerminalModes{
      // ssh.ECHO:          0,     // disable echoing
      ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
      ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
   }

   if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
      session.Close()
      return nil, fmt.Errorf("request for pseudo terminal failed: %s", err)
   }

   return session, nil
}

func PublicKeyFile(file string) ssh.AuthMethod {
   buffer, err := ioutil.ReadFile(file)
   if err != nil {
      return nil
   }

   key, err := ssh.ParsePrivateKey(buffer)
   if err != nil {
      return nil
   }
   return ssh.PublicKeys(key)
}

func SSHAgent() ssh.AuthMethod {
   if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
      return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
   }
   return nil
}

func SendScriptToInstance(key string, url string){
    clientSCPConfig, _ := auth.PrivateKey("ubuntu", key, ssh.InsecureIgnoreHostKey())
    clientSCP := scp.NewClient(url+":22", &clientSCPConfig)

    err := clientSCP.Connect()
    if err != nil {
        fmt.Println("Couldn't establisch a connection to the remote server ", err)
        return
    }
    // Open a file
    f, _ := os.Open("./createUsers.sh")
    // Close client connection after the file has been copied
    defer clientSCP.Close()
    // Close the file after it has been copied
    defer f.Close()
    // Finaly, copy the file over
    // Usage: CopyFile(fileReader, remotePath, permission)
    clientSCP.CopyFile(f, "/home/ubuntu/createUsers.sh", "0655")
}

func CreateUsers(url string, key string, usuarios []string) {
   SendScriptToInstance(key, url);

   sshConfig := &ssh.ClientConfig{
      User: "ubuntu",
      HostKeyCallback: ssh.InsecureIgnoreHostKey(),
      Auth: []ssh.AuthMethod {
         PublicKeyFile(key),
      },
   }

   client := &SSHClient {
      Config: sshConfig,
      Host:   url,
      Port:   22,
   }

   for _, usuario := range usuarios {
      file, err := os.Create("./" + usuario + ".pem")
      if err != nil {
         fmt.Println(err)
      }
      log.Printf("Creating user: %s\n", usuario)
      cmd := &SSHCommand{
         Path:   "sudo ./createUsers.sh " + usuario,
         Env:    []string{"LC_DIR=/"},
         Stdin:  os.Stdin,
         Stdout: file,
         Stderr: os.Stderr,
      }

      if err := client.RunCommand(cmd); err != nil {
         fmt.Fprintf(os.Stderr, "command run error: %s\n", err)
         os.Exit(1)
      }
      file.Sync()
      file.Close()
   }
}

func CreateInstance(instanceImage string, instanceType string, instanceName string, region string, keyName string) (string) {
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
      return "noinstance"
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
      return "noinstance"
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
   return publicDns
}



