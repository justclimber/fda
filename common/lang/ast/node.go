package ast

type NodeKey int32

const (
	KeyIllegal NodeKey = iota
	KeyFunctionDefinition
	KeyFunction
	KeyFunctionCall
	KeyPackage
	KeyStatementsBlock
	KeyVoidedExpression
	KeyExpressionList
	KeyNamedExpressionList
	KeyIfStatement
	KeyAssignment
	KeyIdentifier
	KeyArithmeticOperation
	KeyComparisonOperation
	KeyUnaryMinus
	KeyVarAndType
	KeyStructDefinition
	KeyStruct
	KeyStructFieldIdentifier
	KeyStructFieldCall
	KeyStructFieldAssignment
	KeyNumInt
	KeyNumFloat
	KeyBool
)

type Node interface {
	ID() int64
	NodeKey() NodeKey
}
