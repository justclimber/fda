package user

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/common/hasher/bcrypt"
)

type NoHasher struct{}

func (n NoHasher) Hash(raw string) (string, error) {
	return raw, nil
}

func (n NoHasher) IsValid(hash string, raw string) bool {
	return hash == raw
}

func TestUser_CheckPassword(t *testing.T) {
	noHasher := NoHasher{}
	bc := bcrypt.Bcrypt{}
	const pass = "123"
	wrongPass := "1234"
	bcryptedPass, _ := bc.Hash(pass)

	tests := []struct {
		name         string
		h            hasher.Hasher
		passwordHash string
		rawPass      string
		want         bool
	}{
		{
			name:         "test with NoHasHer to true",
			h:            noHasher,
			passwordHash: pass,
			rawPass:      pass,
			want:         true,
		},
		{
			name:         "test with NoHasHer to false",
			h:            noHasher,
			passwordHash: pass,
			rawPass:      wrongPass,
			want:         false,
		},
		{
			name:         "test with bcrypt to true",
			h:            bc,
			passwordHash: bcryptedPass,
			rawPass:      pass,
			want:         true,
		},
		{
			name:         "test with bcrypt to false",
			h:            bc,
			passwordHash: pass,
			rawPass:      wrongPass,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Id:           1,
				Name:         "foo",
				passwordHash: tt.passwordHash,
			}
			assert.Equalf(t, tt.want, u.CheckPassword(tt.rawPass, tt.h), "CheckPassword(%v)", tt.rawPass)
		})
	}
}
