package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

type NodeKey int32

const (
	KeyIllegal NodeKey = iota
	KeyFunction
	KeyPackage
	KeyStatementsBlock
	KeyVoidedExpression
	KeyExpressionList
	KeyIfStatement
	KeyAssignment
	KeyIdentifier
	KeyUnaryMinus
	KeyNumInt
	KeyBool
)

type Node interface {
	ID() int64
	NodeKey() NodeKey
}

type Stmt interface {
	Node
	Exec(env *fdalang.Environment, execQueue *ExecFnList) error
}

type Expr interface {
	Node
	Exec(env *fdalang.Environment, result *Result, execQueue *ExecFnList) error
}
