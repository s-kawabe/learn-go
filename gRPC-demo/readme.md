# gRPC-web sandbox

## introduction

[introduction gRPC](https://grpc.io/docs/what-is-grpc/introduction/)

[go quiqstart](https://grpc.io/docs/languages/go/)

### RPCとは
RemoteProcedureCall
他サーバー上の関数を呼び出し、実行する事を可能にする手法。
また、同じRPCの仕様に準拠していれば呼び出される関数は別の言語で実装されていても問題ない。

### gRPCの概要
gRPCクライアントとサーバーは、Google内のサーバーから独自のデスクトップまで、さまざまな環境で
実行および通信でき、サポートされている任意の言語で記述できる。
- ProtocolBuffersを利用して高速かつ低容量での通信を実現
- 多言語に対応
- HTTP/2でデータを送受信する
- IDLからサーバとクライアントのコードを生成できる
- １対１、多対１、多対多など様々なバリエーションの通信方法が実現可能

### Protocol Buffers
Protocol Buffersは構造化されたデータをシリアライズするデータフォーマットで、gRPCdeha
デフォルトのフォーマットとなっておりGoogleが開発している。

## setup

```
go version
=> go version go1.16.4 darwin/amd64
```

```
go mod init gRPC-demo/api
```

```
brew install protobuf protoc-gen-grpc-web protoc-gen-go
```

```
protoc --version
=> libprotoc 3.17.3
```

```
npx create-next-app client --ts
```