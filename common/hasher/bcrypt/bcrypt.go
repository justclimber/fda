package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// Bcrypt is a Hasher implementation for bcrypt
type Bcrypt struct{}

// Hash returns a hash for the provided string
func (h Bcrypt) Hash(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// IsValid checks if a hash and a string match
func (h Bcrypt) IsValid(hash, raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return err == nil
}
