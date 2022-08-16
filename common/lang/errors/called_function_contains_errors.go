package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrCalledFunctionContainsErrors struct {
	node         ast.Node
	functionName string
	packageName  string
}

func NewErrCalledFunctionContainsErrors(node ast.Node, functionName string, packageName string) ErrCalledFunctionContainsErrors {
	return ErrCalledFunctionContainsErrors{node: node, functionName: functionName, packageName: packageName}
}

func (e ErrCalledFunctionContainsErrors) Error() string {
	return fmt.Sprintf("called function %s#%s contains errors", e.packageName, e.functionName)
}
