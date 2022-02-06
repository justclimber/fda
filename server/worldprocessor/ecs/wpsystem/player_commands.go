package wpsystem

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type playerCs struct {
	PowerSettable wpcomponent.PowerSettable
	PlayerC       *wpcomponent.Player
	delayLeft     int
}

type PlayerCommands struct {
	components map[entity.Id]*playerCs
	debugger   *debugger.Nested
}

func NewPlayerCommands(debugger *debugger.Nested) *PlayerCommands {
	return &PlayerCommands{
		components: map[entity.Id]*playerCs{},
		debugger:   debugger,
	}
}

func (p *PlayerCommands) String() string {
	return "PlayerCommands"
}

func (p *PlayerCommands) RequiredComponentKeys() []component.Key {
	return []component.Key{wpcomponent.CPlayer, wpcomponent.CMovable}
}

func (p *PlayerCommands) AddEntity(e *entity.Entity, in []interface{}) error {
	if len(in) != 2 {
		return ErrInvalidComponent
	}
	pl, ok1 := in[0].(*wpcomponent.Player)
	powerSettable, ok2 := in[1].(wpcomponent.PowerSettable)
	if !ok1 || !ok2 {
		return ErrInvalidComponent
	}

	p.components[e.Id] = &playerCs{
		PowerSettable: powerSettable,
		PlayerC:       pl,
		delayLeft:     pl.Delay,
	}
	return nil
}

func (p *PlayerCommands) RemoveEntity(_ *entity.Entity) {}

func (p *PlayerCommands) DoTick(tick tick.Tick) (error, bool) {
	for _, cs := range p.components {
		cs.delayLeft--
		//p.debugger.LogF("DoTick", "[tick %d], delayLeft: %d", tick, cs.delayLeft)
		if cs.delayLeft > 0 {
			continue
		}
		cs.delayLeft = cs.PlayerC.Delay

		select {
		case cmd := <-cs.PlayerC.CmdCh:
			p.debugger.LogF("DoTick", "get commands")
			cs.PowerSettable.SetPower(cmd.Move)
			//default:
		}
	}

	return nil, false
}
