package wpsystem

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/tick"
	servcomponent2 "github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type playerCs struct {
	PowerSettable servcomponent2.PowerSettable
	PlayerC       *servcomponent2.Player
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
	return []ecs.ComponentKey{servcomponent2.CPlayer, servcomponent2.CMovable}
}

func (p *PlayerCommands) AddEntity(e *ecs.Entity, in []interface{}) error {
	if len(in) != 2 {
		return ErrInvalidComponent
	}
	pl, ok1 := in[0].(*servcomponent2.Player)
	powerSettable, ok2 := in[1].(servcomponent2.PowerSettable)
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
