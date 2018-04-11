#!/bin/bash
for file in `\find . -name '*.proto'`; do
    temp="$temp $file"
    echo $file
done
protoc --go_out=plugins=grpc:./client $temp