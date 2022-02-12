package ecs

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

//go:generate go run ../../../common/generators/entityrepo/main.go wprepo

//generate:entities
var _ = [][]component.Key{
	{wpcomponent.KeyPosition, wpcomponent.KeyMoving},
	{wpcomponent.KeyPosition, wpcomponent.KeyMoving, wpcomponent.KeyPlayer},
}

//generate:systems
var _ = [][]component.Key{
	{wpcomponent.KeyPosition, wpcomponent.KeyMoving},
	{wpcomponent.KeyMoving, wpcomponent.KeyPlayer},
}
