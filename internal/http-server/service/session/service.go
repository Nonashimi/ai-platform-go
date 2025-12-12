package sessionService

import "project-go/internal/models"

type Service struct {
	sessionRepo SessionGetAll
}

type SessionGetAll interface {
	GetAllSessions(userId uint) ([]models.SessionHistory, error)
}

func New(sessionRepo SessionGetAll) *Service {
	return &Service{
		sessionRepo: sessionRepo,
	}
}

func (s *Service) GetAllSessions(userId uint) ([]models.SessionHistory, error) {
	sessions, err := s.sessionRepo.GetAllSessions(userId)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}
