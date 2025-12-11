package store

import (
	"project-go/internal/http-server/repository/chat"
	"project-go/internal/http-server/repository/session"
	"project-go/internal/http-server/repository/test/test"
	"project-go/internal/http-server/repository/user"

	"gorm.io/gorm"
)

type Store struct {
	UserRepo    *user.UserRepository
	SessionRepo *session.SessionRepository
	ChatRepo    *chat.ChatRepository
	TestRepo    *test.TestRepository
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		UserRepo:    user.NewUserRepo(db),
		SessionRepo: session.NewSessionRepo(db),
		ChatRepo:    chat.NewChatnRepo(db),
		TestRepo:    test.NewTestRepo(db),
	}
}
