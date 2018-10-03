#  Como configurar o ambiente.

Abaixo instruções de como configurar o ambiente para poder trabalhar na solução.

## Criando a área de trabalho.

Na _Go_, a variável GOPATH define o _workspace_, ou seja, a área de trabalho na qual você está desenvolvendo sua solução. Para cada projeto desenvolvido na linguagem _Go_ deve ser criado um _workspace_ diferente. O diretório não precisa ter um nome específico ou obedecer qualquer regra de nomenclatura. 

Por padrão, dentro do seu _workspace_, há um diretório _bin_ e outro _src_.  Vamos começar criando um diretório para ser seu _workspace_. Como já foi dito, pode ser qualquer caminho, mas depois precisamos atualizar a variável GOPATH.

```
$ mkdir workspace
$ export GOPATH=$PWD/workspace
```

A  variável PWD tem seu diretório corrente, onde você está executando os comandos. Dessa forma, a GOPATH estará definida com o caminho completo (por exemplo, se seu diretório corrente for _/home/usuario_, a variável vai conter o valor _/home/usuario/workspace_). Você precisa configurar a variável para toda sessão que iniciar e deve alterá-la quando passar para projetos diferentes ou reiniciar o sistema. 

## Recuperando os pacotes.

Depois da variável configurada, no *mesmo terminal*, pode baixar o projeto.

```
$ go get github.com/jmhal/slackaws
```

Isso irá baixar o projeto para workspace/src/github.com/jmhal/slack. Agora, devemos construir a solução. Isso significa compilar o código do pacote _main_ que contém a função _main_. Para construir a solução:

```
$ go install github.com/jmhal/slackaws
````

Lembrando que deve ser executado no terminal no qual a GOPATH foi configurada. No diretório _workspace/bin_ estará o binário *slackaws*. Veja que não precisamos informar o caminho completo (_/home/usuario/workspace/src/github.com/jmhal/slackaws_). Com a variável configurada, basta mencionar _github/jmhal/slackaws_).

