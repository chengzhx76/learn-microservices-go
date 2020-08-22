package main

import (
	"flag"
	"fmt"
)

type serverFlag []string

func (s *serverFlag) String() string {
	return "server flag"
}

func (s *serverFlag) Set(value string) error {
	*s = append(*s , value)
	return nil
}

var (
	ip = flag.String("addr", "127.0.0.1", "服务器地址")
	port = flag.Int("port", 9090, "服务器端口")
	enable = flag.Bool("enble", false, "是否激活服务")

	servers = &serverFlag{}

)

func main() {
	fmt.Println("hello world!")

	flag.Var(servers, "server", "服务器资源")

	flag.Parse()
	fmt.Printf("%v:%d->%v", *ip, *port, *enable)

	fmt.Println()

	fmt.Printf("server->%v", *servers)




}


