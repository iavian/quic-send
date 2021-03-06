package server

import (
	"context"
	"log"

	"github.com/iavian/quic-send/common"
	"github.com/lucas-clemente/quic-go"
)

// SessionHandler ..
type SessionHandler struct {
	Ctx     context.Context
	Session quic.Session
	Streams map[int64]*StreamHandler
}

// NewSessionHandler ..
func NewSessionHandler(session *quic.Session) *SessionHandler {
	return &SessionHandler{
		Ctx:     context.Background(),
		Session: *session,
		Streams: make(map[int64]*StreamHandler, 0),
	}
}

//Run ..
func (h *SessionHandler) Run() {
	for {
		stream, err := h.Session.AcceptStream(h.Ctx)
		if err != nil {
			if err.Error() != common.NoError {
				log.Printf("accept stream error: %v", err)
			}
			break
		}
		streamHandler := NewStreamHandler(&stream)
		go streamHandler.Run()
	}
}
