package auth

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		db: DB,
	}
}

func (r *UserRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*User, error) {
	var user User

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	var user User

	err := r.db.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(DB *gorm.DB) *SessionRepository {
	return &SessionRepository{
		db: DB,
	}
}

func (r *SessionRepository) Create(session *Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepository) FindByToken(token string) (*Session, error) {
	var session Session
	err := r.db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepository) DeleteByToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&Session{}).Error
}

func (r *SessionRepository) DeleteExpired() error {
	return r.db.Where("expires_at <= ?", time.Now()).Delete(&Session{}).Error
}
