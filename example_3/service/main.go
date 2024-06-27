package main

import (
	"context"
	"fmt"
	wake_grpc3 "go_grpc_study/example_3/grpc_proto/wake_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

// 新版本 gRPC 要求必须嵌入 UnimplementedGreeterServer 结构体
type VoiceWakeServer struct {
	wake_grpc3.UnimplementedVoiceWakeServiceServer
}
type FaceWakeServer struct {
	wake_grpc3.UnimplementedFaceWakeServiceServer
}

func (VoiceWakeServer) DogBark(ctx context.Context, request *wake_grpc3.Request) (pd *wake_grpc3.Response, err error) {
	fmt.Println("语音唤醒入参：", request.Name)
	pd = new(wake_grpc3.Response)
	pd.Sound = "汪汪汪～"
	return
}

func (FaceWakeServer) ASlap(ctx context.Context, request *wake_grpc3.Request) (pd *wake_grpc3.Response, err error) {
	fmt.Println("人脸唤醒入参：", request.Name)
	pd = new(wake_grpc3.Response)
	pd.Sound = "塞班～"
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
	wake_grpc3.RegisterVoiceWakeServiceServer(s, &VoiceWakeServer{})
	wake_grpc3.RegisterFaceWakeServiceServer(s, &FaceWakeServer{})
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
