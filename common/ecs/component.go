package ecs

type ComponentKey string

type Component interface {
	Key() ComponentKey
}
