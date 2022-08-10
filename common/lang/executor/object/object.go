package object

import (
	"bytes"
	"fmt"
	"strings"
)

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	TypeInt   Type = "std#int"
	TypeFloat Type = "std#float"
	TypeBool  Type = "std#bool"
)

type Object interface {
	GetType() Type
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

func (i *ObjInteger) GetType() Type   { return TypeInt }
func (i *ObjInteger) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type ObjFloat struct {
	Emptier
	Value float64
}

func (f *ObjFloat) GetType() Type   { return TypeFloat }
func (f *ObjFloat) Inspect() string { return fmt.Sprintf("%.2f", f.Value) }

type ObjBoolean struct {
	Value bool
}

func (b *ObjBoolean) GetType() Type   { return TypeBool }
func (b *ObjBoolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type ObjArray struct {
	Emptier
	ElementsType Type
	Elements     []Object
}

func (a *ObjArray) GetType() Type {
	return Type(fmt.Sprintf("[]%s", a.ElementsType))
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

func (f *ObjFunction) GetType() Type { return f.Definition.Type() }
func (f *ObjFunction) Inspect() string {
	return fmt.Sprintf("fn %s", f.Definition.Type())
}

type ObjStruct struct {
	Emptier
	Definition *StructDefinition
	Fields     map[string]Object
}

func (s *ObjStruct) GetType() Type { return s.Definition.Type() }
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
