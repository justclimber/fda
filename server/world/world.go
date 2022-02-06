package world

import (
	"github.com/justclimber/fda/common/ecs/entity"
)

type World struct{}

func NewWorld() *World {
	return &World{}
}

func (w *World) RegisterNewEntity(*entity.Entity) error {
	return nil
}

func (w *World) AllocateEntity(*entity.Entity) error {
	return nil
}
