package main

import (
	"log"
	"time"

	"github.com/iavian/quic-send/client"
	"github.com/iavian/quic-send/common"
)

func main() {
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

}
