package main

import (
	"context"
	"fmt"
	hello_grpc2 "go_grpc_study/example_1/grpc_proto/hello_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

// 得有一个结构体，需要实现这个服务的全部方法,叫什么名字不重要
// 新版本 gRPC 要求必须嵌入 UnimplementedGreeterServer 结构体
type HelloServer struct {
	hello_grpc2.UnimplementedHelloServiceServer
}

func (HelloServer) SayHello(ctx context.Context, request *hello_grpc2.HelloRequest) (pd *hello_grpc2.HelloResponse, err error) {
	fmt.Println("入参：", request.Name, request.Message)
	pd = new(hello_grpc2.HelloResponse)
	pd.Name = "你好"
	pd.Message = "ok"
	return
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	// 将server结构体注册为gRPC服务。
	hello_grpc2.RegisterHelloServiceServer(s, &HelloServer{})
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
