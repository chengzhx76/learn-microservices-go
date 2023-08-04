package main

import (
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"
	"io"
	"log"
	"os"
)

var (
	caCert = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\ca-gm-cert.crt"

	clientAuthFileKey  = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\client-gm-auth-key.pem"
	clientAuthFileCert = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\client-gm-auth-cert.crt"
)

func main() {
	//gmSocketClientSingleAuth()
	gmHttpClientDoubleAuth()
}

func gmSocketClientSingleAuth() {
	// 信任的根证书
	certPool := x509.NewCertPool()
	cacert, err := os.ReadFile(caCert)
	if err != nil {
		log.Fatal(err)
	}
	ok := certPool.AppendCertsFromPEM(cacert)
	if !ok {
		return
	}

	config := &gmtls.Config{
		GMSupport:          &gmtls.GMSupport{},
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	}

	conn, err := gmtls.Dial("tcp", "umf.com:50052", config)
	if err != nil {
		fmt.Printf("%s\r\n", err)
		return
	}
	defer conn.Close()

	req := []byte("GET / HTTP/1.1\r\n" +
		"Host: 127.0.0.1\r\n" +
		"Connection: close\r\n\r\n")
	_, _ = conn.Write(req)
	buff := make([]byte, 1024)
	for {
		n, _ := conn.Read(buff)
		if n <= 0 {
			break
		} else {
			fmt.Printf("收到回应[%s]\r\n", buff[0:n])
		}
	}
	fmt.Println("国密单向校验通过,纯Socket模式")
}

func gmHttpClientSingleAuth() {
	// 1. 提供根证书链
	certPool := x509.NewCertPool()
	cacert, err := os.ReadFile(caCert)
	if err != nil {
		panic(err)
	}
	certPool.AppendCertsFromPEM(cacert)
	// 3. 构造HTTP客户端。
	httpClient := gmtls.NewHTTPSClient(certPool)
	// 4. 调用API访问HTTPS。
	response, err := httpClient.Get("https://umf.com:50052")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	// 使用 response 做你需要的事情...
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println("国密单向校验通过,httpClient模式")
}

func gmSocketClientDoubleAuth() {
	// 信任的根证书
	certPool := x509.NewCertPool()
	cacert, err := os.ReadFile(caCert)
	if err != nil {
		log.Fatal(err)
	}
	ok := certPool.AppendCertsFromPEM(cacert)
	if !ok {
		return
	}
	cert, err := gmtls.LoadX509KeyPair(clientAuthFileCert, clientAuthFileKey)
	if err != nil {
		log.Fatal(err)
	}
	config := &gmtls.Config{
		GMSupport:          &gmtls.GMSupport{},
		RootCAs:            certPool,
		Certificates:       []gmtls.Certificate{cert},
		InsecureSkipVerify: false,
	}

	conn, err := gmtls.Dial("tcp", "umf.com:50052", config)
	if err != nil {
		fmt.Printf("%s\r\n", err)
		return
	}
	defer conn.Close()

	req := []byte("GET / HTTP/1.1\r\n" +
		"Host: 127.0.0.1\r\n" +
		"Connection: close\r\n\r\n")
	_, _ = conn.Write(req)
	buff := make([]byte, 1024)
	for {
		n, _ := conn.Read(buff)
		if n <= 0 {
			break
		} else {
			fmt.Printf("收到回应[%s]\r\n", buff[0:n])
		}
	}
	fmt.Println("国密双向校验通过,纯Socket模式")
}

func gmHttpClientDoubleAuth() {
	// 1. 提供根证书链
	certPool := x509.NewCertPool()
	cacert, err := os.ReadFile(caCert)
	if err != nil {
		panic(err)
	}
	certPool.AppendCertsFromPEM(cacert)
	// 2. 提供客户端认证证书、密钥对。
	clientAuthCert, err := gmtls.LoadX509KeyPair(clientAuthFileCert, clientAuthFileKey)
	// 3. 构造HTTP客户端。
	httpClient := gmtls.NewAuthHTTPSClient(certPool, &clientAuthCert)
	// 4. 调用API访问HTTPS。
	response, err := httpClient.Get("https://umf.com:50052")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	// 使用 response 做你需要的事情...
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println("国密双向校验通过,httpClient模式")
}
