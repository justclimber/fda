package program

import (
	"github.com/justclimber/fda/common/lang/validator/ast"
)

type Packagist struct {
	packages map[string]*ast.Package
}

func NewPackagist() *Packagist {
	return &Packagist{
		packages: make(map[string]*ast.Package),
	}
}

func (p *Packagist) PackageByName(name string) (*ast.Package, bool) {
	pkg, ok := p.packages[name]
	return pkg, ok
}

func (p *Packagist) RegisterPackage(pkg *ast.Package) error {
	p.packages[pkg.Name] = pkg
	return nil
}
