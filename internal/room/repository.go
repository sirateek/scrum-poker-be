package room

type roomRepository struct {
}

type Repository interface {
}

func NewRepository() Repository {
	return &roomService{}
}
