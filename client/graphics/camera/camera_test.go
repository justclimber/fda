package camera

import (
	"testing"

	"github.com/golang/geo/r1"
	"github.com/golang/geo/r2"
	"github.com/stretchr/testify/assert"
)

func TestCamera_Offset(t *testing.T) {
	tests := []struct {
		name      string
		camera    *Camera
		args      r2.Point
		wantP     r2.Point
		wantIsOut bool
	}{
		{
			name: "simple_inbound1",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 1, Hi: 11},
				Y: r1.Interval{Lo: 3, Hi: 13},
			}, 0),
			args:      r2.Point{X: 8, Y: 6},
			wantP:     r2.Point{X: 7, Y: 3},
			wantIsOut: false,
		},
		{
			name: "cameraCentered_inbound2",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 10, Hi: 30},
				Y: r1.Interval{Lo: 15, Hi: 45},
			}, 0),
			args:      r2.Point{X: 13, Y: 20},
			wantP:     r2.Point{X: 3, Y: 5},
			wantIsOut: false,
		},
		{
			name: "pointTooRight_outOfBound1",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 2, Hi: 7},
				Y: r1.Interval{Lo: 2, Hi: 7},
			}, 0),
			args:      r2.Point{X: 8, Y: 3},
			wantP:     r2.Point{X: 0, Y: 0},
			wantIsOut: true,
		},
		{
			name: "pointTooLeft_outOfBound2",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 2, Hi: 7},
				Y: r1.Interval{Lo: 2, Hi: 7},
			}, 0),
			args:      r2.Point{X: 1, Y: 3},
			wantP:     r2.Point{X: 0, Y: 0},
			wantIsOut: true,
		},
		{
			name: "pointTooAbove_outOfBound3",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 2, Hi: 7},
				Y: r1.Interval{Lo: 2, Hi: 7},
			}, 0),
			args:      r2.Point{X: 4, Y: 1},
			wantP:     r2.Point{X: 0, Y: 0},
			wantIsOut: true,
		},
		{
			name: "pointTooBelow_outOfBound4",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 2, Hi: 7},
				Y: r1.Interval{Lo: 2, Hi: 7},
			}, 0),
			args:      r2.Point{X: 4, Y: 8},
			wantP:     r2.Point{X: 0, Y: 0},
			wantIsOut: true,
		},
		{
			name: "leftUpCorner_inBound3",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 2, Hi: 7},
				Y: r1.Interval{Lo: 2, Hi: 7},
			}, 0),
			args:      r2.Point{X: 2, Y: 2},
			wantP:     r2.Point{X: 0, Y: 0},
			wantIsOut: false,
		},
		{
			name: "rightBottomCorner_inBound4",
			camera: NewCamera(r2.Rect{
				X: r1.Interval{Lo: 2, Hi: 7},
				Y: r1.Interval{Lo: 2, Hi: 7},
			}, 0),
			args:      r2.Point{X: 7, Y: 7},
			wantP:     r2.Point{X: 5, Y: 5},
			wantIsOut: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPoint, wantIsOut := tt.camera.Offset(tt.args)
			assert.Equalf(t, tt.wantP, gotPoint, "Offset(%v)", tt.args)
			assert.Equalf(t, tt.wantIsOut, wantIsOut, "Offset(%v)", tt.args)
		})
	}
}
