package ecs

type ComponentKey string
type EntityId int64

type Entity struct {
	Id         EntityId
	Components map[ComponentKey]interface{}
}

func (e *Entity) AddComponent(key ComponentKey, component interface{}) {
	e.Components[key] = component
}
