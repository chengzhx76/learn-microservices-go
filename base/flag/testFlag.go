package main

import (
	"flag"
	"fmt"
)

type A int32   //类型定义，生成新的
type B = int32 //别名定义，仅仅是 alias

// 类型等价定义，相当于类型重命名 此处相当于 serverFlag 等于 []string
// 相当于增加了方法 `String()` `Set(value string)`
type serverFlag []string

func (s *serverFlag) String() string {
	return "server flag"
}

func (s *serverFlag) Set(value string) error {
	fmt.Println("call->Set(value string)")
	*s = append(*s, value)
	return nil
}

var (
	ip     = flag.String("addr", "127.0.0.1", "服务器地址")
	port   = flag.Int("port", 9090, "服务器端口")
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
