package ast

func NewStructFieldIdentifier(fieldName string, structIdentifier *Identifier) *StructFieldIdentifier {
	return &StructFieldIdentifier{
		key:              KeyStructFieldIdentifier,
		fieldName:        fieldName,
		structIdentifier: structIdentifier,
	}
}

type StructFieldIdentifier struct {
	id               int64
	key              NodeKey
	fieldName        string
	structIdentifier *Identifier
}

func (sf *StructFieldIdentifier) ID() int64        { return sf.id }
func (sf *StructFieldIdentifier) NodeKey() NodeKey { return sf.key }
