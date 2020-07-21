#!/usr/bin/env bash

make

tar zcf chat.tar chat

ssh root@192.168.0.110 "mkdir -p /ivargo/app/chat/logs"

scp chat.tar root@192.168.0.110:/ivargo/app/chat
scp config.yml root@192.168.0.110:/ivargo/app/chat/

ssh root@192.168.0.110 "cd /ivargo/app/chat && tar zxf chat.tar  && supervisorctl reload"

rm -rf chat chat.tar
