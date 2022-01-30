// Copyright 2014 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fgeom

import (
	"math"
	"testing"
)

var (
	sw = Point{0, 0.25}
	se = Point{0.5, 0.25}
	ne = Point{0.5, 0.75}
	nw = Point{0, 0.75}

	empty   = EmptyRect()
	rect    = RectFromPoints(sw, ne)
	rectMid = RectFromPoints(Point{0.25, 0.5}, Point{0.25, 0.5})
	rectSW  = RectFromPoints(sw, sw)
	rectNE  = RectFromPoints(ne, ne)
)

func float64Eq(x, y float64) bool { return math.Abs(x-y) < 1e-14 }

func pointsApproxEqual(a, b Point) bool {
	return float64Eq(a.X, b.X) && float64Eq(a.Y, b.Y)
}

func TestOrtho(t *testing.T) {
	tests := []struct {
		p    Point
		want Point
	}{
		{Point{0, 0}, Point{0, 0}},
		{Point{0, 1}, Point{-1, 0}},
		{Point{1, 1}, Point{-1, 1}},
		{Point{-4, 7}, Point{-7, -4}},
		{Point{1, math.Sqrt(3)}, Point{-math.Sqrt(3), 1}},
	}
	for _, test := range tests {
		if got := test.p.Ortho(); !pointsApproxEqual(got, test.want) {
			t.Errorf("%v.Ortho() = %v, want %v", test.p, got, test.want)
		}
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		p    Point
		op   Point
		want float64
	}{
		{Point{0, 0}, Point{0, 0}, 0},
		{Point{0, 1}, Point{0, 0}, 0},
		{Point{1, 1}, Point{4, 3}, 7},
		{Point{-4, 7}, Point{1, 5}, 31},
	}
	for _, test := range tests {
		if got := test.p.Dot(test.op); !float64Eq(got, test.want) {
			t.Errorf("%v.Dot(%v) = %v, want %v", test.p, test.op, got, test.want)
		}
	}
}

func TestCross(t *testing.T) {
	tests := []struct {
		p    Point
		op   Point
		want float64
	}{
		{Point{0, 0}, Point{0, 0}, 0},
		{Point{0, 1}, Point{0, 0}, 0},
		{Point{1, 1}, Point{-1, -1}, 0},
		{Point{1, 1}, Point{4, 3}, -1},
		{Point{1, 5}, Point{-2, 3}, 13},
	}

	for _, test := range tests {
		if got := test.p.Cross(test.op); !float64Eq(got, test.want) {
			t.Errorf("%v.Cross(%v) = %v, want %v", test.p, test.op, got, test.want)
		}
	}
}

func TestNorm(t *testing.T) {
	tests := []struct {
		p    Point
		want float64
	}{
		{Point{0, 0}, 0},
		{Point{0, 1}, 1},
		{Point{-1, 0}, 1},
		{Point{3, 4}, 5},
		{Point{3, -4}, 5},
		{Point{2, 2}, 2 * math.Sqrt(2)},
		{Point{1, math.Sqrt(3)}, 2},
		{Point{29, 29 * math.Sqrt(3)}, 29 * 2},
		{Point{1, 1e15}, 1e15},
		{Point{1e14, math.MaxFloat32 - 1}, math.MaxFloat32},
	}

	for _, test := range tests {
		if !float64Eq(test.p.Norm(), test.want) {
			t.Errorf("%v.Norm() = %v, want %v", test.p, test.p.Norm(), test.want)
		}
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		have Point
		want Point
	}{
		{Point{}, Point{}},
		{Point{0, 0}, Point{0, 0}},
		{Point{0, 1}, Point{0, 1}},
		{Point{-1, 0}, Point{-1, 0}},
		{Point{3, 4}, Point{0.6, 0.8}},
		{Point{3, -4}, Point{0.6, -0.8}},
		{Point{2, 2}, Point{math.Sqrt(2) / 2, math.Sqrt(2) / 2}},
		{Point{7, 7 * math.Sqrt(3)}, Point{0.5, math.Sqrt(3) / 2}},
		{Point{1e21, 1e21 * math.Sqrt(3)}, Point{0.5, math.Sqrt(3) / 2}},
		{Point{1, 1e16}, Point{0, 1}},
		{Point{1e4, math.MaxFloat32 - 1}, Point{0, 1}},
	}

	for _, test := range tests {
		if got := test.have.Normalize(); !pointsApproxEqual(got, test.want) {
			t.Errorf("%v.Normalize() = %v, want %v", test.have, got, test.want)
		}
	}

}

func TestPoint_Near(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		p2   Point
		d    float64
		want bool
	}{
		{
			name: "From 0, 0 to 0, 0, 0.1 - true",
			p:    Point{X: 0, Y: 0},
			p2:   Point{X: 0, Y: 0},
			d:    0.1,
			want: true,
		},
		{
			name: "From 0, 0 to 1, 0, 0.1 - false",
			p:    Point{X: 0, Y: 0},
			p2:   Point{X: 1, Y: 0},
			d:    0.1,
			want: false,
		},
		{
			name: "From 0, 0 to 1, 1, 1.415 - true",
			p:    Point{X: 0, Y: 0},
			p2:   Point{X: 1, Y: 1},
			d:    1.415,
			want: true,
		},
		{
			name: "From 0, 0 to 1, 1, 1.413 - false",
			p:    Point{X: 0, Y: 0},
			p2:   Point{X: 1, Y: 1},
			d:    1.413,
			want: false,
		},
		{
			name: "From 1, 1 to 0, 0, 1.415 - true",
			p:    Point{X: 1, Y: 1},
			p2:   Point{X: 0, Y: 0},
			d:    1.415,
			want: true,
		},
		{
			name: "From 1, 1 to 0, 0, 1.413 - false",
			p:    Point{X: 1, Y: 1},
			p2:   Point{X: 0, Y: 0},
			d:    1.413,
			want: false,
		},
		{
			name: "From 0, 0 to -1, 1, 1.415 - true",
			p:    Point{X: 0, Y: 0},
			p2:   Point{X: -1, Y: 1},
			d:    1.415,
			want: true,
		},
		{
			name: "From 0, 0 to 1, -1, 1.413 - false",
			p:    Point{X: 0, Y: 0},
			p2:   Point{X: 1, Y: -1},
			d:    1.413,
			want: false,
		},
		{
			name: "From 4, 3 to 10, 5, 6.32 + 0.1 - true",
			p:    Point{X: 4, Y: 3},
			p2:   Point{X: 10, Y: 5},
			d:    6.32 + 0.1,
			want: true,
		},
		{
			name: "From 4, 3 to 10, 5, 6.32 - 0.1 - false",
			p:    Point{X: 4, Y: 3},
			p2:   Point{X: 10, Y: 5},
			d:    6.32 - 0.1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Near(tt.p2, tt.d); got != tt.want {
				t.Errorf("Near() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_EqualApprox(t *testing.T) {
	tests := []struct {
		name string
		p    Point
		p2   Point
		d    float64
		want bool
	}{
		{
			name: "Exact same point - true",
			p:    Point{1, 1},
			p2:   Point{1, 1},
			d:    0.1,
			want: true,
		},
		{
			name: "false",
			p:    Point{0, 0},
			p2:   Point{0, 1},
			d:    0.5,
			want: false,
		},
		{
			name: "near - true",
			p:    Point{0, 0},
			p2:   Point{0.1, 0.1},
			d:    0.11,
			want: true,
		},
		{
			name: "near - true 2",
			p:    Point{0.1, 0.1},
			p2:   Point{0, 0},
			d:    0.11,
			want: true,
		},
		{
			name: "near - false",
			p:    Point{0.1, 0.1},
			p2:   Point{0, 0},
			d:    0.09,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.EqualApprox(tt.p2, tt.d); got != tt.want {
				t.Errorf("EqualApprox() = %v, want %v", got, tt.want)
			}
		})
	}
}
