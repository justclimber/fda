package result

import (
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewResult() *Result {
	return &Result{
		ResultTypeList: make([]object.Type, 0),
	}
}

func NewSingleResult(objType object.Type) *Result {
	return &Result{
		ResultTypeList: []object.Type{objType},
	}
}

type Result struct {
	ResultTypeList []object.Type
}

func (r *Result) Add(objType object.Type) {
	r.ResultTypeList = append(r.ResultTypeList, objType)
}

func (r *Result) Merge(r2 *Result) {
	for _, rt := range r2.ResultTypeList {
		r.Add(rt)
	}
}

func (r *Result) Get() object.Type {
	return r.ResultTypeList[0]
}

func (r *Result) Count() int {
	return len(r.ResultTypeList)
}

func (r *Result) GetByIndex(index int) object.Type {
	return r.ResultTypeList[index]
}

func NewNamedResult() *NamedResult {
	return &NamedResult{
		ResultTypeList: make(map[string]object.Type),
	}
}

type NamedResult struct {
	ResultTypeList map[string]object.Type
}

func (nm *NamedResult) Get(name string) object.Type {
	return nm.ResultTypeList[name]
}

func (nm *NamedResult) Set(name string, objType object.Type) {
	nm.ResultTypeList[name] = objType
}
