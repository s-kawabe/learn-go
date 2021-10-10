protoc \
  --experimental_allow_proto3_optional \
  --go_out=api/hello --go_opt=paths=source_relative \
  --go-grpc_out=api/hello --go-grpc_opt=paths=source_relative \
  proto/hello.proto