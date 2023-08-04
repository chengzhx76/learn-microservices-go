package main

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"
	"io"
	"net/http"
	"os"
	"time"
)

func main2() {

	caCert := "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\ca-gm-cert.crt"

	signFileKey := "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-sign-key.pem"
	signFileCert := "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-sign-cert.crt"

	encFileKey := "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-enc-key.pem"
	encFileCert := "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\server-gm-enc-cert.crt"

	/*caCert := "D:\\golang\\src\\learn-microservices-go\\gmtls\\java\\chain-ca.crt"

	signFileKey := "D:\\golang\\src\\learn-microservices-go\\gmtls\\java\\server_sign.key"
	signFileCert := "D:\\golang\\src\\learn-microservices-go\\gmtls\\java\\server_sign.crt"

	encFileKey := "D:\\golang\\src\\learn-microservices-go\\gmtls\\java\\server_enc.key"
	encFileCert := "D:\\golang\\src\\learn-microservices-go\\gmtls\\java\\server_enc.crt"*/

	sigCert, err := gmtls.LoadX509KeyPair(signFileCert, signFileKey)
	if err != nil {
		panic(err)
	}
	encCert, err := gmtls.LoadX509KeyPair(encFileCert, encFileKey)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	cacert, err := os.ReadFile(caCert)
	if err != nil {
		panic(err)
	}

	certPool.AppendCertsFromPEM(cacert)

	config := &gmtls.Config{
		GMSupport:    &gmtls.GMSupport{},
		Certificates: []gmtls.Certificate{sigCert, encCert},
		ClientAuth:   gmtls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	if err != nil {
		panic(err)
	}

	conn, err := gmtls.Listen("tcp", ":50055", config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))

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

		/*req := request.Clone(context.Background())
		reqBody, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		url := fmt.Sprintf("http://%s%s", "10.10.77.118:7051", req.RequestURI)
		proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(reqBody))

		proxyReq.Header = make(http.Header)
		for h, val := range req.Header {
			proxyReq.Header[h] = val
		}

		client := &http.Client{}
		resp, err := client.Do(proxyReq)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		respBody, err := switchContentEncoding(resp)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadGateway)
			return
		}
		result, err := io.ReadAll(respBody)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadGateway)
			return
		}
		writer.Write(result)*/

	})

	fmt.Println(">> HTTP :50055 [GMSSL] Client Auth running...")
	err = http.Serve(conn, serveMux)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Minute)
}

func switchContentEncoding(resp *http.Response) (bodyReader io.Reader, err error) {
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		bodyReader, err = gzip.NewReader(resp.Body)
	case "deflate":
		bodyReader = flate.NewReader(resp.Body)
	default:
		bodyReader = resp.Body
	}
	return
}
