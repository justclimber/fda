package ide

import (
	"github.com/justclimber/fda/client/ide/ast"
)

type Tab struct {
	name       string
	packageAst *ast.Package
}

func NewTab(name string, packageAst *ast.Package) *Tab {
	return &Tab{name: name, packageAst: packageAst}
}

func (t *Tab) Draw(r ast.Renderer) {
	r.DrawTabBody()
	t.packageAst.Draw(r, "")
}
