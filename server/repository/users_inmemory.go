package repository

import (
	"fmt"

	"github.com/justclimber/fda/server/user"
)

type UsersInMemory struct {
	users          map[uint64]*user.User
	userIdsByName  map[string]uint64
	userIdsByToken map[string]uint64
	lastUserId     uint64
}

func NewUsersInMemory() *UsersInMemory {
	return &UsersInMemory{
		users:          make(map[uint64]*user.User),
		userIdsByName:  make(map[string]uint64),
		userIdsByToken: make(map[string]uint64),
		lastUserId:     0,
	}
}

func (u *UsersInMemory) Register(t user.ToRegister) (uint64, error) {
	if _, found := u.userIdsByName[t.Name]; found {
		return 0, user.ErrAlreadyExists
	}
	u.lastUserId++
	id := u.lastUserId
	newUser, err := user.NewUser(id, t.Name, t.PasswordHash)

	if err != nil {
		return 0, fmt.Errorf("can't create user obj: %w", err)
	}

	u.users[id] = newUser
	u.userIdsByName[t.Name] = id
	return id, nil
}

func (u *UsersInMemory) FindById(id uint64) (*user.User, error) {
	usr, ok := u.users[id]
	if !ok {
		return nil, nil
	}
	return usr, nil
}

func (u *UsersInMemory) FindByToken(token string) (*user.User, error) {
	id, ok := u.userIdsByToken[token]
	if !ok {
		return nil, nil
	}
	return u.FindById(id)
}

func (u *UsersInMemory) StoreToken(id uint64, token string) error {
	u.userIdsByToken[token] = id
	return nil
}
