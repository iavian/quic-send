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
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

func main() {
	if len(os.Args) > 1 {
		start := time.Now()
		c := client.NewFileClient(common.ClientServerAddr)
		defer c.Close()
		file := "tfile"
		err := c.Upload(file)
		elapsed := time.Since(start)
		if err != nil {
			log.Printf("upload/download file error: %v\n", err)
		} else {
			log.Printf("upload/download file success: %s %s\n", file, elapsed)
		}
	} else {
		pool, err := x509.SystemCertPool()
		if err != nil {
			log.Fatal(err)
		}
		roundTripper := &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				RootCAs:            pool,
				InsecureSkipVerify: true,
			}, DisableCompression: false, QuicConfig: &quic.Config{KeepAlive: true},
		}
		/*	roundTripper.QuicConfig.GetLogWriter = func(connectionID []byte) io.WriteCloser {
			filename := fmt.Sprintf("logs/client_%x.qlog", connectionID)
			f, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Creating qlog file %s.\n", filename)
			return struct {
				io.Writer
				io.Closer
			}{bufio.NewWriter(f), f}
		}*/
		defer roundTripper.Close()
		hclient := &http.Client{
			Transport: roundTripper,
		}
		start := time.Now()
		file, err := os.Open("tfile")
		if err != nil {
			panic(err)
		}
		res, err := hclient.Post("https://vpn.iavian.net:8080/upload", "binary/octet-stream", file)
		elapsed := time.Since(start)
		if err != nil {
			panic(err)
		} else {
			log.Printf("upload/download file success: %s %s\n", file.Name(), elapsed)
		}
		defer res.Body.Close()
		fmt.Println("All done")
	}
}
