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

func (el *EntitiesLogs) ApplyLogBatch(b worldlog.LogBatch) {
	for eid, tickComponents := range b.EntitiesLogs {
		for _, tickComponent := range tickComponents {
			for t := tickComponent.TickFrom; t <= tickComponent.TickTo; t++ {
				if el.Logs[t] == nil {
					el.Logs[t] = map[entity.Id][]component.Component{}
				}
				if el.Logs[t][eid] == nil {
					el.Logs[t][eid] = []component.Component{}

				}
				el.Logs[t][eid] = append(el.Logs[t][eid], tickComponent.Component)
			}
		}
	}
}
