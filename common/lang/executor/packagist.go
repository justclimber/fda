package executor

import (
	"github.com/justclimber/fda/common/lang/executor/ast"
)

type Packagist struct {
	mainPackage *ast.Package
	packages    map[string]*ast.Package
}

func NewPackagist(mainPackage *ast.Package) *Packagist {
	return &Packagist{
		mainPackage: mainPackage,
		packages:    make(map[string]*ast.Package),
	}
}

func (p *Packagist) Main() *ast.Package {
	return p.mainPackage
}
