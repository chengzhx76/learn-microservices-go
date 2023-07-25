package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

var (
	caCerts    = "D:\\golang\\src\\learn-microservices-go\\tls\\certs\\ca.crt"
	serverCert = "D:\\golang\\src\\learn-microservices-go\\tls\\certs\\server.crt"
	serverKey  = "D:\\golang\\src\\learn-microservices-go\\tls\\certs\\server.key"
)

type myServer struct {
}

func (server *myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func main() {
	pool := x509.NewCertPool()

	caCrt, err := os.ReadFile(caCerts)
	if err != nil {
		panic(err)
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8443",
		Handler: &myServer{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS(serverCert, serverKey)

	if err != nil {
		panic(err)
	}
}
