package main

import (
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/gmtls/gmcredentials"
	"github.com/tjfoc/gmsm/x509"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"learn-microservices-go/gmgrpc/echo"
	"log"
	"os"
)

const address = "umf.com:50051"

var (
	caFile       = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\ca-gm-cert.crt"
	userCertFile = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\client-gm-auth-cert.crt"
	userKeyFile  = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\client-gm-auth-key.pem"
)

func main() {
	clientRun()
}

func clientRun() {
	cert, err := gmtls.LoadX509KeyPair(userCertFile, userKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	caCert, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(caCert)
	creds := gmcredentials.NewTLS(&gmtls.Config{
		GMSupport: &gmtls.GMSupport{},
		//ServerName:   "test.example.com",
		Certificates: []gmtls.Certificate{cert},
		RootCAs:      certPool,
		ClientAuth:   gmtls.RequireAndVerifyClientCert,
	})
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("cannot to connect: %v", err)
	}
	defer conn.Close()
	c := echo.NewEchoClient(conn)
	echoTest(c)
}

func echoTest(c echo.EchoClient) {
	r, err := c.Echo(context.Background(), &echo.EchoRequest{Req: "hello"})
	if err != nil {
		log.Fatalf("failed to echo: %v", err)
	}
	fmt.Printf("%s\n", r.Result)
}
