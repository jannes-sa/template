#!/bin/bash
export GOPATH=/var/lib/jenkins/workspace/SAV_TXN
# export GOPATH=/home/jannes/work/TN/sav-txn

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

cat $GOPATH/src/txn/conf/ci/mq.json
cat $GOPATH/src/txn/conf/ci/mongodb.json
cd $GOPATH/src/txn && $GOPATH/bin/bee migrate -driver=postgres -conn="postgres://postgres:root@172.17.0.1:5432/postgres?sslmode=disable"
cd $GOPATH/src/txn/models/logic && go test -v --cover
cd $GOPATH/src/txn/models/logic/accrue && go test -v --cover
cd $GOPATH/src/txn/models/logic/mymo-posting && go test -v --cover
cd $GOPATH/src/txn/models/logic/posting && go test -v --cover
cd $GOPATH/src/txn/models/logic/posting/interdomain && go test -v --cover
cd $GOPATH/src/txn/models/logic/prepost/interdomain && go test -v --cover
cd $GOPATH/src/txn/models/logic/batch/earning && go test -v --cover
cd $GOPATH/src/txn/models/logic/batch/deactivate && go test -v --cover
cd $GOPATH/src/txn/models/logic/batch/reverseredemption && go test -v --cover

cd $GOPATH/src/txn/routers/component/accrue && go test -v
cd $GOPATH/src/txn/routers/component/interdomain && go test -v
cd $GOPATH/src/txn/routers/component/commontransaction && go test -v
cd $GOPATH/src/txn/routers/component/loyaltyTransaction && go test -v --cover
cd $GOPATH/src/txn/models/logic/batch/loyaltyrecon && go test -v --cover
cd $GOPATH/src/txn/models/logic/batch/clearingpoint && go test -v --cover

cd $GOPATH/src/txn/routers/component/batch && go test -v
cd $GOPATH/src/txn/routers/component/reconcile && go test -v
cd $GOPATH/src/txn/routers/component/clearingpoint && go test -v
