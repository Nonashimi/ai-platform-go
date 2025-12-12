package res

import (
	"project-go/internal/models"
	"time"
)

type ChatResponse struct {
	ID              uint      `json:"id"`
	SessionID       uint      `json:"session_id"`
	SessionTitle    string    `json:"session_title"`
	MessageFromUser string    `json:"message_from_user"`
	MessageFromBot  string    `json:"message_from_bot"`
	CreatedAt       time.Time `json:"created_at"`
}

func ChatResponseFromModel(chat *models.ChatHistory) ChatResponse {
	return ChatResponse{
		ID:              chat.ID,
		SessionID:       chat.SessionID,
		SessionTitle:    chat.Session.Title,
		MessageFromUser: chat.MessageFromUser,
		MessageFromBot:  chat.MessageFromBot,
		CreatedAt:       chat.CreatedAt,
	}
}
