package component

type Key string

type Component interface {
	Key() Key
}
