package service

import (
	"github.com/vikasd145/nfi-test-project/pkg/repository"
)

type UserService interface {
	RegisterUser() (int, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) RegisterUser() (int, error) {
	return s.userRepository.CreateUser()
}
