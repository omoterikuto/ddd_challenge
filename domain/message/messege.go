package message

import (
	"fmt"
	"time"
)

type message struct {
	id        uint64
	text      text
	chatID    uint64
	userID    uint64
	createdAt time.Time
}

type text string

func newText(t string) (text, error) {
	if t == "" {
		return text(""), fmt.Errorf("text must not be empty")
	}
	return text(t), nil
}

func NewMessage(text string, chatID, userID uint64) (*message, error) {
	textVO, err := newText(text)
	if err != nil {
		return nil, err
	}

	return &message{
		text:   textVO,
		chatID: chatID,
		userID: userID,
	}, nil
}

type Repository interface {
	SendMessage(m message) error
}
