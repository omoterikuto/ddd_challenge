package scheduled_message

import (
	"fmt"
	"time"
)

type scheduledMessage struct {
	id        uint64
	text      text
	chatID    uint64
	userID    uint64
	sendTime  sendTime
	createdAt time.Time
}

type sendTime time.Time

func newSendTime(t time.Time) (sendTime, error) {
	if t.Before(time.Now()) {
		return sendTime(time.Now()), fmt.Errorf("sendTime must not be before now")
	}
	return sendTime(t), nil
}

type text string

func newText(t string) (text, error) {
	if t == "" {
		return text(""), fmt.Errorf("text must not be empty")
	}
	return text(t), nil
}

func NewScheduledMessage(text string, chatID, userID uint64, sendTime time.Time) (*scheduledMessage, error) {
	sendTimeVO, err := newSendTime(sendTime)
	if err != nil {
		return nil, err
	}

	textVO, err := newText(text)
	if err != nil {
		return nil, err
	}

	return &scheduledMessage{
		text:     textVO,
		chatID:   chatID,
		userID:   userID,
		sendTime: sendTimeVO,
	}, nil
}

func (s scheduledMessage) EditSendTime(sendTime time.Time) error {
	var err error
	s.sendTime, err = newSendTime(sendTime)
	if err != nil {
		return err
	}
	return nil
}

func (s scheduledMessage) EditText(text string) error {
	var err error
	s.text, err = newText(text)
	if err != nil {
		return err
	}
	return nil
}

type Repository interface {
	ScheduleSendingMessage(s scheduledMessage) error
	GetByUserID(userID uint64) ([]scheduledMessage, error)
	DeleteByID(scheduledMessageID uint64) error
	UpdateByID(s scheduledMessage) error
}
