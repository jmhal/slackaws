#!/bin/bash
# Exemplo de invocação da API através do comando curl

# TOKEN do Workspace gerado pela aplicação.
TOKEN=$1

# Testa a autenticação com o TOKEN fornecido
# https://api.slack.com/methods/auth.test
STATUS=$(curl -sS -X POST -H "Authorization: Bearer $TOKEN" -H 'Content-type: application/json' https://slack.com/api/auth.test | tee auth.log | grep -E '(\"ok\")' | cut -f1 -d',' | cut -f2 -d':')
if [ "$STATUS" == "true" ]
then
   echo "Autenticação com Sucesso."
   rm auth.log
else
   echo "Problemas na autenticação."
   cat auth.log | python -m json.tool
   rm auth.log
   exit 1
fi

# Listar usuários
# https://api.slack.com/methods/users.list
curl -sS -X GET -H "Authorization: Bearer $TOKEN" -H 'Content-type: application/json' https://slack.com/api/users.list | python -m json.tool > users_output.log
OLDIFS=$IFS
IFS=','
declare -A users
ID=""
for input in `cat users_output.log | grep -E '(\"name\"|\"id\")'`
do 
   inputtype=$(echo "$input" | sed 's/"//g' | cut -f1 -d':' | tr -d '[[:space:]]' )
   inputvalue=$(echo "$input" | sed 's/"//g' | cut -f2 -d':' | tr -d '[[:space:]] ')
   if [ "$inputtype" == "id" ]
   then
      ID=$inputvalue     
   else
      users[$inputvalue]=$ID
   fi
done
rm users_output.log
IFS=$OLDIFS

echo "Usuários do Grupo:"
for key in "${!users[@]}" 
do
   if [ ! "$key" == "slackbot" ]
   then
      echo "$key"
   fi
done

# Envia mensagem para um usuário
# https://api.slack.com/methods/chat.postMessage
read -p "Informe o usuário que deseja enviar mensagem: " USER
read -p "Informe a mensagem: " MENSAGEM
curl -sS -X POST -H "Authorization: Bearer $TOKEN" -H 'Content-type: application/json' \
--data "{\"channel\":\"${users[$USER]}\", \"text\":\"$MENSAGEM\" }" \
https://slack.com/api/chat.postMessage | python -m json.tool > message.log 
STATUS=$(cat message.log | grep -E '(\"ok\")' | cut -f1 -d',' | cut -f2 -d':' | tr -d '[[:space:]]')
if [ "$STATUS" == "true" ]
then
   echo "Mensagem enviada com Sucesso."
   rm message.log
else
   echo "Problemas no envio da mensagem."
   cat message.log | python -m json.tool
   rm message.log
   exit 1
fi


