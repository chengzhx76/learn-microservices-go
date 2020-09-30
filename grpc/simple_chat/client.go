package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"learn-microservices-go/grpc/simple_chat/chat"
)

// golang grpc 负载均衡 https://blog.csdn.net/weixin_33750452/article/details/92740999
// gRPC服务发现&负载均衡 https://blog.csdn.net/qq_21816375/article/details/78159297
// gRPC服务发现&负载均衡 https://www.cnblogs.com/FireworksEasyCool/p/12912839.html

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := chat.NewChatServiceClient(conn)

	response, err := client.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	response, err = client.BroadcastMessage(context.Background(), &chat.Message{Body: "Message to Broadcast!"})
	if err != nil {
		log.Fatalf("Error when calling Broadcast Message: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
