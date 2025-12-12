package app

import (
	"log/slog"
	"net/http"
	chatCreate "project-go/internal/http-server/handlers/chat/create"
	GetChatBySessionId "project-go/internal/http-server/handlers/chat/getBySessionId"
	GetAllSessions "project-go/internal/http-server/handlers/session/getAll"
	TestCreate "project-go/internal/http-server/handlers/test/test"
	UserCreate "project-go/internal/http-server/handlers/user/create"
	UserLogin "project-go/internal/http-server/handlers/user/login"
	"project-go/internal/http-server/repository/store"
	chatservice "project-go/internal/http-server/service/chat"
	sessionService "project-go/internal/http-server/service/session"
	testservice "project-go/internal/http-server/service/test"
	userservice "project-go/internal/http-server/service/user"
	"project-go/internal/lib/jwt"
)

type App struct {
	CreateChatHandler         http.HandlerFunc
	GetChatBySessionIdHandler http.HandlerFunc
	GetAllSessions            http.HandlerFunc
	TestCreate                http.HandlerFunc
	UserCreate                http.HandlerFunc
	UserLogin                 http.HandlerFunc
}

func New(log *slog.Logger, store *store.Store, jwtKey string) *App {
	// репозитории
	chatRepo := store.ChatRepo
	sessionRepo := store.SessionRepo
	testRepo := store.TestRepo
	questionRepo := store.QuestionRepo
	userRepo := store.UserRepo
	// сервисы
	chatService := chatservice.New(chatRepo, sessionRepo)
	sessionService := sessionService.New(sessionRepo)
	testService := testservice.New(testRepo, questionRepo)
	userService := userservice.New(userRepo)
	// хендлеры
	CreateChatHandler := chatCreate.New(log, chatService)
	GetChatBySessionIdHandler := GetChatBySessionId.New(log, chatService)
	GetAllSessions := GetAllSessions.New(log, sessionService)
	TestCreate := TestCreate.New(log, testService)
	UserCreate := UserCreate.New(log, userService)
	UserLogin := UserLogin.New(log, userService, jwt.NewJWTService(jwtKey))
	return &App{
		CreateChatHandler:         CreateChatHandler,
		GetChatBySessionIdHandler: GetChatBySessionIdHandler,
		GetAllSessions:            GetAllSessions,
		TestCreate:                TestCreate,
		UserCreate:                UserCreate,
		UserLogin:                 UserLogin,
	}
}
