package main

import (
	"fmt"
	"go_grpc_study/example_4/grpc_proto/stream_proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

//服务端流式

// 新版本 gRPC 要求必须嵌入 UnimplementedGreeterServer 结构体
type ServiceStream struct {
	stream_proto.UnimplementedServiceStreamServer
}

func (ServiceStream) DownloadFile(request *stream_proto.Request, stream stream_proto.ServiceStream_DownloadFileServer) error {
	fmt.Println("DownloadFile", request)
	//分片读的方式读取图片
	file, err := os.Open("../static/grpc.png")
	if err != nil {
		return err
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
		stream.Send(&stream_proto.FileResponse{
			Content: buf,
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
	s := grpc.NewServer()
	// 将server结构体注册为gRPC服务。
	stream_proto.RegisterServiceStreamServer(s, &ServiceStream{})
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
