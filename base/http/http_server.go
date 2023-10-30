package main

import (
	"io"
	"log"
	"net/http"
)

// Golang构建HTTP服务（一）--- net/http库源码笔记 https://www.jianshu.com/p/be3d9cdc680b
// Golang构建HTTP服务（二）--- Handler，ServeMux与中间件 https://www.jianshu.com/p/16210100d43d

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world!")
}

func main() {
	http.HandleFunc("/hello", helloServer)
	err := http.ListenAndServe(":8092", nil)
	if err != nil {
		log.Fatal("listenAndserve: ", err)
	}
}
