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
}

func (s *scheduledMessage) ID() uint64 {
	return s.id
}

type text string

func newText(t string) (text, error) {
	if t == "" {
		return "", fmt.Errorf("text must not be empty")
	}
	return text(t), nil
}

func (s scheduledMessage) ChangeText(text string) error {
	var err error
	s.text, err = newText(text)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduledMessage) Text() text {
	return s.text
}

func (s *scheduledMessage) ChatID() uint64 {
	return s.chatID
}

func (s *scheduledMessage) UserID() uint64 {
	return s.userID
}

type scheduledSendingTime time.Time

func newScheduledSendingTime(t time.Time) (scheduledSendingTime, error) {
	if t.Before(time.Now()) {
		return scheduledSendingTime(time.Now()), fmt.Errorf("scheduledSendingTime must not be before now")
	}
	return scheduledSendingTime(t), nil
}

func (s *scheduledMessage) ScheduledSendingTime() scheduledSendingTime {
	return s.scheduledSendingTime
}

func (s scheduledMessage) ChangeScheduledSendingTime(scheduledSendingTime time.Time) error {
	var err error
	s.scheduledSendingTime, err = newScheduledSendingTime(scheduledSendingTime)
	if err != nil {
		return err
	}
	return nil
}

func NewScheduledMessage(id uint64, text string, chatID, userID uint64, scheduledSendingTime time.Time) (*scheduledMessage, error) {
	scheduledSendingTimeVO, err := newScheduledSendingTime(scheduledSendingTime)
	if err != nil {
		return nil, err
	}

	textVO, err := newText(text)
	if err != nil {
		return nil, err
	}

	return &scheduledMessage{
		id:                   id,
		text:                 textVO,
		chatID:               chatID,
		userID:               userID,
		scheduledSendingTime: scheduledSendingTimeVO,
	}, nil
}

type Repository interface {
	ScheduleSendingMessage(s scheduledMessage) error
	GetByUserID(userID uint64) ([]scheduledMessage, error)
	DeleteByID(scheduledMessageID uint64) error
	Update(s scheduledMessage) error
}
