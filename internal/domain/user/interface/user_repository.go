package user_interface

import "comparei-servico-logs/internal/domain/user"

type UserRepository interface {
	CreateUser(user *user.User) error
}
