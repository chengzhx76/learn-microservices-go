package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"github.com/soheilhy/cmux"

)

func main() {
	// 创建监听端口的 Listener
	l, err := net.Listen("tcp", ":23456")
	if err != nil {
		log.Fatal(err)
	}

	// 根据创建成功的 Listener 初始化 cmux 实例
	m := cmux.New(l)

	// 注册不同协议的 *协议匹配器*
	grpcL := m.Match(cmux.HTTP2HeaderField("context-type", "application/grpc")) // 使用 context-type 标识为 grpc 协议
	httpL := m.Match(cmux.HTTP1Fast()) // 使用 cmux 内置函数标识 http1 协议
	trpcL := m.Match(cmux.Any()) // 使用 Any 函数匹配未匹配的任意协议

	// 初始化不同协议的服务实例
	grpcS := grpc.NewServer()
	grpchello.RegisterGreeterServer(grpcs, &server{})

	httpS := http.Server{
		Handler: &helloHTTP1Handler{}
	}


	// 启动所有注册协议的服务
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	m.Serve() // 启动 cmux 服务实例

}