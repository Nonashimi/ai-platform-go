package res

import "time"

type ChatResponse struct {
	ID              uint      `json:"id"`
	SessionID       uint      `json:"session_id"`
	SessionTitle    string    `json:"session_title"`
	MessageFromUser string    `json:"message_from_user"`
	MessageFromBot  string    `json:"message_from_bot"`
	CreatedAt       time.Time `json:"created_at"`
}
