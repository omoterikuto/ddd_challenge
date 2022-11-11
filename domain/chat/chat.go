package chat

type chat struct {
	name      string
	memberIDs []uint64
}

func NewChat(name string, userIDs []uint64) *chat {
	return &chat{
		name:      name,
		memberIDs: userIDs,
	}
}

type Repository interface {
	FindByID(chatID uint64) (*chat, error)
	Create(chat chat) error
}
