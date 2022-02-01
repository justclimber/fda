package level

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/server/ecs/servcomponent"
)

type Level struct{}

func NewLevel() *Level {
	return &Level{}
}

func NewPlayerEntity(id ecs.EntityId, p *servcomponent.Player) *ecs.Entity {
	return &ecs.Entity{
		Id: id,
		Components: map[ecs.ComponentKey]interface{}{
			servcomponent.CPlayer: p,
		},
	}
}

func (l *Level) RegisterNewEntity(*ecs.Entity) error {
	return nil
}

func (l *Level) AllocateEntity(*ecs.Entity) error {
	return nil
}
