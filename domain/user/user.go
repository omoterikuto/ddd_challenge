package user

type user struct {
	id   uint64
	name string
}

type Repository interface {
	FindByID(chatID uint64) (*user, error)
}
