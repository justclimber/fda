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
	"image"
	"math"
)

var EmptyPoint Point

type Point struct {
	X, Y float64
}

// Add returns the sum of p and op.
func (p Point) Add(op Point) Point { return Point{p.X + op.X, p.Y + op.Y} }

// Sub returns the difference of p and op.
func (p Point) Sub(op Point) Point { return Point{p.X - op.X, p.Y - op.Y} }

// Mul returns the scalar product of p and m.
func (p Point) Mul(m float64) Point { return Point{m * p.X, m * p.Y} }

// Ortho returns a counterclockwise orthogonal point with the same norm.
func (p Point) Ortho() Point { return Point{-p.Y, p.X} }

// Dot returns the dot product between p and op.
func (p Point) Dot(op Point) float64 { return p.X*op.X + p.Y*op.Y }

// Cross returns the cross product of p and op.
func (p Point) Cross(op Point) float64 { return p.X*op.Y - p.Y*op.X }

// Norm returns the vector's norm.
func (p Point) Norm() float64 { return math.Hypot(p.X, p.Y) }

// Normalize returns a unit point in the same direction as p.
func (p Point) Normalize() Point {
	if p.X == 0 && p.Y == 0 {
		return p
	}
	return p.Mul(1 / p.Norm())
}

func (p Point) String() string { return fmt.Sprintf("(%.12f, %.12f)", p.X, p.Y) }

// ToImagePoint returns image.Point representation of Point
func (p Point) ToImagePoint() image.Point {
	return image.Point{
		X: int(p.X),
		Y: int(p.Y),
	}
}

// Near returns true if the distance between p2 and p less oq equal than d (without inaccuracy)
func (p Point) Near(p2 Point, d float64) bool {
	pd := (p2.X-p.X)*(p2.X-p.X) + (p2.Y-p.Y)*(p2.Y-p.Y)
	return pd <= d*d
}

// EqualApprox returns true if p2 equals p with d difference
func (p Point) EqualApprox(p2 Point, d float64) bool {
	return math.Abs(p2.Y-p.Y) < d && math.Abs(p2.X-p.X) < d
}

func (p Point) Empty() bool {
	return p == EmptyPoint
}
