package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
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
	KeyVarAndType
	KeyStructDefinition
	KeyStruct
	KeyNumInt
	KeyBool
)

type Node interface {
	ID() int64
	NodeKey() NodeKey
}

type execManager interface {
	AddNextExec(node Node, fn func() error)
	MainPackage() *Package
}

type Stmt interface {
	Node
	Exec(env *environment.Environment, execMngr execManager) error
}

type Expr interface {
	Node
	Exec(env *environment.Environment, result *object.Result, execMngr execManager) error
}
