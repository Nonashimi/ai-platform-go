package req

type TestRequest struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Questions   []QuestionRequest `json:"questions"`
}

type QuestionRequest struct {
	Question string          `json:"question"`
	Options  []OptionRequest `json:"options"`
}

type OptionRequest struct {
	OptionText string `json:"optionText"`
	IsCorrect  bool   `json:"isCorrect"`
}
