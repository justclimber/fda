package server

import (
	"context"
	"fmt"

	"github.com/justclimber/fda/server/user"
)

type authHelper struct {
	usersFinder user.Finder
}

func NewAuthHelper(usersFinder user.Finder) *authHelper {
	return &authHelper{usersFinder: usersFinder}
}

func (a *authHelper) GetUserFromContext(ctx context.Context) (*user.User, error) {
	id, ok := ctx.Value(ContextUserIdKey).(uint64)
	if !ok {
		return nil, fmt.Errorf("empty auth token")
	}
	u, err := a.usersFinder.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("can't get user with auth token")
	}
	return u, nil
}
