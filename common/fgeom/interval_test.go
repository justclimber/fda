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
	"testing"
)

// Some standard intervals for use throughout the tests.
var (
	unit          = Interval{0, 1}
	negunit       = Interval{-1, 0}
	half          = Interval{0.5, 0.5}
	emptyInterval = EmptyInterval()
)

func TestIsEmpty(t *testing.T) {
	var zero Interval
	if unit.IsEmpty() {
		t.Errorf("%v should not be empty", unit)
	}
	if half.IsEmpty() {
		t.Errorf("%v should not be empty", half)
	}
	if !emptyInterval.IsEmpty() {
		t.Errorf("%v should be empty", emptyInterval)
	}
	if zero.IsEmpty() {
		t.Errorf("zero Interval %v should not be empty", zero)
	}
}

func TestIntervalCenter(t *testing.T) {
	tests := []struct {
		interval Interval
		want     float64
	}{
		{unit, 0.5},
		{negunit, -0.5},
		{half, 0.5},
	}
	for _, test := range tests {
		got := test.interval.Center()
		if got != test.want {
			t.Errorf("%v.Center() = %v, want %v", test.interval, got, test.want)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		interval Interval
		want     float64
	}{
		{unit, 1},
		{negunit, 1},
		{half, 0},
	}
	for _, test := range tests {
		if l := test.interval.Length(); l != test.want {
			t.Errorf("%v.Length() = %v, want %v", test.interval, l, test.want)
		}
	}
	if l := emptyInterval.Length(); l >= 0 {
		t.Errorf("empty interval has non-negative length")
	}
}

func TestIntervalContains(t *testing.T) {
	tests := []struct {
		interval         Interval
		p                float64
		contains         bool
		interiorContains bool
	}{
		{
			interval:         unit,
			p:                0.5,
			contains:         true,
			interiorContains: true,
		},
		{
			interval:         unit,
			p:                0,
			contains:         true,
			interiorContains: false,
		},
		{
			interval:         unit,
			p:                1,
			contains:         true,
			interiorContains: false,
		},
	}

	for _, test := range tests {
		if got := test.interval.Contains(test.p); got != test.contains {
			t.Errorf("%v.Contains(%v) = %t, want %t", test.interval, test.p, got, test.contains)
		}
		if got := test.interval.InteriorContains(test.p); got != test.interiorContains {
			t.Errorf("%v.InteriorContains(%v) = %t, want %t", test.interval, test.p, got, test.interiorContains)
		}
	}
}

func TestIntervalOperations(t *testing.T) {
	tests := []struct {
		have               Interval
		other              Interval
		contains           bool
		interiorContains   bool
		intersects         bool
		interiorIntersects bool
	}{
		{
			have:               emptyInterval,
			other:              emptyInterval,
			contains:           true,
			interiorContains:   true,
			intersects:         false,
			interiorIntersects: false,
		},
		{
			have:               emptyInterval,
			other:              unit,
			contains:           false,
			interiorContains:   false,
			intersects:         false,
			interiorIntersects: false,
		},
		{
			have:               unit,
			other:              half,
			contains:           true,
			interiorContains:   true,
			intersects:         true,
			interiorIntersects: true,
		},
		{
			have:               unit,
			other:              unit,
			contains:           true,
			interiorContains:   false,
			intersects:         true,
			interiorIntersects: true,
		},
		{
			have:               unit,
			other:              emptyInterval,
			contains:           true,
			interiorContains:   true,
			intersects:         false,
			interiorIntersects: false,
		},
		{
			have:               unit,
			other:              negunit,
			contains:           false,
			interiorContains:   false,
			intersects:         true,
			interiorIntersects: false,
		},
		{
			have:               unit,
			other:              Interval{0, 0.5},
			contains:           true,
			interiorContains:   false,
			intersects:         true,
			interiorIntersects: true,
		},
		{
			have:               half,
			other:              Interval{0, 0.5},
			contains:           false,
			interiorContains:   false,
			intersects:         true,
			interiorIntersects: false,
		},
	}

	for _, test := range tests {
		if got := test.have.ContainsInterval(test.other); got != test.contains {
			t.Errorf("%v.ContainsInterval(%v) = %t, want %t", test.have, test.other, got, test.contains)
		}
		if got := test.have.InteriorContainsInterval(test.other); got != test.interiorContains {
			t.Errorf("%v.InteriorContainsInterval(%v) = %t, want %t", test.have, test.other, got, test.interiorContains)
		}
		if got := test.have.Intersects(test.other); got != test.intersects {
			t.Errorf("%v.Intersects(%v) = %t, want %t", test.have, test.other, got, test.intersects)
		}
		if got := test.have.InteriorIntersects(test.other); got != test.interiorIntersects {
			t.Errorf("%v.InteriorIntersects(%v) = %t, want %t", test.have, test.other, got, test.interiorIntersects)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		x, y Interval
		want Interval
	}{
		{unit, half, half},
		{unit, negunit, Interval{0, 0}},
		{negunit, half, emptyInterval},
		{unit, emptyInterval, emptyInterval},
		{emptyInterval, unit, emptyInterval},
	}
	for _, test := range tests {
		if got := test.x.Intersection(test.y); !got.Equal(test.want) {
			t.Errorf("%v.Intersection(%v) = %v, want equal to %v", test.x, test.y, got, test.want)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		x, y Interval
		want Interval
	}{
		{Interval{99, 100}, emptyInterval, Interval{99, 100}},
		{emptyInterval, Interval{99, 100}, Interval{99, 100}},
		{Interval{5, 3}, Interval{0, -2}, emptyInterval},
		{Interval{0, -2}, Interval{5, 3}, emptyInterval},
		{unit, unit, unit},
		{unit, negunit, Interval{-1, 1}},
		{negunit, unit, Interval{-1, 1}},
		{half, unit, unit},
	}
	for _, test := range tests {
		if got := test.x.Union(test.y); !got.Equal(test.want) {
			t.Errorf("%v.Union(%v) = %v, want equal to %v", test.x, test.y, got, test.want)
		}
	}
}

func TestIntervalAddPoint(t *testing.T) {
	tests := []struct {
		interval Interval
		point    float64
		want     Interval
	}{
		{emptyInterval, 5, Interval{5, 5}},
		{Interval{5, 5}, -1, Interval{-1, 5}},
		{Interval{-1, 5}, 0, Interval{-1, 5}},
		{Interval{-1, 5}, 6, Interval{-1, 6}},
	}
	for _, test := range tests {
		if got := test.interval.AddPoint(test.point); !got.Equal(test.want) {
			t.Errorf("%v.AddPoint(%v) = %v, want equal to %v", test.interval, test.point, got, test.want)
		}
	}
}

func TestIntervalClampPoint(t *testing.T) {
	tests := []struct {
		interval Interval
		clamp    float64
		want     float64
	}{
		{Interval{0.1, 0.4}, 0.3, 0.3},
		{Interval{0.1, 0.4}, -7.0, 0.1},
		{Interval{0.1, 0.4}, 0.6, 0.4},
	}
	for _, test := range tests {
		if got := test.interval.ClampPoint(test.clamp); got != test.want {
			t.Errorf("%v.ClampPoint(%v) = %v, want equal to %v", test.interval, test.clamp, got, test.want)
		}
	}
}

func TestExpanded(t *testing.T) {
	tests := []struct {
		interval Interval
		margin   float64
		want     Interval
	}{
		{emptyInterval, 0.45, emptyInterval},
		{unit, 0.5, Interval{-0.5, 1.5}},
		{unit, -0.5, Interval{0.5, 0.5}},
		{unit, -0.51, emptyInterval},
	}
	for _, test := range tests {
		if got := test.interval.Expanded(test.margin); !got.Equal(test.want) {
			t.Errorf("%v.Expanded(%v) = %v, want equal to %v", test.interval, test.margin, got, test.want)
		}
	}
}

func TestIntervalString(t *testing.T) {
	i := Interval{2, 4.5}
	if s, exp := i.String(), "[2.0000000, 4.5000000]"; s != exp {
		t.Errorf("i.String() = %q, want %q", s, exp)
	}
}

func TestApproxEqual(t *testing.T) {
	// Choose two values lo and hi such that it's okay to shift an endpoint by
	// kLo (i.e., the resulting interval is equivalent) but not by kHi.
	const lo = 4 * dblEpsilon // < max_error default
	const hi = 6 * dblEpsilon // > max_error default

	tests := []struct {
		interval Interval
		other    Interval
		want     bool
	}{
		// Empty intervals.
		{EmptyInterval(), EmptyInterval(), true},
		{Interval{0, 0}, EmptyInterval(), true},
		{EmptyInterval(), Interval{0, 0}, true},
		{Interval{1, 1}, EmptyInterval(), true},
		{EmptyInterval(), Interval{1, 1}, true},
		{EmptyInterval(), Interval{0, 1}, false},
		{EmptyInterval(), Interval{1, 1 + 2*lo}, true},
		{EmptyInterval(), Interval{1, 1 + 2*hi}, false},

		// Singleton intervals.
		{Interval{1, 1}, Interval{1, 1}, true},
		{Interval{1, 1}, Interval{1 - lo, 1 - lo}, true},
		{Interval{1, 1}, Interval{1 + lo, 1 + lo}, true},
		{Interval{1, 1}, Interval{1 - hi, 1}, false},
		{Interval{1, 1}, Interval{1, 1 + hi}, false},
		{Interval{1, 1}, Interval{1 - lo, 1 + lo}, true},
		{Interval{0, 0}, Interval{1, 1}, false},

		// Other intervals.
		{Interval{1 - lo, 2 + lo}, Interval{1, 2}, true},
		{Interval{1 + lo, 2 - lo}, Interval{1, 2}, true},

		{Interval{1 - hi, 2 + lo}, Interval{1, 2}, false},
		{Interval{1 + hi, 2 - lo}, Interval{1, 2}, false},
		{Interval{1 - lo, 2 + hi}, Interval{1, 2}, false},
		{Interval{1 + lo, 2 - hi}, Interval{1, 2}, false},
	}

	for d, test := range tests {
		if got := test.interval.ApproxEqual(test.other); got != test.want {
			t.Errorf("%d. %v.ApproxEqual(%v) = %t, want %t", d,
				test.interval, test.other, got, test.want)
		}
	}
}
