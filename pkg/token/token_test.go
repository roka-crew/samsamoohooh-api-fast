package token

import (
	"testing"

	"github.com/roka-crew/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	cfg := &config.Config{
		Token: config.Token{
			SecretKey: []byte("test-secret-key"),
		},
	}
	tokenService := NewToken(cfg)

	userID := 12345
	tokenString, err := tokenService.GenerateToken(userID)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestParseToken(t *testing.T) {
	cfg := &config.Config{
		Token: config.Token{
			SecretKey: []byte("test-secret-key"),
		},
	}
	tokenService := NewToken(cfg)

	userID := 12345
	tokenString, err := tokenService.GenerateToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	parsedPayload, err := tokenService.ParseToken(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, parsedPayload)
	assert.Equal(t, userID, parsedPayload.UserID)
}

func TestParseToken_InvalidToken(t *testing.T) {
	cfg := &config.Config{
		Token: config.Token{
			SecretKey: []byte("test-secret-key"),
		},
	}
	tokenService := NewToken(cfg)

	invalidToken := "invalid.token.string"

	parsedPayload, err := tokenService.ParseToken(invalidToken)
	assert.Error(t, err)
	assert.Nil(t, parsedPayload)
}

func TestParseToken_WrongSecret(t *testing.T) {
	cfg := &config.Config{
		Token: config.Token{
			SecretKey: []byte("test-secret-key"),
		},
	}
	tokenService := NewToken(cfg)

	userID := 12345
	tokenString, err := tokenService.GenerateToken(userID)
	assert.NoError(t, err)

	// 다른 SecretKey를 가진 토큰 서비스
	wrongCfg := &config.Config{
		Token: config.Token{
			SecretKey: []byte("test-secret-key-another"),
		},
	}
	wrongTokenService := NewToken(wrongCfg)

	parsedPayload, err := wrongTokenService.ParseToken(tokenString)
	assert.Error(t, err)
	assert.Nil(t, parsedPayload)
}
