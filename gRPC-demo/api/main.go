package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/s-kawabe/learn-go/gRPC-demo/api/hello/proto"

	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

// Helloメソッドを実装することでprotoファイルで定義した内容と連携できます
func (*server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Getxxxメソッドも自動に作成されています。
	log.Printf("Recieved : %s", in.GetName())
	name := in.GetName()
	email := in.GetEmail()
	age := in.GetAge()
	log.Println(name, email, age)
	result := fmt.Sprintf("Hello,%s (%d). email is %s", name, age, email)

	// 受け取ったメッセージを連結したレスポンスを返します。
	return &pb.HelloResponse{
		Message: result,
	}, nil
}

func main() {
	// リッスンするportをを設定します
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println(lis)
	var opts []grpc.ServerOption

	// サーバをインスタンス化します。
	s := grpc.NewServer(opts...)

	// RegisterXXXXメソッドも自動的に作成されています。
	pb.RegisterHelloServiceServer(s, &server{})

	// 起動確認様にログ出力をさせてみます
	log.Println("Waiting for gRPC request ....")
	log.Println("--------------")

	// サービスの起動を行います。
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
