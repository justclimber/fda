// Code generated by entityrepo generator. DO NOT EDIT.
package wprepo

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

var (
    Mask3 = component.NewMask([]component.Key{
        wpcomponent.KeyPosition,
        wpcomponent.KeyMoving,
    })
    Mask6 = component.NewMask([]component.Key{
        wpcomponent.KeyMoving,
        wpcomponent.KeyPlayer,
    })
    Mask7 = component.NewMask([]component.Key{
        wpcomponent.KeyPosition,
        wpcomponent.KeyMoving,
        wpcomponent.KeyPlayer,
    })
)
