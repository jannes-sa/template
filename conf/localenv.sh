#!/bin/bash
export GOENV=local
export GOAPP=template
export CRED_PGSQL=postgres://postgres:root@127.0.0.1:5432/template?sslmode=disable
export CRED_MONGO=mongodb://172.17.0.1:27017
export AUTH=0
export AUTHKEY=AUTHKEYDATA
export AUTHEXP=60
export TZ=Asia/Bangkok

export DBMAXCONN=50
export DBMAXIDLE=50

export DEBUG=1

export VERSION=v1