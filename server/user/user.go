package user

import (
	"errors"
	"fmt"

	"github.com/justclimber/fda/common/hasher"
)

var NilHasherProvided = errors.New("nil hasher provided")
var EmptyPasswordProvided = errors.New("empty password provided")

type User struct {
	Id           uint64
	Name         string
	passwordHash string
}

type ForClient struct {
	Id   uint64
	Name string
}

func NewUser(id uint64, name, password string, h hasher.Hasher) (*User, error) {
	if password == "" {
		return nil, EmptyPasswordProvided
	}

	passwordHash, err := h.Hash(password)

	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	return &User{
		Id:           id,
		Name:         name,
		passwordHash: passwordHash,
	}, nil
}

func (u *User) CheckPassword(p string, h hasher.Hasher) (bool, error) {
	if h == nil {
		return false, NilHasherProvided
	}
	return h.IsValid(u.passwordHash, p), nil
}
