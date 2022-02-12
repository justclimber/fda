package component

import (
	"fmt"
)

type Key int64

type Mask struct {
	fmt.Stringer
	data int64
}

func NewMask(keys []Key) Mask {
	m := Mask{}
	for _, k := range keys {
		m.Add(k)
	}
	return m
}

func (m *Mask) Add(key Key) {
	m.data = m.data | int64(key)
}

func (m *Mask) String() string {
	return fmt.Sprintf("%d [%[1]b]", m.data)
}

func (m *Mask) Intersect(m2 Mask) bool {
	return m.data&m2.data != 0
}

func (m *Mask) Contains(m2 Mask) bool {
	return m.data&m2.data == m2.data
}

type Component interface {
	Key() Key
}
