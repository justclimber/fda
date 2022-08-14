package errors

import (
	"strings"

	"github.com/justclimber/fda/common/lang/ast"
)

func NewErrContainer(node ast.Node) *ErrContainer {
	return &ErrContainer{
		node:    node,
		errList: make([]error, 0),
	}
}

type ErrContainer struct {
	node    ast.Node
	errList []error
}

func (vs *ErrContainer) Add(e error) {
	if e != nil {
		vs.errList = append(vs.errList, e)
	}
}

func (vs *ErrContainer) NotEmpty() bool {
	return len(vs.errList) > 0
}

func (vs *ErrContainer) Error() string {
	var s []string
	for _, err := range vs.errList {
		s = append(s, err.Error())
	}
	return strings.Join(s, "\n")
}
