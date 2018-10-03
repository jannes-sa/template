#!/bin/bash
mkdir -p /var/lib/jenkins/workspace/DLOR-Collect_service/src/template

rm -R /var/lib/jenkins/workspace/DLOR-Collect_service/src/template/*

cp -r /var/lib/jenkins/workspace/DLOR-Collect_Prep/* /var/lib/jenkins/workspace/DLOR-Collect_service/src/template

docker start mongo
docker start rabbitmq
docker start spostgres