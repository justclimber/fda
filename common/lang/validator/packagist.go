package validator

type Packagist struct {
	mainPackage *Package
	packages    map[string]*Package
}

func NewPackagist(mainPackage *Package) *Packagist {
	return &Packagist{
		mainPackage: mainPackage,
		packages:    make(map[string]*Package),
	}
}

func (p *Packagist) Main() *Package {
	return p.mainPackage
}
