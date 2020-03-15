package main

import (
	"context"
	"crypto/tls"
	"fmt"

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
	session, err := quic.DialAddr(common.ClientServerAddr, tlsConf, nil)
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
