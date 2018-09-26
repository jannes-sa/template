#!/bin/bash

export GOENV=devci
export GOAPP=txn
export RPC=ci
export APP_DIRECTORY=/go/src/txn
export CRED_MONGODB=mongodb://172.17.0.1:27017
export CRED_MQ=amqp://guest:guest@172.17.0.1:5672/
export CRED_REDIS=172.17.0.1:6379
export CRED_PGSQL=postgres://postgres:root@172.17.0.1:5432/postgres?sslmode=disable
export BANK_CODE=11
export BANK_NAME=PERMATA
export DOMAIN=SAV
export SMITHBANKCODE=009
export GOPATH=/var/lib/jenkins/workspace/SAV_TXN

# process download gometalinter
pwd

go version
go env
echo $GOENV
cd $GOPATH
go get github.com/beego/bee
# chmod +x $GOPATH/src/txn/goget.sh
# $GOPATH/src/txn/goget.sh
go get -u gopkg.in/alecthomas/gometalinter.v1 && $GOPATH/bin/gometalinter.v1 --install

# run linter
chmod +x $GOPATH/src/txn/linter.sh
cd $GOPATH/src/txn && ./linter.sh