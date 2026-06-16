package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"shop/internal/api"
	"time"

	"github.com/go-redis/redis/v8"
)

const SESSION_DURATION = 30 * time.Minute

type Service struct {
	redis *redis.Client
}

func NewService(redis *redis.Client) *Service {
	return &Service{redis: redis}
}

func (s *Service) CreateSession(user *api.User) (string, time.Time, error) {
	sessionId, err := generateSessionId()
	if err != nil {
		return "", time.Time{}, err
	}

	timeNow := time.Now()
	expiredAt := timeNow.Add(SESSION_DURATION)

	sessionData := api.SessionData{
		UserID:    user.Id,
		Email:     user.Email,
		CreatedAt: timeNow,
		ExpiresAt: expiredAt,
	}

	sessionJSON, err := json.Marshal(sessionData)
	if err != nil {
		return "", time.Time{}, err
	}

	err = s.redis.Set(
		context.Background(),
		sessionId,
		sessionJSON,
		SESSION_DURATION).Err()
	if err != nil {
		return "", time.Time{}, err
	}

	return sessionId, expiredAt, nil
}

func generateSessionId() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
