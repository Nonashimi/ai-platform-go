package userservice

import "project-go/internal/models"

type Service struct {
	UserRepo UserRepotisory
}

type UserRepotisory interface {
	CreateUser(u *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}

func New(UserRepo UserRepotisory) *Service {
	return &Service{
		UserRepo: UserRepo,
	}
}

func (s *Service) CreateUser(user *models.User) (*models.User, error) {
	newUser, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) FindUserByEmail(email string) (*models.User, error) {
	findUser, err := s.UserRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return findUser, nil
}
