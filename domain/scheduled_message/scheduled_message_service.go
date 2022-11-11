package scheduled_message

import (
	"ddd_challenge/domain/chat"
	"ddd_challenge/domain/user"
	"fmt"
)

type scheduledMessageService struct {
	scheduledMessageRepository Repository
	userRepository             user.Repository
	chatRepository             chat.Repository
}

func NewScheduledMessageService(s Repository, u user.Repository, c chat.Repository) *scheduledMessageService {
	return &scheduledMessageService{
		scheduledMessageRepository: s,
		userRepository:             u,
		chatRepository:             c,
	}
}

func (m *scheduledMessageService) ScheduleSendingMessage(message scheduledMessage) error {
	userID := message.userID
	u, err := m.userRepository.FindByID(userID)
	if err != nil {
		return err
	}
	if u == nil {
		return fmt.Errorf("this user is not exists. user id is %d", message.userID)
	}

	chatID := message.chatID
	c, err := m.chatRepository.FindByID(chatID)
	if err != nil {
		return err
	}
	if c == nil {
		return fmt.Errorf("this chat is not exists. chat id is %d", chatID)
	}

	if err := m.scheduledMessageRepository.ScheduleSendingMessage(message); err != nil {
		return err
	}
	return nil
}
