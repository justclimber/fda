package [[ .PackageName ]]

import (
	"github.com/justclimber/fda/common/ecs/entity"
[[- range $key, $value := .KeysPackages]]
	"github.com/justclimber/fda/server/worldprocessor/ecs/[[ $key ]]"
[[- end]]
)

const chunkSize[[ .MaskName ]] = 10

type Chunk[[ .MaskName ]] struct {
	Size     int
	Ids      [chunkSize[[ $.MaskName -]] ]entity.Id
[[- range .Keys]]
	[[ .StrWithoutPrefix]]   [chunkSize[[ $.MaskName -]] ][[ .PackageName ]].[[ .StrWithoutPrefix ]]
[[- end]]
}

func (ch *Chunk[[ .MaskName ]]) Add(e entity.Entity) {
	ch.Ids[ch.Size] = e.Id
[[- range .Keys]]
	ch.[[ .StrWithoutPrefix ]][ch.Size] = *e.Components[ [[- .FullStr -]] ].(*[[ .PackageName ]].[[ .StrWithoutPrefix ]])
[[- end]]
	ch.Size++
}