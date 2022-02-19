package playersprocessor

import (
	"log"

	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldlog"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type EntitiesState struct {
	store map[tick.Tick]map[entity.Id]map[component.Key]component.Component
}

func NewEntitiesState() *EntitiesState {
	return &EntitiesState{
		store: map[tick.Tick]map[entity.Id]map[component.Key]component.Component{},
	}
}

func (es *EntitiesState) ApplyLogBatch(b worldlog.Batch) {
	var maxTick tick.Tick
	for eid, repeatableComponents := range b.Repeatable {
		for _, rc := range repeatableComponents {
			for t := rc.TickFrom; t <= rc.TickTo; t++ {
				if es.store[t] == nil {
					es.store[t] = map[entity.Id]map[component.Key]component.Component{}
				}
				if es.store[t][eid] == nil {
					es.store[t][eid] = map[component.Key]component.Component{}
				}
				es.store[t][eid][rc.Component.Key()] = rc.Component
			}
			if rc.TickTo > maxTick {
				maxTick = rc.TickTo
			}
		}
	}

	for eid, calculated := range b.Calculated {
		for _, cc := range calculated {
			// todo: should be polymorph? but lets start as is
			if cc.Component.Key() == wpcomponent.KeyPosition {
				es.calculatePosition(eid, cc.Component.(wpcomponent.Position), cc.TickFrom, maxTick)
			}
		}
	}
}

func (es *EntitiesState) calculatePosition(id entity.Id, p wpcomponent.Position, tickFrom tick.Tick, tickTo tick.Tick) {
	es.store[tickFrom][id][wpcomponent.KeyPosition] = p

	if tickFrom == tickTo {
		return
	}

	for t := tickFrom; t < tickTo; t++ {
		mc, ok := es.store[t][id][wpcomponent.KeyMoving]
		m, ok2 := mc.(wpcomponent.Moving)
		if !ok || !ok2 {
			log.Fatalf("moving component must exist for tick %d and eid %d", t, id)
		}
		p = wpcomponent.Position{Pos: p.Pos.Add(m.D)}
		es.store[t+1][id][wpcomponent.KeyPosition] = p
	}
}
