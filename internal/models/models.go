package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Role string

const (
	Admin   Role = "admin"
	Teacher Role = "teacher"
	Student Role = "student"
)

func IsValidRole(r string) bool {
	switch Role(r) {
	case Admin, Teacher, Student:
		return true
	}
	return false
}

// Пользователь
type User struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"size:100;not null"`
	Email     string  `gorm:"size:100;unique;not null"`
	Password  string  `gorm:"size:255;not null"`
	Role      Role    `gorm:"type:role_enum;not null"`                                                       // admin / teacher / student
	Students  []*User `gorm:"many2many:teachers_students;joinForeignKey:TeacherID;JoinReferences:StudentID"` // для учителей
	Teachers  []*User `gorm:"many2many:teachers_students;joinForeignKey:StudentID;JoinReferences:TeacherID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Связь учителей и учеников
type TeacherStudent struct {
	TeacherID uint `gorm:"primaryKey"`
	StudentID uint `gorm:"primaryKey"`
}

type SessionHistory struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	Student   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title     string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Chats     []ChatMessage `gorm:"foreignKey:SessionID"`
}

type ChatMessage struct {
	ID        uint `gorm:"primaryKey"`
	SessionID uint
	Session   SessionHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// "user" или "bot"
	Role string `gorm:"type:varchar(10)"`

	// Сам текст сообщения
	Content string `gorm:"type:text"`

	CreatedAt time.Time
}

// Тест
type Test struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Questions   []TestQuestion
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Вопросы теста
type TestQuestion struct {
	ID       uint `gorm:"primaryKey"`
	TestID   uint
	Test     Test         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Question string       `gorm:"type:text;not null"`
	Options  []TestOption `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Результаты теста
type TestResult struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	Student   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TestID    uint
	Test      Test `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Score     int  `gorm:"not null"`
	CreatedAt time.Time
}
type TestOption struct {
	ID             uint `gorm:"primaryKey"`
	TestQuestionID uint
	OptionText     string `gorm:"type:text;not null"`
	IsCorrect      bool   `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Карточки
type CardHolder struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	Student   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title     string `gorm:"type:text;not null"`
	Cards     []Card
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Card struct {
	ID           uint `gorm:"primaryKey"`
	CardHolderID uint
	CardHolder   CardHolder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Question     string     `gorm:"type:text;not null"`
	Answer       string     `gorm:"type:text;not null"`
}

// Игры
type Game struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:text;not null"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
}

type GameResult struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	Student   User `gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	GameID    uint
	Game      Game `gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Score     int  `gorm:"not null"`
	CreatedAt time.Time
}
