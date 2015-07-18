package main

import (
	"errors"

	log "github.com/gonet2/libs/nsq-logger"
	. "github.com/gonet2/libs/services/proto"
)

import (
	. "types"
)

var (
	ERROR_NOT_AUTHORIZED = errors.New("User not authorized")
)

// forward messages to game server
func forward(sess *Session, p []byte) error {
	frame := &Game_Frame{
		Type:    Game_Message,
		Message: p,
	}

	// validation
	if sess.Flag&SESS_AUTHORIZED == 0 {
		return ERROR_NOT_AUTHORIZED
	}

	// forward the frame to game
	if err := sess.Stream.Send(frame); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
