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

type Response struct {
	User         UserResponse `json:"user"`
	SessionToken string       `json:"session_token"`
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

func (s *UserService) Register(ctx context.Context, username, useremail, password string) (*Response, error) {

	if username == "" || useremail == "" || password == "" {
		return nil, errors.New("all fields are required")
	}
	hash, err := HashPassword(password, DefaultParams)
	if err != nil {
		return nil, err
	}

	user := &User{
		UserName: username,
		Email:    useremail,
		Password: hash,
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	session, err := s.SessionService.CreateSession(ctx, user.ID, "", "")
	if err != nil {
		return nil, err
	}
	return &Response{
		User: UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.UserName,
		},
		SessionToken: session,
	}, nil
}

func (s *UserService) Login(ctx context.Context, useremail, password, ip string) (*Response, error) {
	user, err := s.userRepository.FindByEmail(useremail)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if ok, err := VerifyPassword(password, user.Password); !ok || err != nil {
		return nil, errors.New("invalid credentials")
	}

	session, err := s.SessionService.CreateSession(ctx, user.ID, ip, "")
	if err != nil {
		return nil, err
	}

	return &Response{
		User: UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.UserName,
		},
		SessionToken: session,
	}, nil
}

func (s *UserService) Logout(ctx context.Context, token string) (bool, error) {
	err := s.SessionService.RevokeSession(ctx, token)
	if err != nil {
		return false, err
	}
	return true, nil
}
