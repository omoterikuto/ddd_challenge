package scheduled_message

import (
	"fmt"
	"time"
)

type scheduledMessage struct {
	id                   uint64
	text                 text
	chatID               uint64
	userID               uint64
	scheduledSendingTime scheduledSendingTime
	createdAt            time.Time
}

type scheduledSendingTime time.Time

func newScheduledSendingTime(t time.Time) (scheduledSendingTime, error) {
	if t.Before(time.Now()) {
		return scheduledSendingTime(time.Now()), fmt.Errorf("scheduledSendingTime must not be before now")
	}
	return scheduledSendingTime(t), nil
}

type text string

func newText(t string) (text, error) {
	if t == "" {
		return "", fmt.Errorf("text must not be empty")
	}
	return text(t), nil
}

func NewScheduledMessage(text string, chatID, userID uint64, scheduledSendingTime time.Time) (*scheduledMessage, error) {
	scheduledSendingTimeVO, err := newScheduledSendingTime(scheduledSendingTime)
	if err != nil {
		return nil, err
	}

	textVO, err := newText(text)
	if err != nil {
		return nil, err
	}

	return &scheduledMessage{
		text:                 textVO,
		chatID:               chatID,
		userID:               userID,
		scheduledSendingTime: scheduledSendingTimeVO,
	}, nil
}

func (s scheduledMessage) EditSendTime(scheduledSendingTime time.Time) error {
	var err error
	s.scheduledSendingTime, err = newScheduledSendingTime(scheduledSendingTime)
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
