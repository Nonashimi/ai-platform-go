package res

import "time"

type TestResponse struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Questions   []QuestionResponse `json:"questions"`
	CreatedAt   time.Time          `json:"createdAt"`
}

type QuestionResponse struct {
	Question string           `json:"question"`
	Options  []OptionResponse `json:"options"`
}

type OptionResponse struct {
	ID         uint   `json:"id"`
	OptionText string `json:"optionText"`
	IsCorrect  bool   `json:"isCorrect"`
}
