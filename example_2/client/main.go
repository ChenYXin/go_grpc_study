package main

import (
	"context"
	"fmt"
	wake_grpc2 "go_grpc_study/example_2/grpc_proto/wake_proto"
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

	voiceClient := wake_grpc2.NewVoiceWakeServiceClient(conn)
	res, err := voiceClient.DogBark(context.Background(), &wake_grpc2.Request{
		Name: "张三",
	})
	fmt.Println(res, err)

	faceClient := wake_grpc2.NewFaceWakeServiceClient(conn)
	res, err = faceClient.ASlap(context.Background(), &wake_grpc2.Request{
		Name: "李四",
	})
	fmt.Println(res, err)

}
