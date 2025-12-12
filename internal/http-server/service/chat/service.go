package chatservice

import (
	"errors"
	"project-go/internal/models"
)

type ChatRepository interface {
	CreateChat(chat *models.ChatHistory) (*models.ChatHistory, error)
	GetChatBySessionId(sessionId uint) ([]models.ChatHistory, error)
}

type SessionRepository interface {
	CreateSession(session *models.SessionHistory) (*models.SessionHistory, error)
}
type Service struct {
	chatRepo    ChatRepository
	sessionRepo SessionRepository
}

func New(chatRepo ChatRepository, sessionRepo SessionRepository) *Service {
	return &Service{
		chatRepo:    chatRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *Service) CreateMessage(userID uint, sessionID *uint, message string) (*models.ChatHistory, error) {
	if message == "" {
		return nil, errors.New("message cannot be empty")
	}

	if sessionID == nil {
		session := &models.SessionHistory{
			StudentID: userID,
			Title:     "Dragon history",
		}
		newSession, err := s.sessionRepo.CreateSession(session)
		if err != nil {
			return nil, err
		}
		sessionID = &newSession.ID
	}

	chat := &models.ChatHistory{
		SessionID:       *sessionID,
		MessageFromUser: message,
		MessageFromBot:  "Hello world",
	}

	return s.chatRepo.CreateChat(chat)
}

func (s *Service) GetChatBySessionId(sessionId uint) ([]models.ChatHistory, error) {
	chats, err := s.chatRepo.GetChatBySessionId(sessionId)
	if err != nil {
		return nil, err
	}
	return chats, nil
}
