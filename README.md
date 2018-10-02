#  Como configurar o ambiente.

## Criando a área de trabalho.

Na Go, a variável GOPATH define o _workspace_, ou seja, a área de trabalho na qual você está desenvolvendo sua solucao. Por padrao, dentro do seu _workspace_, existe um diretório _bin_ e outro _src_.  Vamos comecar criando um diretorio para ser seu _workspace_.  Pode ser qualquer caminho, mas depois  precisamos atualizar a variável GOPATH.


```
$ mkdir workspace
$ export GOPATH=$PWD/workspace
```

A  variável PWD tem seu diretório corrente, onde você está executando os comandos. Dessa forma, a GOPATH estará definida com o caminho completo. Você precisa configurar a variável para toda sessao que iniciar e deve alterá-la quando passar para projetos diferentes. 

## Recuperando os pacotes.

Depois da variável configurada, pode baixar o projeto.

```
$ go get github.com/jmhal/slackaws
```
Isso irá baixar o projeto para workspace/src/github.com/jmhal/slack. Para construir a solucao:

```
$ go install github.com/jmhal/slackaws
````
No diretório workspace/bin estará o binário *slackaws*.

