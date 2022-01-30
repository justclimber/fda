package ecs

type ComponentKey string
type EntityId int64

type Entity struct {
	Id         EntityId
	Components map[ComponentKey]interface{}
}
