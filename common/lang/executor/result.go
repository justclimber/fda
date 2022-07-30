package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

func NewResult() *Result {
	return &Result{
		objectList: make([]fdalang.Object, 0),
	}
}

type Result struct {
	objectList []fdalang.Object
}

func (r *Result) Add(object fdalang.Object) {
	r.objectList = append(r.objectList, object)
}
