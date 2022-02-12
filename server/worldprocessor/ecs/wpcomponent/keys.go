package wpcomponent

import (
	"github.com/justclimber/fda/common/ecs/component"
)

const (
	KeyPosition component.Key = 1 << iota
	KeyMoving
	KeyPlayer
	KeyBody
)
