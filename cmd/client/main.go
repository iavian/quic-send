package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/iavian/quic-send/common"
	"github.com/lucas-clemente/quic-go"
)

const addr = "localhost:4242"
const message = "foobar"

func main() {

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-file"},
	}

	quicConfig := &quic.Config{GetLogWriter: func(connectionID []byte) io.WriteCloser {
		filename := fmt.Sprintf("client_%x.qlog", connectionID)
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Creating qlog file %s.\n", filename)
		return f
	}}
	session, err := quic.DialAddr(common.ClientServerAddr, tlsConf, quicConfig)
	if err != nil {
		panic(err)
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Client: Sending '%s'\n", message)
	_, err = stream.Write([]byte(message))
	if err != nil {
		panic(err)
	}

}
