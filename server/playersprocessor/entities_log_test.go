package playersprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldlog"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

func TestEntitiesLogs_ApplyLogBatch(t *testing.T) {
	tests := []struct {
		name     string
		logs     worldlog.LogBatch
		wantLogs map[tick.Tick]map[entity.Id][]component.Component
	}{
		{
			name:     "empty",
			logs:     worldlog.LogBatch{},
			wantLogs: map[tick.Tick]map[entity.Id][]component.Component{},
		},
		{
			name: "1_comp_1_tick",
			logs: worldlog.LogBatch{
				StartTick: 10,
				EndTick:   10,
				EntitiesLogs: map[entity.Id][]worldlog.TickComponent{
					10: {
						{
							TickFrom:  1,
							TickTo:    1,
							Component: wpcomponent.NewMoving(fgeom.Point{X: 1}),
						},
					},
				},
			},
			wantLogs: map[tick.Tick]map[entity.Id][]component.Component{
				1: {
					10: []component.Component{wpcomponent.NewMoving(fgeom.Point{X: 1})},
				},
			},
		},
		{
			name: "2_comp_1_tick",
			logs: worldlog.LogBatch{
				StartTick: 10,
				EndTick:   10,
				EntitiesLogs: map[entity.Id][]worldlog.TickComponent{
					10: {
						{
							TickFrom:  1,
							TickTo:    1,
							Component: wpcomponent.NewMoving(fgeom.Point{X: 1}),
						},
						{
							TickFrom:  1,
							TickTo:    1,
							Component: wpcomponent.NewPosition(fgeom.Point{X: 2}),
						},
					},
				},
			},
			wantLogs: map[tick.Tick]map[entity.Id][]component.Component{
				1: {
					10: []component.Component{
						wpcomponent.NewMoving(fgeom.Point{X: 1}),
						wpcomponent.NewPosition(fgeom.Point{X: 2}),
					},
				},
			},
		},
		{
			name: "2_comp_2_tick",
			logs: worldlog.LogBatch{
				StartTick: 10,
				EndTick:   10,
				EntitiesLogs: map[entity.Id][]worldlog.TickComponent{
					10: {
						{
							TickFrom:  1,
							TickTo:    1,
							Component: wpcomponent.NewMoving(fgeom.Point{X: 1}),
						},
						{
							TickFrom:  2,
							TickTo:    2,
							Component: wpcomponent.NewPosition(fgeom.Point{X: 2}),
						},
					},
				},
			},
			wantLogs: map[tick.Tick]map[entity.Id][]component.Component{
				1: {10: []component.Component{wpcomponent.NewMoving(fgeom.Point{X: 1})}},
				2: {10: []component.Component{wpcomponent.NewPosition(fgeom.Point{X: 2})}},
			},
		},
		{
			name: "2_comp_2_tick_2_entities",
			logs: worldlog.LogBatch{
				StartTick: 10,
				EndTick:   10,
				EntitiesLogs: map[entity.Id][]worldlog.TickComponent{
					10: {
						{
							TickFrom:  1,
							TickTo:    1,
							Component: wpcomponent.NewMoving(fgeom.Point{X: 1}),
						},
						{
							TickFrom:  2,
							TickTo:    2,
							Component: wpcomponent.NewPosition(fgeom.Point{X: 2}),
						},
					},
					20: {
						{
							TickFrom:  2,
							TickTo:    2,
							Component: wpcomponent.NewPosition(fgeom.Point{X: 5}),
						},
					},
				},
			},
			wantLogs: map[tick.Tick]map[entity.Id][]component.Component{
				1: {10: []component.Component{wpcomponent.NewMoving(fgeom.Point{X: 1})}},
				2: {
					10: []component.Component{wpcomponent.NewPosition(fgeom.Point{X: 2})},
					20: []component.Component{wpcomponent.NewPosition(fgeom.Point{X: 5})},
				},
			},
		},
		{
			name: "1_comp_1-3_tick",
			logs: worldlog.LogBatch{
				StartTick: 1,
				EndTick:   3,
				EntitiesLogs: map[entity.Id][]worldlog.TickComponent{
					10: {
						{
							TickFrom:  1,
							TickTo:    3,
							Component: wpcomponent.NewMoving(fgeom.Point{X: 1}),
						},
					},
				},
			},
			wantLogs: map[tick.Tick]map[entity.Id][]component.Component{
				1: {10: []component.Component{wpcomponent.NewMoving(fgeom.Point{X: 1})}},
				2: {10: []component.Component{wpcomponent.NewMoving(fgeom.Point{X: 1})}},
				3: {10: []component.Component{wpcomponent.NewMoving(fgeom.Point{X: 1})}},
			},
		},
		{
			name: "mixed_1",
			logs: worldlog.LogBatch{
				EntitiesLogs: map[entity.Id][]worldlog.TickComponent{
					10: {
						{
							TickFrom:  1,
							TickTo:    3,
							Component: wpcomponent.NewMoving(fgeom.Point{X: 1}),
						},
						{
							TickFrom:  2,
							TickTo:    4,
							Component: wpcomponent.NewPosition(fgeom.Point{X: 2}),
						},
					},
					20: {
						{
							TickFrom:  2,
							TickTo:    3,
							Component: wpcomponent.NewMoving(fgeom.Point{X: 1}),
						},
					},
				},
			},
			wantLogs: map[tick.Tick]map[entity.Id][]component.Component{
				1: {10: []component.Component{wpcomponent.NewMoving(fgeom.Point{X: 1})}},
				2: {
					10: []component.Component{
						wpcomponent.NewMoving(fgeom.Point{X: 1}),
						wpcomponent.NewPosition(fgeom.Point{X: 2}),
					},
					20: []component.Component{
						wpcomponent.NewMoving(fgeom.Point{X: 1}),
					},
				},
				3: {
					10: []component.Component{
						wpcomponent.NewMoving(fgeom.Point{X: 1}),
						wpcomponent.NewPosition(fgeom.Point{X: 2}),
					},
					20: []component.Component{
						wpcomponent.NewMoving(fgeom.Point{X: 1}),
					},
				},
				4: {
					10: []component.Component{
						wpcomponent.NewPosition(fgeom.Point{X: 2}),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			el := NewEntitiesLogs()
			el.ApplyLogBatch(tt.logs)
			assert.Equal(t, tt.wantLogs, el.Logs)
		})
	}
}
