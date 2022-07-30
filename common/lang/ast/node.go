package ast

type Key int32

const (
	KeyIllegal Key = iota
	KeyAssignment
	KeyIdentifier
)

type Node interface {
	Key() Key
}

type Stmt interface{ StmtMixin() }
type Expr interface{ ExprMixin() }

type Identifier struct{}

type Assignment struct {
	Left  Identifier
	Value Expr
}

func (a Assignment) Key() Key { return KeyAssignment }
func (i Identifier) Key() Key { return KeyIdentifier }

func (a Assignment) StmtMixin() {}

func (i Identifier) ExprMixin() {}
