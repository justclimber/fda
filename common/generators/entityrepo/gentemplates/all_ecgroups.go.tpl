// Code generated by entityrepo generator. DO NOT EDIT.
package [[ .PackageName ]]

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entityrepo"
)

func GetAllECGroups() map[component.Mask]entityrepo.CGroup {
	return map[component.Mask]entityrepo.CGroup{
	    [[- range .ECGroups ]]
		[[ .MaskName ]]: NewECGroup[[ .MaskName ]](),
		[[- end ]]
	}
}