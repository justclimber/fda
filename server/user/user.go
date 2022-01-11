package user

import (
	"errors"
	"fmt"

	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/server/token"
)

var ErrNilHasher = errors.New("nil hasher provided")
var ErrEmptyName = errors.New("empty name provided")
var ErrEmptyPassword = errors.New("empty password provided")
var ErrAlreadyExists = errors.New("user already exists")

type User struct {
	Id           uint64
	Name         string
	passwordHash string
}

type ToRegister struct {
	Name         string
	PasswordHash string
}

type Repository interface {
	Register(u ToRegister) (uint64, error)
	StoreToken(id uint64, token string) error
	TokenFinder
	Finder
}

type Finder interface {
	FindById(id uint64) (*User, error)
}

type TokenFinder interface {
	FindByToken(token string) (*User, error)
}

func NewUserToRegister(name, password string, h hasher.Hasher) (ToRegister, error) {
	if name == "" {
		return ToRegister{}, ErrEmptyName
	}
	if password == "" {
		return ToRegister{}, ErrEmptyPassword
	}

	passwordHash, err := h.Hash(password)
	if err != nil {
		return ToRegister{}, fmt.Errorf("hash password: %w", err)
	}

	return ToRegister{
		Name:         name,
		PasswordHash: passwordHash,
	}, nil
}

func NewUser(id uint64, name, passwordHash string) (*User, error) {
	if name == "" {
		return nil, ErrEmptyName
	}

	if passwordHash == "" {
		return nil, ErrEmptyPassword
	}

	return &User{
		Id:           id,
		Name:         name,
		passwordHash: passwordHash,
	}, nil
}

func (u *User) CheckPassword(p string, h hasher.Hasher) (bool, error) {
	if h == nil {
		return false, ErrNilHasher
	}
	return h.IsValid(u.passwordHash, p), nil
}

func (u *User) GenerateToken(generator token.Generator) string {
	return generator.Generate()
}
