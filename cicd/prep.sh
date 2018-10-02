#!/bin/bash
mkdir -p /var/lib/jenkins/workspace/DLOR-Loan/src/template

rm -R /var/lib/jenkins/workspace/DLOR-Loan/src/template/*

cp -r /var/lib/jenkins/workspace/DLOR-Loan_Prep/* /var/lib/jenkins/workspace/DLOR-Loan/src/template

docker start mongo
docker start rabbitmq
docker start spostgres