package main

import (
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/gmtls/gmcredentials"
	"github.com/tjfoc/gmsm/x509"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"learn-microservices-go/gmgrpc/echo"
	"log"
	"net"
	"os"
)

const (
	port = ":50051"
)

var (
	caFile       = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\ca-gm-cert.crt"
	signCertFile = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\server-gm-sign-cert.crt"
	signKeyFile  = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\server-gm-sign-key.pem"
	encCertFile  = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\server-gm-enc-cert.crt"
	encKeyFile   = "D:\\golang\\src\\learn-microservices-go\\gmgrpc\\certs\\server-gm-enc-key.pem"
)

func main() {
	serverRun()
}

type server struct{}

func (s *server) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Result: req.Req}, nil
}

func serverRun() {
	signCert, err := gmtls.LoadX509KeyPair(signCertFile, signKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	encCert, err := gmtls.LoadX509KeyPair(encCertFile, encKeyFile)
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
		GMSupport:    &gmtls.GMSupport{},
		ClientAuth:   gmtls.RequireAndVerifyClientCert,
		Certificates: []gmtls.Certificate{signCert, encCert},
		ClientCAs:    certPool,
	})

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	grpcSev := grpc.NewServer(grpc.Creds(creds))

	echo.RegisterEchoServer(grpcSev, &server{})
	err = grpcSev.Serve(listen)
	if err != nil {
		log.Fatalf("Serve: %v", err)
	}
}
