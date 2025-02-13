package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/roka-crew/pkg/config"
)

type Token struct {
	cfg *config.Config
}

func NewToken(cfg *config.Config) *Token {
	return &Token{cfg: cfg}
}

func (t *Token) GenerateToken(userID int) (string, error) {
	payload := Payload{
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(t.cfg.Token.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *Token) ParseToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(_ *jwt.Token) (any, error) {
		return t.cfg.Token.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	return payload, nil
}

type Payload struct {
	// 사용자의 ID
	UserID int `json:"userID"`
}

func (p Payload) Valid() error {
	return nil
}
