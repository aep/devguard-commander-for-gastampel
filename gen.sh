#!/bin/sh

openapi-generator generate \
    -i api.yaml \
    -g go-gin-server  \
    -p packageName=server \
    -o  .

for i in ../carrier/proto/*.proto
do
    protoc -I ./ --go_out=protos -I../carrier/proto/ $i
done
