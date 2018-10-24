#!/bin/bash
usuario=$1

adduser --quiet --disabled-password --force-badname --gecos "" ${usuario}
ssh-keygen -q -t rsa -N "" -f ${usuario}.pem
mkdir /home/${usuario}/.ssh
cat ${usuario}.pem.pub > /home/${usuario}/.ssh/authorized_keys
chown -R ${usuario}.${usuario} /home/${usuario}/.ssh/
chmod 0700 /home/${usuario}/.ssh
chmod 0600 /home/${usuario}/.ssh/authorized_keys
cat ${usuario}.pem
