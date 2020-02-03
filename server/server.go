package server

import (
	"bufio"
	"context"
	"crypto/tls"
	"log"
	"os"

	"github.com/lucas-clemente/quic-go"
)

type FileServer struct {
	Address    string
	TLSConfig  *tls.Config
	QuicConfig *quic.Config
	Sessions   map[int64]*quic.Session
	Listener   quic.Listener
	Ctx        context.Context
}

func NewFileServer(address string, tlsConfig *tls.Config, quicConfig *quic.Config) *FileServer {
	return &FileServer{
		Address:    address,
		TLSConfig:  tlsConfig,
		QuicConfig: quicConfig,
		Sessions:   make(map[int64]*quic.Session, 0),
		Ctx:        context.Background(),
	}
}

func (s *FileServer) Run() error {
	var err error
	s.Listener, err = quic.ListenAddr(s.Address, s.TLSConfig, s.QuicConfig)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	log.Println("listening on port ", s.Address)
	for {
		sess, err := s.Listener.Accept(s.Ctx)
		if err != nil {
			log.Printf("accept session error: %v\n", err)
			continue
		}

		sessionHandler := NewSessionHandler(&sess)
		go sessionHandler.Run()
	}
}

func writeFile(file string) (*bufio.Writer, error) {
	fp, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return bufio.NewWriterSize(fp, 1000), nil
}
