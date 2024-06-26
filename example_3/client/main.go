package main

import (
	"bufio"
	"context"
	"fmt"
	"go_grpc_study/example_3/grpc_proto/stream_proto"
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

	client := stream_proto.NewServiceStreamClient(conn)
	//发送消息给服务端
	stream, err := client.DownloadFile(context.Background(), &stream_proto.Request{
		Name: "李四",
	})
	//缓冲写的方式另存为新的图片
	//os.O_CREATE ： 创建并打开一个新文件
	//os.O_TRUNC ：打开一个文件并截断它的长度为零（必须有写权限）
	//os.O_WRONLY ：以只写的方式打开
	file, err := os.OpenFile("../static/grpc_new.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	var index int
	for {
		index++
		response, errRecv := stream.Recv()
		if errRecv == io.EOF {
			break
		}
		if errRecv != nil {
			fmt.Println(errRecv)
			break
		}
		fmt.Printf("第 %d 次，写入 %d 数据\n", index, len(response.Content))
		writer.Write(response.Content)
	}
	writer.Flush()
}
