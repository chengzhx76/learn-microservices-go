package test

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"testing"
)

var (
	encFileKey  = "D:/golang/src/learn-microservices-go/gmtls/certs2/server-gm-enc-key.pem"
	encFileCert = "D:/golang/src/learn-microservices-go/gmtls/certs2/server-gm-enc-cert.crt"
)

// 读取公私钥
func TestSm2_read_pub_pri(t *testing.T) {
	cert, err := gmtls.LoadX509KeyPair(encFileCert, encFileKey)
	priKey, ok := cert.PrivateKey.(*sm2.PrivateKey)
	if !ok {
		t.Fatal("tls: certificate private key does not implement crypto.Decrypter")
		return
	}
	pubKey := &priKey.PublicKey

	privPem, err := x509.WritePrivateKeyToPem(priKey, nil) // 生成密钥文件
	if err != nil {
		fmt.Printf("Error: write private key: %v\n", err)
		return
	}
	fmt.Printf("read priKey: \n%s\n", string(privPem))

	pubkeyPem, err := x509.WritePublicKeyToPem(pubKey) // 生成公钥文件
	if err != nil {
		fmt.Printf("Error: write public key: %v\n", err)
		return
	}
	fmt.Printf("read pubKey: \n%s\n", string(pubkeyPem))

	/*err = os.WriteFile("priv.pem", privPem, os.FileMode(0644))
	if err != nil {
		fmt.Printf("Error: write priv file: %v\n", err)
		return
	}*/

}

// 私钥解密
func TestSm2_priKeyToHexStrDecryptAsn1(t *testing.T) {
	plaintextHex := "01014d1ed402d6d88e128744bbf5480254971fb4237264dee532fc10e47f9a8fc2641556084f68f9d92204d7bce62ff6"
	encryptHex := "30819a0221008eb41dac5abdafa3d3a1eaed3753f58e265f27f20c4948beb03ff7eafc2d28e7022100a1834d4ec03de7230b5e71d7a567de602ccd3fef90b4a5f69cc1c33aa0b26e35042092e151166b1c2586ffb2d54b35b80468e7ddd0611004f107aacca38d7d78fb040430ae0179e051cad0a18a10797ae0e14238a591833916f246c60915dba8253add55f09f00e9de0f772b5f41e9e6108ded7f"

	cert, err := gmtls.LoadX509KeyPair(encFileCert, encFileKey)
	priv, ok := cert.PrivateKey.(*sm2.PrivateKey)
	if !ok {
		t.Fatal("tls: certificate private key does not implement crypto.Decrypter")
		return
	}
	//pub := &priv.PublicKey

	encryptBytes, err := hex.DecodeString(encryptHex)
	if err != nil {
		fmt.Printf("Error: hex decode")
		return
	}

	decryptBytes, err := sm2.DecryptAsn1(priv, encryptBytes)
	if err != nil {
		fmt.Printf("Error: failed to decrypt: %v\n", err)
	}
	//fmt.Printf("clear text = %s\n", decryptBytes)
	decryptStr := hex.EncodeToString(decryptBytes)
	fmt.Printf("decryptStr   text = %v\n", decryptStr)
	fmt.Printf("plaintextHex text = %v\n", plaintextHex)

}

// 公钥加密
func TestSm2_pubKeyToEncryptAsn1(t *testing.T) {
	plaintextHex := "010177f9368e67b0d0ed8e72642f00d0af448184f02515ee2afd6e54d73dbfae76bc291dbfa7f881fb0d3594c2ceabcb"
	encryptHex := "308199022100e24139de2b4733ee170536a77480093d205659fa7f302ec010a3ba997d1d89e902204ffb41f1cb64d779d6aedff31c5a6deda5fb66438fb275ee472c6a4e51a8b7220420790898b70d435c1aad325966c87f0be14d72b9918ba8896c5962bc8d2a5d87150430dd0000feff5e4c40461322cdad2335110b34595b1aef0cd4a4bc18b917e2c0a146a4fd695b65426ce43ff17e7a3d7359"

	cert, err := gmtls.LoadX509KeyPair(encFileCert, encFileKey)
	priv, ok := cert.PrivateKey.(*sm2.PrivateKey)
	if !ok {
		t.Fatal("tls: certificate private key does not implement crypto.Decrypter")
		return
	}
	pub := &priv.PublicKey

	plaintextBytes, err := hex.DecodeString(plaintextHex)
	if err != nil {
		fmt.Printf("Error: hex decode")
		return
	}

	encryptBytes, err := sm2.EncryptAsn1(pub, plaintextBytes, rand.Reader)
	if err != nil {
		fmt.Printf("Error: failed to encrypt %v\n", err)
		return
	}
	//fmt.Printf("Cipher text = %v\n", encryptBytes)
	encryptStr := hex.EncodeToString(encryptBytes)
	fmt.Printf("encryptStr text = %v\n", encryptStr)
	fmt.Printf("encryptHex text = %v\n", encryptHex)

}

func TestSm2_c(t *testing.T) {
	//priv, err := GenerateKey(rand.Reader) // 生成密钥对

	cert, err := gmtls.LoadX509KeyPair(encFileCert, encFileKey)
	priv, ok := cert.PrivateKey.(*sm2.PrivateKey)
	if !ok {
		t.Fatal("tls: certificate private key does not implement crypto.Decrypter")
		return
	}
	pub := &priv.PublicKey

	//msg := []byte("123456")
	msg, err := hex.DecodeString("010119df0a6bfcb58391d845da04388b68b923e8dd11705393688347a00e54ef132e9fc95ae57b564f9a359f7150d6e4")
	if err != nil {
		fmt.Printf("Error: hex decode")
		return
	}

	encryptBytes, err := sm2.EncryptAsn1(pub, msg, rand.Reader)
	if err != nil {
		fmt.Printf("Error: failed to encrypt %s: %v\n", msg, err)
		return
	}
	//fmt.Printf("Cipher text = %v\n", encryptBytes)
	encryptStr := hex.EncodeToString(encryptBytes)
	fmt.Printf("encryptStr text = %v\n", encryptStr)

	decryptBytes, err := sm2.DecryptAsn1(priv, encryptBytes)
	if err != nil {
		fmt.Printf("Error: failed to decrypt: %v\n", err)
	}
	//fmt.Printf("clear text = %s\n", decryptBytes)
	decryptStr := hex.EncodeToString(decryptBytes)
	fmt.Printf("decryptStr text = %v\n", decryptStr)
}

func Test_req(t *testing.T) {
	h := md5.New()
	str := hex.EncodeToString(h.Sum(nil))

	//b, err := hex.DecodeString(str)

	t.Log(str)
}
