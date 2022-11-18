package chat

import "fmt"

type chat struct {
	name      name
	memberIDs []uint64
}

func NewChat(name string, userIDs []uint64) (*chat, error) {
	nameVO, err := newName(name)
	if err != nil {
		return nil, err
	}

	return &chat{
		name:      nameVO,
		memberIDs: userIDs,
	}, nil
}

type name string

func newName(n string) (name, error) {
	if n == "" {
		return "", fmt.Errorf("name must not be empty")
	}
	return name(n), nil
}

func (c *chat) EditName(n string) error {
	var err error
	c.name, err = newName(n)
	if err != nil {
		return err
	}
	return nil
}

func (c *chat) AppendMembers(userIDs []uint64) {
	c.memberIDs = append(c.memberIDs, userIDs...)
}

func (c *chat) DeleteMembers(userIDs []uint64) {
	isDeleted := make(map[uint64]bool)
	var newMemberIDs []uint64

	for _, uid := range userIDs {
		isDeleted[uid] = true
	}
	for _, mid := range c.memberIDs {
		if !isDeleted[mid] {
			newMemberIDs = append(newMemberIDs, mid)
		}
	}
}

type Repository interface {
	FindByID(chatID uint64) (*chat, error)
	Create(chat chat) error
}
