package user_interface

import "comparei-servico-logs/internal/domain/user"

type UserRepository interface {
	CreateUser(user *user.User) error
	GetUserById(id string) (*user.User, error)
	UpdateUserScore(user_id string, new_score float32) (*user.User, error)
	UpdateUserLevel(user_id string, new_level int) (*user.User, error)
}
