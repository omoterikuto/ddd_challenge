package user

type user struct {
	id   uint64
	name string
}

func (u *user) ID() uint64 {
	return u.id
}

func (u *user) Name() string {
	return u.name
}

type Repository interface {
	FindByID(chatID uint64) (*user, error)
}
