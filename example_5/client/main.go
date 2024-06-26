package main

import (
	"context"
	"fmt"
	"go_grpc_study/example_5/grpc_proto/stream_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
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
	client := stream_proto.NewClientStreamClient(conn)
	stream, err := client.UploadFile(context.Background())

	//分片读的方式读取图片
	file, err := os.Open("../static/grpc.png")
	if err != nil {
		log.Fatalf(fmt.Sprintf("open file err [%s]", err))
	}

	defer file.Close()
	for {
		buf := make([]byte, 1024)
		_, err = file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		stream.Send(&stream_proto.FileRequest{
			Content: buf,
		})
	}

	response, err := stream.CloseAndRecv()
	fmt.Println(response, err)
}
