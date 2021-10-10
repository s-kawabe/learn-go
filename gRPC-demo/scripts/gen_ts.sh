#!/bin/bash

protoc \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:client/src/api \
  --js_out=import_style=commonjs:client/src/api \
  -I=proto proto/hello.proto