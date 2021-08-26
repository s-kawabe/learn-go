# gRPC-web sandbox

## 参考

https://qiita.com/Le0tk0k/items/af2fbf1d246f0a788ec3
https://zenn.dev/mikankitten/articles/0437fa6fb7de82

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
Protocol Buffersは構造化されたデータをシリアライズするデータフォーマットで、gRPCでは
デフォルトのフォーマットとなっておりGoogleが開発している。

### GOPATHとは
> Go1.16がリリースされたことでGo-Moduleによるプロジェクト構成が標準で推奨されることになりました。

最初はGOPATHではなくGoModuleモードでやればよいらしい。

### Goにおけるモジュールとパッケージ

- モジュール
パッケージを一つまたは複数のサブパッケージを取りまとめたカタマリ。

- パッケージ
フォルダ単位で単一ファイルまたは複数ファイルのカタマリ。

- サブパッケージ
サブフォルダにおくだけで扱いはパッケージと同等。

```
go mod init <モジュール名>
```

あるフォルダ内でこれを実行するとそのフォルダ以下が`モジュール`となる。
カレントフォルダにgo.modファイルが作成され、このカレントフォルダ以下がまるっと「モジュール」の扱いになる。

## setup

```
go version
=> go version go1.16.4 darwin/amd64
```

```
go mod init gRPC-demo/api
```

- brewでの自動生成モジュールインストール?

```
brew install protobuf protoc-gen-grpc-web protoc-gen-go
```

- go intallでGo用のプロトコルコンパイラプラグイン? →入れたけど関係ない？

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

```
protoc --version
=> libprotoc 3.17.3
```

```
npx create-next-app client --ts
```

## code generate

サーバー

これでなんかいけた
```
--go_out=api/hello --go_opt=paths=source_relative \
--go-grpc_out=api/hello --go-grpc_opt=paths=source_relative \
proto/hello.proto
```

クライアント
https://github.com/grpc/grpc-web#code-generator-plugin

Code Generator Plugin

```
$ sudo mv ~/Downloads/protoc-gen-grpc-web-1.2.1-darwin-x86_64 \
    /usr/local/bin/protoc-gen-grpc-web
$ chmod +x /usr/local/bin/protoc-gen-grpc-web
```

```
protoc --grpc-web_out=import_style=typescript,mode=grpcwebtext:client/src/api --js_out=import_style=commonjs:client/src/api -I=proto proto/hello.proto
```
→ここで権限系の注意が出るが許可する。

## error集

- Go側のコードをprotoから生成する際

```
protoc-gen-go: unable to determine Go import path for "hello.proto"

Please specify either:
        • a "go_package" option in the .proto source file, or
        • a "M" argument on the command line.
```

→ proto内にはgo_packageの指定をしなければいけないらしい
→ GOPATHを使用している場合と、していない場合(go_modules)で対処法が異なる。
