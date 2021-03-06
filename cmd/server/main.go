package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/iavian/quic-send/common"
	"github.com/iavian/quic-send/server"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	if len(os.Args) > 1 {
		log.Printf("Listening on port %s for QUIC\n", common.ServerAddr)
		quicConfig := &quic.Config{}
		/*quicConfig.GetLogWriter = func(connectionID []byte) io.WriteCloser {
			filename := fmt.Sprintf("client_%x.qlog", connectionID)
			f, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Creating qlog file %s.\n", filename)
			return f
		}*/
		s := server.NewFileServer(common.ServerAddr, generateTLSConfig(), quicConfig)
		s.Run()
	} else {
		log.Printf("Listening on port %s for http3\n", common.ServerAddr)
		http.HandleFunc("/upload", uploadFile)
		http3.ListenAndServeQUIC(":8080", "./certs/quic.cert", "./certs/quic.key", nil)
	}
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("./result")
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
	fmt.Printf("Nice %d\n", n)
}
