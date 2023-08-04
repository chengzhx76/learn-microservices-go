package main

import (
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"
	"net/http"
	"os"
	"time"
)

var (
	caCert = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\ca-gm-cert.crt"

	signFileKey  = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-sign-key.pem"
	signFileCert = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-sign-cert.crt"

	encFileKey  = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-enc-key.pem"
	encFileCert = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-enc-cert.crt"
)

/*var (
	caCert = "ca-gm-cert.crt"

	signFileKey  = "server-gm-sign-key.pem"
	signFileCert = "server-gm-sign-cert.crt"

	encFileKey  = "server-gm-enc-key.pem"
	encFileCert = "server-gm-enc-cert.crt"
)*/

/**
GO (国密,标准Https) 单向,双向认证Demo
https://blog.csdn.net/HeroRazor/article/details/121633211
*/

func main() {
	//gmServerSingleAuth()
	gmServerDoubleAuth()
}

func gmServerSingleAuth() {
	sigCert, err := gmtls.LoadX509KeyPair(signFileCert, signFileKey)
	if err != nil {
		fmt.Println(err)
	}
	encCert, err := gmtls.LoadX509KeyPair(encFileCert, encFileKey)
	if err != nil {
		fmt.Println(err)
	}
	config := &gmtls.Config{
		GMSupport:    gmtls.NewGMSupport(),
		Certificates: []gmtls.Certificate{sigCert, encCert},
	}
	if err != nil {
		fmt.Println(err)
	}
	ln, err := gmtls.Listen("tcp", ":50052", config)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		p := fmt.Sprintf(`
Method:%v|Proto:%v|Host:%v
RequestURI:%v|UserAgent:%v|Referer:%v
Scheme:%v|URL.Host:%v|URL.Port:%v
URL.Path:%v|URL.RawPath:%v|URL.RawQuery:%v
URL.RawFragment:%v|URL.EscapedFragment:%v|URL.EscapedPath:%v
URL.RequestURI:%v|URL.String:%v
`,
			request.Method, request.Proto, request.Host,
			request.RequestURI, request.UserAgent(), request.Referer(),
			request.URL.Scheme, request.URL.Host, request.URL.Port(),
			request.URL.Path, request.URL.RawPath, request.URL.RawQuery,
			request.URL.RawFragment, request.URL.EscapedFragment(), request.URL.EscapedPath(),
			request.URL.RequestURI(), request.URL.String())

		fmt.Println(p)

		//fmt.Println("" + request.TLS)

		_, _ = writer.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))

	})
	fmt.Println(">> HTTP :50052 [GMSSL] Client Auth running...")
	err = http.Serve(ln, serveMux)
	if err != nil {
		fmt.Println(err)
	}

}

func gmServerDoubleAuth() {
	sigCert, err := gmtls.LoadX509KeyPair(signFileCert, signFileKey)
	if err != nil {
		fmt.Println(err)
	}
	encCert, err := gmtls.LoadX509KeyPair(encFileCert, encFileKey)
	if err != nil {
		fmt.Println(err)
	}
	certPool := x509.NewCertPool()
	cacert, err := os.ReadFile(caCert)
	if err != nil {
		fmt.Println(err)
	}
	certPool.AppendCertsFromPEM(cacert)
	config := &gmtls.Config{
		GMSupport:    &gmtls.GMSupport{},
		Certificates: []gmtls.Certificate{sigCert, encCert},
		ClientAuth:   gmtls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	if err != nil {
		fmt.Println(err)
	}
	ln, err := gmtls.Listen("tcp", ":50052", config)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		p := fmt.Sprintf(`
Method:%v|Proto:%v|Host:%v
RequestURI:%v|UserAgent:%v|Referer:%v
Scheme:%v|URL.Host:%v|URL.Port:%v
URL.Path:%v|URL.RawPath:%v|URL.RawQuery:%v
URL.RawFragment:%v|URL.EscapedFragment:%v|URL.EscapedPath:%v
URL.RequestURI:%v|URL.String:%v
`,
			request.Method, request.Proto, request.Host,
			request.RequestURI, request.UserAgent(), request.Referer(),
			request.URL.Scheme, request.URL.Host, request.URL.Port(),
			request.URL.Path, request.URL.RawPath, request.URL.RawQuery,
			request.URL.RawFragment, request.URL.EscapedFragment(), request.URL.EscapedPath(),
			request.URL.RequestURI(), request.URL.String())

		fmt.Println(p)

		_, _ = writer.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))

	})
	fmt.Println(">> HTTP :50052 [GMSSL] Client Auth running...")
	err = http.Serve(ln, serveMux)
	if err != nil {
		fmt.Println(err)
	}
}
