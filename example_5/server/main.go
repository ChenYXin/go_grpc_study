package main

import (
	"bufio"
	"fmt"
	"go_grpc_study/example_5/grpc_proto/stream_proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

// 新版本 gRPC 要求必须嵌入 UnimplementedGreeterServer 结构体
type ClientStream struct {
	stream_proto.UnimplementedClientStreamServer
}

func (ClientStream) UploadFile(stream stream_proto.ClientStream_UploadFileServer) error {
	//os.O_CREATE ： 创建并打开一个新文件
	//os.O_TRUNC ：打开一个文件并截断它的长度为零（必须有写权限）
	//os.O_WRONLY ：以只写的方式打开
	file, err := os.OpenFile("../static/grpc_x.png", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
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
		writer.Write(response.Content)
		fmt.Printf("第 %d 次写入数据\n", index)
	}
	writer.Flush()
	stream.SendAndClose(&stream_proto.Response{Message: "服务端接收到上传的文件了"})
	return nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	// 将server结构体注册为gRPC服务。
	stream_proto.RegisterClientStreamServer(s, &ClientStream{})
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
