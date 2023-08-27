package room

type roomService struct {
}

type Service interface{}

func NewService() Service {
	return &roomService{}
}
