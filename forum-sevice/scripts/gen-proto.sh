#!/bin/bash
CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/proto-service/genprotos
for x in $(find ${CURRENT_DIR}/proto-service/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/proto-service/forum-protos -I /usr/local/go --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done
