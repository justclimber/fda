package assets

import (
	"embed"
)

//go:embed **
var EmbeddedFS embed.FS
