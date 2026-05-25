package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"lanos/pkg/util"
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

func (s *SessionService) CreateSession(ctx context.Context, userID uint, userAgent, ip string) (string, error) {
	token, err := GenerateToken()
	if err != nil {
		return "", err
	}
	hashedToken := util.HashTokenWithSha256(token)
	session := &Session{
		UserID:    userID,
		Token:     hashedToken,
		UserAgent: userAgent,
		IPAddress: ip,
		ExpiresAt: time.Now().Add(24 * 7 * time.Hour), // 7 days
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return "", err
	}

	return token, nil
}

func (s *SessionService) RevokeSession(ctx context.Context, token string) error {
	hashedToken := util.HashTokenWithSha256(token)
	return s.sessionRepo.DeleteByToken(hashedToken)
}

func (s *SessionService) GetSessionByToken(ctx context.Context, token string) (*Session, error) {
	hashedToken := util.HashTokenWithSha256(token)
	return s.sessionRepo.FindByToken(hashedToken)
}
func GenerateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
