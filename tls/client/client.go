package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	caCrt      = "D:\\golang\\src\\learn-microservices-go\\tls\\certs\\ca.crt"
	clientCert = "D:\\golang\\src\\learn-microservices-go\\tls\\certs\\client.crt"
	clientKey  = "D:\\golang\\src\\learn-microservices-go\\tls\\certs\\client.key"
)

func main() {
	pool := x509.NewCertPool()

	caCrt, err := os.ReadFile(caCrt)
	if err != nil {
		panic(err)
	}

	pool.AppendCertsFromPEM(caCrt)

	clientCrt, err := tls.LoadX509KeyPair(clientCert, clientKey)

	if err != nil {
		panic(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{clientCrt},
		},
	}

	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("https://localhost:8443")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
}
