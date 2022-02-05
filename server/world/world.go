package world

import (
	"github.com/justclimber/fda/common/ecs"
)

type World struct{}

func NewWorld() *World {
	return &World{}
}

func (w *World) RegisterNewEntity(*ecs.Entity) error {
	return nil
}

func (w *World) AllocateEntity(*ecs.Entity) error {
	return nil
}
