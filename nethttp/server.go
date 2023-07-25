package main

import (
	"fmt"
	"log"
	"net/http"
)

type FooHandler struct {
	Language string
}

func (f FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", f.Language)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/chinese", FooHandler{Language:"你好"})
	mux.Handle("/english", FooHandler{Language:"hello"})

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}