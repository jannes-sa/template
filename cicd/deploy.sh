#!/bin/bash

export USER_SSH=ubuntu
export HOST_PORT_TEST=32004
export KUBE_PORT=30004
export NODE_PORT=8084
export GRPC_PORT=32104
export DOCKER_USERNAME=fakename
export DOCKER_PASSWORD=fakename
export VERSION_IMG=v1stag
export VERSION_MINOR_IMG=0
export OURAPPNAME=template
export GOPATH=/var/lib/jenkins/workspace/DLOR-Collect_service

# deploy staging
kubectl get pod
cd $GOPATH/src/template/
chmod +x ./deploy.sh
whoami
./deploy.sh $DOCKER_USERNAME $DOCKER_PASSWORD $VERSION_IMG $BUILD_ID $OURAPPNAME $VERSION_MINOR_IMG $NODE_PORT $KUBE_PORT $HOST_PORT_TEST $GRPC_PORT