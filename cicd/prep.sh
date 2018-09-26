#!/bin/bash
mkdir -p /var/lib/jenkins/workspace/SAV_TXN/src/txn

rm -R /var/lib/jenkins/workspace/SAV_TXN/src/txn/*

cp -r /var/lib/jenkins/workspace/SAV_TXN_Prep/* /var/lib/jenkins/workspace/SAV_TXN/src/txn

docker start mongo
docker start rabbitmq
docker start spostgres