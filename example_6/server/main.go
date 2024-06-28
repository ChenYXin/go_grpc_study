package main

import (
	"fmt"
	"go_grpc_study/example_6/grpc_proto/each_proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 新版本 gRPC 要求必须嵌入 UnimplementedGreeterServer 结构体
type EachStream struct {
	each_proto.UnimplementedEachStreamServer
}

func (EachStream) Chat(stream each_proto.EachStream_ChatServer) error {
	for i := 0; i < 5; i++ {
		request, _ := stream.Recv()
		fmt.Println(request)
		stream.Send(&each_proto.Response{
			Message: fmt.Sprintf("第 %d 次回应你好", i+1),
		})
	}
	return nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个gRPC服务器实例。
	server := grpc.NewServer()
	// 将server结构体注册为gRPC服务。
	each_proto.RegisterEachStreamServer(server, &EachStream{})
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	server.Serve(listen)
}
