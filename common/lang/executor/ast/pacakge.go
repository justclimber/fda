package ast

func NewPackage(mainFunction *Function) *Package {
	return &Package{
		key:               KeyPackage,
		mainFunction:      mainFunction,
		structDefinitions: make(map[string]*StructDefinition),
	}
}

type Package struct {
	id                int64
	key               NodeKey
	mainFunction      *Function
	structDefinitions map[string]*StructDefinition
}

func (p *Package) ID() int64        { return p.id }
func (p *Package) NodeKey() NodeKey { return p.key }

func (p *Package) RegisterStructDefinition(s *StructDefinition) {
	p.structDefinitions[s.name] = s
}

func (p *Package) StructDefinition(name string) (*StructDefinition, bool) {
	s, ok := p.structDefinitions[name]
	return s, ok
}
