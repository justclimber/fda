package gentemplates

import (
	"embed"
)

//go:embed **
var EmbeddedFS embed.FS
