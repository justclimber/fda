package playersprocessor

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldlog"
)

type EntitiesLogs struct {
	Logs map[tick.Tick]map[entity.Id][]component.Component
}

func NewEntitiesLogs() *EntitiesLogs {
	return &EntitiesLogs{
		Logs: map[tick.Tick]map[entity.Id][]component.Component{},
	}
}

func (el *EntitiesLogs) ApplyLogBatch(b worldlog.Batch) {
	for eid, repeatableComponents := range b.Repeatable {
		for _, rc := range repeatableComponents {
			for t := rc.TickFrom; t <= rc.TickTo; t++ {
				if el.Logs[t] == nil {
					el.Logs[t] = map[entity.Id][]component.Component{}
				}
				if el.Logs[t][eid] == nil {
					el.Logs[t][eid] = []component.Component{}

				}
				el.Logs[t][eid] = append(el.Logs[t][eid], rc.Component)
			}
		}
	}
}
