package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type FooHandler struct {
	Language string
	Time     time.Time
}

func (f FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s-%s", f.Language, f.Time.Format("2006-01-02 15:04:05"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/test-gm/time", FooHandler{Language: "国密测试，服务器时间", Time: time.Now()})
	mux.Handle("/test-gm/chinese", FooHandler{Language: "你好", Time: time.Now()})
	mux.Handle("/test-gm/english", FooHandler{Language: "hello", Time: time.Now()})

	server := &http.Server{
		Addr:    ":8881",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
