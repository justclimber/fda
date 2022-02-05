package servsystem

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/ecs/servcomponent"
)

type playerCs struct {
	PowerSettable servcomponent.PowerSettable
	PlayerC       *servcomponent.Player
	delayLeft     int
}

type PlayerCommands struct {
	components map[ecs.EntityId]*playerCs
	debugger   *debugger.Nested
}

func NewPlayerCommands(debugger *debugger.Nested) *PlayerCommands {
	return &PlayerCommands{
		components: map[ecs.EntityId]*playerCs{},
		debugger:   debugger,
	}
}

func (p *PlayerCommands) String() string {
	return "PlayerCommands"
}

func (p *PlayerCommands) RequiredComponentKeys() []ecs.ComponentKey {
	return []ecs.ComponentKey{servcomponent.CPlayer, servcomponent.CMovable}
}

func (p *PlayerCommands) AddEntity(e *ecs.Entity, in []interface{}) error {
	if len(in) != 2 {
		return ErrInvalidComponent
	}
	pl, ok1 := in[0].(*servcomponent.Player)
	powerSettable, ok2 := in[1].(servcomponent.PowerSettable)
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

func (p *PlayerCommands) RemoveEntity(_ *ecs.Entity) {}

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
