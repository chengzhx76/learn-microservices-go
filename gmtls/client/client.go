package main

import (
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"
	"io"
	"os"
)

var (
	authPriKeyPath = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\client-gm-auth-key.pem"
	authCertPath   = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\client-gm-auth-cert.crt"
	rootCertPath   = "D:\\golang\\src\\learn-microservices-go\\gmtls\\certs2\\ca-gm-cert.crt"
)

func main() {
	/*config, err := createClientGMTLSConfig("D:\\golang\\src\\learn-microservices-go\\gmtls\\certs\\sm2_auth_key.pem",
	"D:\\golang\\src\\learn-microservices-go\\gmtls\\certs\\sm2_auth_cert.cer",
	[]string{"D:\\golang\\src\\learn-microservices-go\\gmtls\\certs\\SM2_CA.cer"})*/

	config, err := createClientGMTLSConfig(authPriKeyPath, authCertPath, []string{rootCertPath})

	if err != nil {
		panic(err)
	}
	httpClient := gmtls.NewCustomHTTPSClient(config)

	/*reqTicker := time.NewTicker(3 * time.Second)
	defer reqTicker.Stop()
	for {
		<-reqTicker.C
		go func() {*/
	//response, err := httpClient.Get("https://umf.com:50055/UChains/poeBatch/batch/umfnet_poeBatch/test")
	response, err := httpClient.Get("https://umf.com:18431/UChains/baseinfo")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	raw, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("====> " + string(raw))
	/*}()
	}*/

}

func createClientGMTLSConfig(keyPath string, certPath string, caPaths []string) (*gmtls.Config, error) {

	cfg := &gmtls.Config{
		GMSupport: &gmtls.GMSupport{},
	}
	cfg.Certificates = []gmtls.Certificate{}
	if keyPath != "" && certPath != "" {
		cert, err := gmtls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, fmt.Errorf("load gm X509 keyPair error: %v", err)
		}
		cfg.Certificates = append(cfg.Certificates, cert)
	}

	var pool *x509.CertPool = nil
	if len(caPaths) > 0 {
		pool = x509.NewCertPool()
		for _, certPath := range caPaths {
			caCrt, err := os.ReadFile(certPath)
			if err != nil {
				return nil, err
			}
			ok := pool.AppendCertsFromPEM(caCrt)
			if !ok {
				return nil, fmt.Errorf("append cert to pool fail at %s", certPath)
			}
		}
	}

	cfg.MinVersion = gmtls.VersionGMSSL
	cfg.MaxVersion = gmtls.VersionTLS12

	cfg.PreferServerCipherSuites = true
	// cfg.CipherSuites use default value []uint16{GMTLS_SM2_WITH_SM4_SM3, GMTLS_ECDHE_SM2_WITH_SM4_SM3}

	cfg.RootCAs = pool
	// cfg.ServerName = "localhost"
	cfg.InsecureSkipVerify = false

	return cfg, nil

}
