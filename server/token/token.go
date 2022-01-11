package token

import (
	"fmt"
	"math/rand"
)

const Length = 32

type Generator interface {
	Generate() string
}

type SimpleTokenGenerator struct{}

func (s *SimpleTokenGenerator) Generate() string {
	b := make([]byte, Length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
