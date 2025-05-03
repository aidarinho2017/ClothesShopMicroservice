package service

import (
	"awesomeProject2/internal/repository"
	"awesomeProject2/models"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *models.User) error {
	return s.repo.Register(user)
}

func (s *UserService) Login(username, password string) (string, error) {
	return s.repo.Login(username, password)
}
