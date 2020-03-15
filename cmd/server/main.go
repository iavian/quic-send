package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/iavian/quic-send/common"
	"github.com/lucas-clemente/quic-go"
)

func main() {

	quicConfig := &quic.Config{}
	quicConfig.GetLogWriter = func(connectionID []byte) io.WriteCloser {
		filename := fmt.Sprintf("server_%x.qlog", connectionID)
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Creating qlog file %s.\n", filename)
		return f
	}
	listener, err := quic.ListenAddr(common.ServerAddr, generateTLSConfig(), quicConfig)
	if err != nil {
		panic(err)
	}
	sess, err := listener.Accept(context.Background())
	if err != nil {
		panic(err)
	}
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}
	log.Printf("Stream Accepted")

	tmpAbsPath, err := ioutil.TempFile(".", "*")
	defer tmpAbsPath.Close()

	writen, err := io.Copy(tmpAbsPath, stream)

	if err != nil {
		panic(err)
	}

	log.Printf("accept session error: %v\n", writen)
}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-file"},
	}
}
