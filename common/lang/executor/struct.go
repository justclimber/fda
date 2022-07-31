package executor

func NewStruct(name string, fields *Assignment) *Struct {
	return &Struct{
		key:    KeyStruct,
		name:   name,
		fields: fields,
	}
}

type Struct struct {
	id     int64
	key    NodeKey
	name   string
	fields *Assignment
}

func (s *Struct) ID() int64        { return s.id }
func (s *Struct) NodeKey() NodeKey { return s.key }

func (s *Struct) Exec(env *Environment, result *Result, executor execManager) error {
	definition, _ := executor.MainPackage().StructDefinition(s.name)
	fields := make(map[string]Object)
	newResult := NewResult()
	executor.AddNextExec(s.fields, func() error {
		return s.fields.Exec(env, newResult, executor)
	})

	for i := range definition.fields {
		ii := i
		executor.AddNextExec(s.fields.left[ii], func() error {
			fields[s.fields.left[ii].value] = newResult.objectList[ii]
			return nil
		})
	}

	executor.AddNextExec(s, func() error {
		result.Add(&ObjStruct{
			Definition: definition,
			Fields:     fields,
		})
		return nil
	})
	return nil
}
