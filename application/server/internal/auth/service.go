package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
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
		SessionId: sessionId,
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

func (s *Service) ValidateSession(r *http.Request) (api.SessionData, error) {
	cocke, err := r.Cookie("session_id")
	if err != nil {
		return api.SessionData{}, err
	}

	sessionId := cocke.Value

	sessionData, err := s.getSessionDataById(sessionId)
	if err != nil {
		return api.SessionData{}, err
	}

	return sessionData, nil
}

// For now i just leave it like this.
func (s *Service) DeleteSessionById(sessionId string) {
	s.redis.Del(context.Background(), sessionId)
}

func (s *Service) getSessionDataById(sessionId string) (api.SessionData, error) {
	sessionDataJSON, err := s.redis.Get(context.Background(), sessionId).Result()
	if err != nil {
		return api.SessionData{}, err
	}
	var sessionData api.SessionData
	// sessionDataJSON.Bytes()
	err = json.Unmarshal([]byte(sessionDataJSON), &sessionData)
	if err != nil {
		return api.SessionData{}, err
	}
	// I don't now if i need to delete data row from redis

	return sessionData, nil
}

func generateSessionId() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
