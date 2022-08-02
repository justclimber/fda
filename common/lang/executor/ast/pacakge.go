package ast

func NewPackage() *Package {
	return &Package{
		key:                 KeyPackage,
		structDefinitions:   make(map[string]*StructDefinition),
		functionDefinitions: make(map[string]*FunctionDefinition),
	}
}

type Package struct {
	id                  int64
	key                 NodeKey
	structDefinitions   map[string]*StructDefinition
	functionDefinitions map[string]*FunctionDefinition
}

func (p *Package) ID() int64        { return p.id }
func (p *Package) NodeKey() NodeKey { return p.key }

func (p *Package) RegisterStructDefinition(s *StructDefinition) {
	p.structDefinitions[s.Name] = s
}

func (p *Package) StructDefinition(name string) (*StructDefinition, bool) {
	s, ok := p.structDefinitions[name]
	return s, ok
}

func (p *Package) RegisterFunctionDefinition(f *FunctionDefinition) {
	p.functionDefinitions[f.Name] = f
}

func (p *Package) FunctionDefinition(name string) (*FunctionDefinition, bool) {
	f, ok := p.functionDefinitions[name]
	return f, ok
}
