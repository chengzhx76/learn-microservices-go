package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	"learn-microservices-go/grpc-load-balancing/etcdv3"
	pb "learn-microservices-go/grpc-load-balancing/proto"
)

var (
	grpcClient pb.SimpleClient
)

var etcdEndpoints = []string{"180.76.183.68:2379"}
var serName = "simple_grpc"

func main() {
	r := etcdv3.NewServiceDiscovery(etcdEndpoints)
	resolver.Register(r)
	// 连接服务器
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", r.Scheme(), serName), grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
	for i := 0; i < 100; i++ {
		route(i)
		time.Sleep(1 * time.Second)
	}

}

// route 调用服务端Route方法
func route(i int) {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc " + strconv.Itoa(i),
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
