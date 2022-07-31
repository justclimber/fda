package executor

var (
	ReservedObjTrue  = &ObjBoolean{Value: true}
	ReservedObjFalse = &ObjBoolean{Value: false}
)

func NewResult() *Result {
	return &Result{
		objectList: make([]Object, 0),
	}
}

type Result struct {
	objectList []Object
}

func (r *Result) Add(object Object) {
	r.objectList = append(r.objectList, object)
}

func toReservedBoolObj(value bool) *ObjBoolean {
	if value == true {
		return ReservedObjTrue
	}
	return ReservedObjFalse
}
