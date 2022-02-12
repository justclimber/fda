// Code generated by entityrepo generator. DO NOT EDIT.
package [[ .PackageName ]]

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
[[- range $key, $value := .KeysPackages]]
	"github.com/justclimber/fda/server/worldprocessor/ecs/[[ $key ]]"
[[- end -]]
)

type ECGroup[[ .MaskName ]] struct {
	Chunks []*Chunk[[ .MaskName ]]
	last   *Chunk[[ .MaskName ]]
}

func NewECGroup[[ .MaskName ]]() *ECGroup[[ .MaskName ]] {
	chunks := []*Chunk[[ .MaskName ]]{{}}
	return &ECGroup[[ .MaskName ]]{
		Chunks: chunks,
		last:   chunks[0],
	}
}

func (eg *ECGroup[[ .MaskName ]]) Add(e entity.Entity) (int, int) {
	if eg.last.Size == chunkSize[[ .MaskName ]] {
		eg.last = &Chunk[[ .MaskName ]]{}
		eg.Chunks = append(eg.Chunks, eg.last)
	}
	eg.last.Add(e)
	return len(eg.Chunks) - 1, eg.last.Size - 1
}

func (eg *ECGroup[[ .MaskName ]]) Get(addr entityrepo.EAddress) (entity.Entity, bool) {
	if addr.ChunkIndex >= len(eg.Chunks) {
		return entity.Entity{}, false
	}

	chunk := eg.Chunks[addr.ChunkIndex]

	if addr.Index >= chunk.Size {
		return entity.Entity{}, false
	}

	return entity.Entity{
		Id: chunk.Ids[addr.Index],
		Components: map[component.Key]component.Component{
[[- range .Keys]]
			[[ .FullStr ]]: &chunk.[[ .StrWithoutPrefix]][addr.Index],
[[- end]]
		},
		CMask: component.NewMask([]component.Key{
[[- range .Keys]]
			[[ .FullStr ]],
[[- end]]
		}),
	}, true
}