package auth

import (
	"context"
	"errors"
)

type UserService struct {
	userRepository UserRepository
	SessionService *SessionService
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginResponse struct {
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	SessionId    string       `json:"session_id"`
}

func NewUserService(
	repo UserRepository,
	sessionService *SessionService,
) *UserService {
	return &UserService{
		userRepository: repo,
		SessionService: sessionService,
	}
}

func (s *UserService) Register(ctx context.Context, username, useremail, password string) (string, error) {

	if username == "" || useremail == "" || password == "" {
		return "", errors.New("all fields are required")
	}
	hash, err := HashPassword(password, DefaultParams)
	if err != nil {
		return "", err
	}

	user := &User{
		UserName: username,
		Email:    useremail,
		Password: hash,
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return "", err
	}

	session, err := s.SessionService.CreateSession(ctx, user.ID, "", "")
	if err != nil {
		return "", err
	}

	return session.Token, nil
}

func (s *UserService) Login(ctx context.Context, useremail, password, ip string) (*LoginResponse, error) {
	return nil, nil
}
