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

func (s *UserService) GetUserById(id string) (*user.User, error) {
	return s.mysqlRepo.GetUserById(id)
}

func (s *UserService) UpdateUserScore(id string, new_score float32) (*user.User, error) {
	return s.mysqlRepo.UpdateUserScore(id, new_score)
}

func (s *UserService) UpdateUserLevel(id string, new_level int) (*user.User, error) {
	return s.mysqlRepo.UpdateUserLevel(id, new_level)
}
