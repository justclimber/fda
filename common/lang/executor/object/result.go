package object

var (
	ReservedObjTrue  = &ObjBoolean{Value: true}
	ReservedObjFalse = &ObjBoolean{Value: false}
)

func NewResult() *Result {
	return &Result{
		ObjectList: make([]Object, 0),
	}
}

type Result struct {
	ObjectList []Object
}

func (r *Result) Add(object Object) {
	r.ObjectList = append(r.ObjectList, object)
}

func ToReservedBoolObj(value bool) *ObjBoolean {
	if value == true {
		return ReservedObjTrue
	}
	return ReservedObjFalse
}
