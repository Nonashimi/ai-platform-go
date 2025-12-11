package chatCreate

import (
	"log/slog"
	"net/http"
	req "project-go/internal/http-server/dto/request"
	res "project-go/internal/http-server/dto/response"
	"project-go/internal/lib/api/response"
	"project-go/internal/lib/auth"
	"project-go/internal/models"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type ChatCreate interface {
	CreateChat(chat *models.ChatHistory) (*models.ChatHistory, error)
}

type SessionCreate interface {
	CreateSession(session *models.SessionHistory) (*models.SessionHistory, error)
}

type Response struct {
	response.Response
	Chat res.ChatResponse
}

func New(log *slog.Logger, ChatCreate ChatCreate, SessionCreate SessionCreate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.chat.create.New"
		log = log.With(
			slog.String("op", op),
			slog.String("req_id", middleware.GetReqID(r.Context())),
		)
		var req req.CreateChatRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request", slog.String("error", err.Error()))
		}
		log.Info("request body decoded", slog.Any("req body:", req))
		userID, idOk := auth.GetUserID(r)
		userRole, _ := auth.GetRole(r)
		log.Info("req context values", slog.String("role", userRole), slog.Any("userID", userID))
		if !idOk || userID == 0 {
			log.Error("user id not  found", slog.Any("userId", userID))
			http.Error(w, "user id not found", http.StatusUnauthorized)
			return
		}
		if req.Message == "" {
			log.Error("message is empty")
			render.JSON(w, r, response.Error("message is empty"))
			return
		}
		sessionId := req.SessionId
		if sessionId == nil {
			session := models.SessionHistory{
				StudentID: userID,
				Title:     "Dragon history",
			}
			createdSession, err := SessionCreate.CreateSession(&session)
			if err != nil {
				log.Error("failed to create a new session", slog.String("error", err.Error()))
				render.JSON(w, r, response.Error("failed to create a new session"))
				return
			}
			sessionId = &createdSession.ID
		}

		newChat := models.ChatHistory{
			SessionID:       *sessionId,
			MessageFromUser: req.Message,
			MessageFromBot:  "Hello world",
		}

		createdChat, err := ChatCreate.CreateChat(&newChat)
		if err != nil {
			log.Error("failed to create chat", slog.String("error", err.Error()))
			render.JSON(w, r, response.Error("failed to create chat"))
			return
		}
		render.JSON(w, r, Response{
			Response: response.OK(),
			Chat: res.ChatResponse{
				ID:              createdChat.ID,
				SessionID:       createdChat.SessionID,
				SessionTitle:    createdChat.Session.Title,
				MessageFromUser: createdChat.MessageFromUser,
				MessageFromBot:  createdChat.MessageFromBot,
				CreatedAt:       createdChat.CreatedAt,
			},
		})

	}
}
