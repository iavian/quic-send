package main

import (
	"log"
	"sync"

	"github.com/iavian/quic-send/client"
	"github.com/iavian/quic-send/common"
)

func main() {
	c := client.NewFileClient(common.ClientServerAddr)
	var wg sync.WaitGroup
	files := [1]string{"tfile"}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			err := c.Upload(file)
			if err != nil {
				log.Printf("upload/download file error: %v\n", err)
			} else {
				log.Printf("upload/download file success: %s\n", file)
			}
			wg.Done()
		}(file)
	}
	wg.Wait()
	c.Close()
}
