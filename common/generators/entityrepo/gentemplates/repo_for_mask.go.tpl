package [[ .PackageName ]]

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
[[- range $key, $value := .KeysPackages]]
	"github.com/justclimber/fda/server/worldprocessor/ecs/[[ $key ]]"
[[- end -]]
)

type RepoFor[[ .MaskName ]] struct {
	cGroups  []entityrepo.CGroup
	repoLink ecs.EntityRepo
}

func NewRepoFor[[ .MaskName ]](repoLink ecs.EntityRepo) *RepoFor[[ .MaskName ]] {
	return &RepoFor[[ .MaskName ]]{
		repoLink: repoLink,
	}
}

func (ci *RepoFor[[ .MaskName ]]) InitRepoLink(mask component.Mask) {
	ci.cGroups = ci.repoLink.GetCGroupsWithMask(mask)
}

func (ci *RepoFor[[ .MaskName ]]) Iterate(f func(
	id entity.Id,
	[[- range $.Keys ]]
	[[ .StrWithoutPrefix ]] [[ .PackageName ]].[[ .StrWithoutPrefix ]],
	[[- end ]]
) ([[ joinKeys ", " "*wpcomponent." "" $.Keys ]])) {
	for _, cGroup := range ci.cGroups {
		switch cg := cGroup.(type) {
		[[ range .ECGroups -]]
		case *[[ . ]]:
			for _, chunk := range cg.Chunks {
				for k := 0; k < chunk.Size; k++ {
                    [[ joinKeys ", " "new" "" $.Keys ]] := f(chunk.Ids[k], [[ joinKeys ", " "chunk." "[k]" $.Keys -]])
                    [[ range $.Keys ]]
					if new[[ .StrWithoutPrefix ]] != nil {
						chunk.[[ .StrWithoutPrefix ]][k] = *new[[ .StrWithoutPrefix ]]
					}
					[[- end ]]
				}
			}
		[[ end -]]
		}
	}
}
