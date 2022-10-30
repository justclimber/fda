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
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Interval represents a closed interval
// Zero-length intervals (where Lo == Hi) represent single points.
// If Lo > Hi then the interval is empty.
type Interval[T Number] struct {
	Lo, Hi T
}

// EmptyInterval returns an empty interval.
func EmptyInterval[T Number]() Interval[T] { return Interval[T]{1, 0} }

// IntervalFromPoint returns an interval representing a single point.
func IntervalFromPoint[T Number](p T) Interval[T] { return Interval[T]{p, p} }

// IsEmpty reports whether the interval is empty.
func (i Interval[T]) IsEmpty() bool { return i.Lo > i.Hi }

// Equal returns true iff the interval contains the same points as oi.
func (i Interval[T]) Equal(oi Interval[T]) bool {
	return i == oi || i.IsEmpty() && oi.IsEmpty()
}

// Center returns the midpoint of the interval.
// It is undefined for empty intervals.
func (i Interval[T]) Center() T { return (i.Lo + i.Hi) / 2 }

// Length returns the length of the interval.
// The length of an empty interval is negative.
func (i Interval[T]) Length() T { return i.Hi - i.Lo }

// Contains returns true iff the interval contains p.
func (i Interval[T]) Contains(p T) bool { return i.Lo <= p && p <= i.Hi }

// ContainsInterval returns true iff the interval contains oi.
func (i Interval[T]) ContainsInterval(oi Interval[T]) bool {
	if oi.IsEmpty() {
		return true
	}
	return i.Lo <= oi.Lo && oi.Hi <= i.Hi
}

// InteriorContains returns true iff the interval strictly contains p.
func (i Interval[T]) InteriorContains(p T) bool {
	return i.Lo < p && p < i.Hi
}

// InteriorContainsInterval returns true iff the interval strictly contains oi.
func (i Interval[T]) InteriorContainsInterval(oi Interval[T]) bool {
	if oi.IsEmpty() {
		return true
	}
	return i.Lo < oi.Lo && oi.Hi < i.Hi
}

// Intersects returns true iff the interval contains any points in common with oi.
func (i Interval[T]) Intersects(oi Interval[T]) bool {
	if i.Lo <= oi.Lo {
		return oi.Lo <= i.Hi && oi.Lo <= oi.Hi // oi.Lo ∈ i and oi is not empty
	}
	return i.Lo <= oi.Hi && i.Lo <= i.Hi // i.Lo ∈ oi and i is not empty
}

// InteriorIntersects returns true iff the interior of the interval contains any points in common with oi, including the latter's boundary.
func (i Interval[T]) InteriorIntersects(oi Interval[T]) bool {
	return oi.Lo < i.Hi && i.Lo < oi.Hi && i.Lo < i.Hi && oi.Lo <= oi.Hi
}

// Intersection returns the interval containing all points common to i and j.
func (i Interval[T]) Intersection(j Interval[T]) Interval[T] {
	// Empty intervals do not need to be special-cased.
	return Interval[T]{
		Lo: Max(i.Lo, j.Lo),
		Hi: Min(i.Hi, j.Hi),
	}
}

func Max[T Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// AddPoint returns the interval expanded so that it contains the given point.
func (i Interval[T]) AddPoint(p T) Interval[T] {
	if i.IsEmpty() {
		return Interval[T]{p, p}
	}
	if p < i.Lo {
		return Interval[T]{p, i.Hi}
	}
	if p > i.Hi {
		return Interval[T]{i.Lo, p}
	}
	return i
}

// ClampPoint returns the closest point in the interval to the given point "p".
// The interval must be non-empty.
func (i Interval[T]) ClampPoint(p T) T {
	return Max(i.Lo, Min(i.Hi, p))
}

// Expanded returns an interval that has been expanded on each side by margin.
// If margin is negative, then the function shrinks the interval on
// each side by margin instead. The resulting interval may be empty. Any
// expansion of an empty interval remains empty.
func (i Interval[T]) Expanded(margin T) Interval[T] {
	if i.IsEmpty() {
		return i
	}
	return Interval[T]{i.Lo - margin, i.Hi + margin}
}

// Union returns the smallest interval that contains this interval and the given interval.
func (i Interval[T]) Union(other Interval[T]) Interval[T] {
	if i.IsEmpty() {
		return other
	}
	if other.IsEmpty() {
		return i
	}
	return Interval[T]{Min(i.Lo, other.Lo), Max(i.Hi, other.Hi)}
}

func (i Interval[T]) String() string { return fmt.Sprintf("[%v, %v]", i.Lo, i.Hi) }

const (
	// epsilon is a small number that represents a reasonable level of noise between two
	// values that can be considered to be equal.
	epsilon = 1e-15
	// dblEpsilon is a smaller number for values that require more precision.
	// This is the C++ DBL_EPSILON equivalent.
	dblEpsilon = 2.220446049250313e-16
)

// ApproxEqual reports whether the interval can be transformed into the
// given interval by moving each endpoint a small distance.
// The empty interval is considered to be positioned arbitrarily on the
// real line, so any interval with a small enough length will match
// the empty interval.
func (i Interval[T]) ApproxEqual(other Interval[float64]) bool {
	if i.IsEmpty() {
		return other.Length() <= 2*epsilon
	}
	if other.IsEmpty() {
		return float64(i.Length()) <= 2*epsilon
	}
	return Abs[float64](other.Lo-float64(i.Lo)) <= epsilon &&
		Abs[float64](other.Hi-float64(i.Hi)) <= epsilon
}

// DirectedHausdorffDistance returns the Hausdorff distance to the given interval. For two
// intervals x and y, this distance is defined as
//     h(x, y) = max_{p in x} min_{q in y} d(p, q).
func (i Interval[T]) DirectedHausdorffDistance(other Interval[T]) T {
	if i.IsEmpty() {
		return 0
	}
	if other.IsEmpty() {
		return T(math.Inf(1))
	}
	return Max(0, Max(i.Hi-other.Hi, other.Lo-i.Lo))
}

func Abs[T Number](x T) T {
	return T(math.Float64frombits(math.Float64bits(float64(x)) &^ (1 << 63)))
}
