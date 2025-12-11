package jwt

import (
	"project-go/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	Key []byte
}

func NewJWTService(key string) *JWTService {
	return &JWTService{Key: []byte(key)}
}

func (s *JWTService) GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.Key)
}
