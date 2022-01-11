package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/hasher"
	"github.com/justclimber/fda/common/hasher/bcrypt"
	"github.com/justclimber/fda/server/token"
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
			check, err := u.CheckPassword(tt.rawPass, tt.h)
			require.NoError(t, err)
			assert.Equalf(t, tt.want, check, "CheckPassword(%v)", tt.rawPass)
		})
	}
}

type tokenGeneratorMock struct {
	token string
}

func (t tokenGeneratorMock) Generate() string {
	return t.token
}

func TestUser_GenerateToken(t *testing.T) {
	g := &tokenGeneratorMock{token: "asd"}
	u := &User{
		Id:   1,
		Name: "Foo",
	}
	tests := []struct {
		name      string
		generator token.Generator
		want      string
	}{
		{
			name:      "simple mock gen",
			generator: g,
			want:      "asd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, u.GenerateToken(tt.generator), "GenerateToken(%v)", tt.generator)
		})
	}
}
