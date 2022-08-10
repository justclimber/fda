package object

import (
	"bytes"
	"fmt"
	"strings"
)

type ObjectType string

func (o ObjectType) String() string {
	return string(o)
}

const (
	TypeInt   ObjectType = "std#int"
	TypeFloat ObjectType = "std#float"
	TypeBool  ObjectType = "std#bool"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Emptier struct {
	Empty bool
}

func (e *Emptier) IsEmpty() bool { return e.Empty }

type ObjInteger struct {
	Emptier
	Value int64
}

func (i *ObjInteger) Type() ObjectType { return TypeInt }
func (i *ObjInteger) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type ObjFloat struct {
	Emptier
	Value float64
}

func (f *ObjFloat) Type() ObjectType { return TypeFloat }
func (f *ObjFloat) Inspect() string  { return fmt.Sprintf("%.2f", f.Value) }

type ObjBoolean struct {
	Value bool
}

func (b *ObjBoolean) Type() ObjectType { return TypeBool }
func (b *ObjBoolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type ObjArray struct {
	Emptier
	ElementsType ObjectType
	Elements     []Object
}

func (a *ObjArray) Type() ObjectType {
	varType := fmt.Sprintf("[]%s", a.ElementsType)
	return ObjectType(varType)
}
func (a *ObjArray) Inspect() string {
	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	return fmt.Sprintf("[]%s{%s}", a.ElementsType, strings.Join(elements, ", "))
}

type ObjFunction struct {
	Definition *FunctionDefinition
}

func (f *ObjFunction) Type() ObjectType { return f.Definition.Type() }
func (f *ObjFunction) Inspect() string {
	return fmt.Sprintf("fn %s", f.Definition.Type())
}

type ObjStruct struct {
	Emptier
	Definition *StructDefinition
	Fields     map[string]Object
}

func (s *ObjStruct) Type() ObjectType { return s.Definition.Type() }
func (s *ObjStruct) Inspect() string {
	var out bytes.Buffer

	var elements []string
	for k, v := range s.Fields {
		elements = append(elements, fmt.Sprintf("%s: %s", k, v.Inspect()))
	}

	out.WriteString(s.Definition.Type().String())
	out.WriteString("{")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("}")

	return out.String()
}
