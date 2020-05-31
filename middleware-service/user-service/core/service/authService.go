package service

import (
	"crypto/rsa"
	"errors"

	"user-service/core/meta"
	"user-service/core/model"
)

var (
	ErrPasswordInvalid       = errors.New("user-service/authService: invalid user password")
	ErrNotConfirmed          = errors.New("user-service/authService: user email not confirmed")
	ErrInvalidRemoteID       = errors.New("user-service/authService: invalid remote OAuth identifier")
	ErrInvalidRemoteProvider = errors.New("user-service/authService: invalid remote OAuth provider")
)

type AuthService interface {
	SingUp(user *model.User) (*meta.AuthToken, error)
	Login(user *model.User) (*meta.AuthToken, error)
	Remote(user *model.User) (*meta.AuthToken, error)
	Guest() (*meta.AuthToken, error)
	GenerateToken(user *model.User) (*meta.AuthToken, error)
	PublicKey() *rsa.PublicKey
	Logout(token string) error
}
