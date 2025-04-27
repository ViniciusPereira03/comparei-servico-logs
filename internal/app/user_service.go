package app

import (
	"comparei-servico-logs/internal/domain/user"
	user_interface "comparei-servico-logs/internal/domain/user/interface"
)

type UserService struct {
	mysqlRepo user_interface.UserRepository
}

func NewUserService(mysqlRepo user_interface.UserRepository) *UserService {
	return &UserService{mysqlRepo: mysqlRepo}
}

func (s *UserService) CreateUser(user *user.User) error {
	err := s.mysqlRepo.CreateUser(user)
	return err
}
