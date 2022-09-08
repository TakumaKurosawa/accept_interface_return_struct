package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) GetUserByID(id string) *User {
	return &User{
		ID: id,
	}
}
