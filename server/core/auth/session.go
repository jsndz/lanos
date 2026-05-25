package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"
)

type SessionService struct {
	sessionRepo *SessionRepository
}

func NewSessionService(repo *SessionRepository) *SessionService {
	return &SessionService{
		sessionRepo: repo,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, userID uint, userAgent, ip string) (*Session, error) {
	token, err := GenerateToken()
	if err != nil {
		return nil, err
	}

	session := &Session{
		UserID:    userID,
		Token:     token,
		UserAgent: userAgent,
		IPAddress: ip,
		ExpiresAt: time.Now().Add(24 * 7 * time.Hour), // 7 days
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *SessionService) ValidateSession(ctx context.Context, token string) (*Session, error) {
	return s.sessionRepo.FindByToken(token)
}

func (s *SessionService) RevokeSession(ctx context.Context, token string) error {
	return s.sessionRepo.DeleteByToken(token)
}

func GenerateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
