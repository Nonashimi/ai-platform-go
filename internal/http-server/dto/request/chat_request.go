package req

type CreateChatRequest struct {
	Message   string `json:"message"`
	SessionId *uint  `json:"session_id"`
}
