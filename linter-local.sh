#!/bin/bash
set -e


GometalinterVariable=(
           "aligncheck"
           "deadcode"
           "dupl"
           "errcheck"
           "gas"
           "goconst"
           "gocyclo"

        #    "goimports"

           "golint"
           "gosimple"

        #    "gotype"
           # about gotype please see this issue https://github.com/alecthomas/gometalinter/issues/206
           # and https://github.com/alecthomas/gometalinter/issues/355
           "ineffassign"
           "interfacer"

           # "lll"
           "misspell"
           "safesql"
           "staticcheck"
           "structcheck"
           "unconvert"
           "unparam"
           "unused"
           "varcheck"
)


Directory=(
            "controllers/grpc"
            "controllers/http"
            
            "helper"
            "helper/timetn"
            
            "models/db"
            "models/db/interfaces"
            "models/db/pgsql"
            "models/logic"
            "models/stub"

            "structs"
            "structs/api/grpc"
            "structs/api/http"
            "structs/db"
            "structs/external"
            "structs/logic"

            "thirdparty/mq"
          )

arrayGometalinterVariable=${#GometalinterVariable[@]}
arrayDirectory=${#Directory[@]}


go get -u gopkg.in/alecthomas/gometalinter.v1
gometalinter.v1 --install

for ((k=0; k<${arrayDirectory}; k++));
do
        #cd ${Directory[$k]}
  for ((i=0; i<${arrayGometalinterVariable}; i++));
  do
        if [ "${Directory[$k]}" == "controllers/http" ] || [ "${Directory[$k]}" == "controllers/grpc" ] || [ "${Directory[$k]}" == "structs" ] || [ "${Directory[$k]}" == "structs/api" ] || [ "${Directory[$k]}" == "structs/external" ] || [ "${Directory[$k]}" == "structs/mongo" ] || [ "${Directory[$k]}" == "structs/pgsql" ] || [ "${Directory[$k]}" == "structs/redis" ] || [ "${Directory[$k]}" == "structs/db" ]
          then
          if [ "${GometalinterVariable[$i]}" != "gocyclo" ] && [ "${GometalinterVariable[$i]}" != "lll" ] && [ "${GometalinterVariable[$i]}" != "dupl" ] && [ "${GometalinterVariable[$i]}" != "goconst" ]
            then
            echo "Currently linter running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
            gometalinter.v1 -j 1 --disable-all  --exclude=_test --exclude=_controllers --exclude=contract_ --enable=${GometalinterVariable[$i]}  ${Directory[$k]}/  2>&1
          fi
        elif [ "${Directory[$k]}" == "routers" ] || [ "${Directory[$k]}" == "routers/grpc" ] || [ "${Directory[$k]}" == "routers/http" ]
          then
          if [ "${GometalinterVariable[$i]}" != "aligncheck" ] && [ "${GometalinterVariable[$i]}" != "errcheck" ] && [ "${GometalinterVariable[$i]}" != "gosimple" ] && [ "${GometalinterVariable[$i]}" != "interfacer" ]
            then
            echo "Currently linter running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
            gometalinter.v1 -j 1 --disable-all  --exclude=_test --exclude=_controllers --exclude=contract_ --enable=${GometalinterVariable[$i]}  ${Directory[$k]}/  2>&1
          fi
        else
          echo "Currently linter running in ${Directory[$k]} == ${GometalinterVariable[$i]}"
          gometalinter.v1 -j 1 --disable-all  --exclude=_test --exclude=_controllers --exclude=contract_ --enable=${GometalinterVariable[$i]}  ${Directory[$k]}/  2>&1
        fi

        sleep 1
        wait

  done
done