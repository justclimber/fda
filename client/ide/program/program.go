package program

type Program struct {
	packagist *Packagist
	// version
}

func NewProgram(packagist *Packagist) *Program {
	return &Program{packagist: packagist}
}
