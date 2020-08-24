package main

import (
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

// https://github.com/DavidCai1993/my-blog/issues/35
// https://zhuanlan.zhihu.com/p/103534192

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func main() {
	flag.Parse()

	h := requestHandler

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/plain; charset=utf8")
	ctx.WriteString("ok")
	ctx.SetStatusCode(200)
}
