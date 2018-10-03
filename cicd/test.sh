#!/bin/bash
export GOPATH=/var/lib/jenkins/workspace/DLOR_Collect
# export GOPATH=/home/tnis/works/TN/dlor
export WORKDIR=$GOPATH/src/template

export GOENV=devci
export GOAPP=template
export RPC=ci
export APP_DIRECTORY=/go/src/template
export CRED_MONGODB=mongodb://172.17.0.1:27017
export CRED_MQ=amqp://guest:guest@172.17.0.1:5672/
export CRED_REDIS=172.17.0.1:6379
export CRED_PGSQL=postgres://postgres:root@172.17.0.1:5432/postgres?sslmode=disable
export BANK_CODE=11
export BANK_NAME=PERMATA
export DOMAIN=SAV
export SMITHBANKCODE=009

export PATH_BATCH_EARNING=storages/batch/earning/
export DOMAIN_ID=02
export DOMAIN_ID_FIX=05
export NUM_ROUTINE=50
export DB_MAXIDLE=50
export DB_MAXCONN=300

export batch=0

export FIX_TXN_DOMAIN_HOST=http://127.0.0.1:7084
export GL_DOMAIN_HOST=http://127.0.0.1:8080
export RPC_RULESTXN=general_transaction_rule@172.17.0.1:58083
export HTTP_TXN=https://txnpost/
export HTTP_SELFSVC=http://127.0.0.1:8084
export HTTP_OTHER=http://127.0.0.1:8084
export MQ_ADDRESS_GL=amqp://guest:guest@127.0.0.1:5672/
export PATH_BATCH_LOYALTY_RECON=storages/batch/loyaltyreconcile/
export STATIC_PATH=/sav_txn/v1/storages
export PATH_BATCH_DEACTIVATE=storages/batch/deactivate/
export PATH_BATCH_REVERSE_REDEMPTION=storages/batch/reverseredemption/
export PATH_BATCH_CLEARING_POINT=storages/batch/clearingpoint/


cat $GOPATH/src/template/conf/ci/mq.json
cat $GOPATH/src/template/conf/ci/mongodb.json

cd $GOPATH/src/template && $GOPATH/bin/bee migrate -driver=postgres -conn="postgres://postgres:root@172.17.0.1:5432/postgres?sslmode=disable"

# Unit test 
cd $WORKDIR &&
go test -v --cover \
./models/logic/cards/... \
-coverprofile=$WORKDIR/sonarqube-report/coverage-report.out

cd $WORKDIR &&
go test -v --cover \
./models/logic/cards/... \
-json > $WORKDIR/sonarqube-report/unit-report.json

# # Component test
# cd $GOPATH/src/template/routers/component && 
# go test -v \
# ./accrue/... \
# ./clearingpoint/...

# Run SonarQube
cd $WORKDIR &&
docker run --rm \
    -v $(pwd):$WORKDIR \
    -w=$WORKDIR --network=sonar \
    nikhuber/sonar-scanner:latest sonar-scanner


# docker run --rm \
#     -v $(pwd):/var/lib/jenkins/workspace/DLOR_Collect/src/template \
#     -w=/var/lib/jenkins/workspace/DLOR_Collect/src/template --network=sonar \
#     nikhuber/sonar-scanner:latest sonar-scanner