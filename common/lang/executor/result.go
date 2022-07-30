package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

var (
	ReservedObjTrue  = &fdalang.ObjBoolean{Value: true}
	ReservedObjFalse = &fdalang.ObjBoolean{Value: false}
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

func toReservedBoolObj(value bool) *fdalang.ObjBoolean {
	if value == true {
		return ReservedObjTrue
	}
	return ReservedObjFalse
}
