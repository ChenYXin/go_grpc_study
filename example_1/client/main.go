package main

import (
	"context"
	"fmt"
	hello_grpc2 "go_grpc_study/example_1/grpc_proto/hello_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	client := hello_grpc2.NewHelloServiceClient(conn)
	result, err := client.SayHello(context.Background(), &hello_grpc2.HelloRequest{
		Name:    "Donkor",
		Message: "hello",
	})
	fmt.Println(result, err)
}
