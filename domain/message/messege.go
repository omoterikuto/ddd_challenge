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

func (m *message) ID() uint64 {
	return m.id
}

type text string

func newText(t string) (text, error) {
	if t == "" {
		return text(""), fmt.Errorf("text must not be empty")
	}
	return text(t), nil
}

func (m *message) Text() text {
	return m.text
}

func (m *message) ChatID() uint64 {
	return m.chatID
}

func (m *message) UserID() uint64 {
	return m.userID
}

func (m *message) CreatedAt() time.Time {
	return m.createdAt
}

func NewMessage(id uint64, text string, chatID, userID uint64) (*message, error) {
	textVO, err := newText(text)
	if err != nil {
		return nil, err
	}

	return &message{
		id:     id,
		text:   textVO,
		chatID: chatID,
		userID: userID,
	}, nil
}

type Repository interface {
	SendMessage(m message) error
}
