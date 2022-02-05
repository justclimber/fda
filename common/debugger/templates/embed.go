package templates

import (
	"embed"
)

//go:embed **
var EmbeddedFS embed.FS
