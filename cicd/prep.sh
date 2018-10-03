#!/bin/bash
mkdir -p /var/lib/jenkins/workspace/DLOR_Collect/src/template

rm -R /var/lib/jenkins/workspace/DLOR_Collect/src/template/*

cp -r /var/lib/jenkins/workspace/DLOR_Collect_Prep/* /var/lib/jenkins/workspace/DLOR_Collect/src/template

docker start mongo
docker start rabbitmq
docker start spostgres