#!/bin/bash
export GOENV=devci
export GOAPP=template
export GOPATH=/var/lib/jenkins/workspace/DLOR-Collect_service

# process download gometalinter
pwd

go version
go env
echo $GOENV
cd $GOPATH
go get github.com/beego/bee
# chmod +x $GOPATH/src/template/goget.sh
# $GOPATH/src/template/goget.sh
go get -u gopkg.in/alecthomas/gometalinter.v1 && $GOPATH/bin/gometalinter.v1 --install

# run linter
chmod +x $GOPATH/src/template/cicd/linter/runLinter
cd $GOPATH/src/template && ./cicd/linter/runLinter