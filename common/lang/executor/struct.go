package executor

func NewStruct(fields *Assignment) *Struct {
	return &Struct{
		key:    KeyStruct,
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
	err := s.fields.Exec(env, newResult, executor)
	if err != nil {
		return err
	}

	for i := range newResult.objectList {
		ii := i
		fields[s.fields.left[ii].value] = newResult.objectList[ii]
	}
	result.Add(&ObjStruct{
		Definition: definition,
		Fields:     fields,
	})
	return nil
}
