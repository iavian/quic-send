package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/iavian/quic-send/client"
	"github.com/iavian/quic-send/common"
	"github.com/lucas-clemente/quic-go/http3"
)

func mainy() {
	pool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: true,
		},
	}
	defer roundTripper.Close()
	hclient := &http.Client{
		Transport: roundTripper,
	}
	file, err := os.Open("tfile")
	if err != nil {
		panic(err)
	}
	res, err := hclient.Post("https://quic.iavian.net:8080/upload", "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("All done")
}

func main() {
	start := time.Now()
	c := client.NewFileClient(common.ClientServerAddr)
	file := "afile"
	err := c.Upload(file)
	elapsed := time.Since(start)
	if err != nil {
		log.Printf("upload/download file error: %v\n", err)
	} else {
		log.Printf("upload/download file success: %s %s\n", file, elapsed)
	}
	c.Close()
}
