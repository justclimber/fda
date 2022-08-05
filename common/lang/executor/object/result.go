package object

var (
	ReservedObjTrue  = &ObjBoolean{Value: true}
	ReservedObjFalse = &ObjBoolean{Value: false}
)

type ArithmeticOperator int8
type ComparisonOperator int8

const (
	OperatorAddition ArithmeticOperator = iota
	OperatorSubtraction
	OperatorMultiplication
	OperatorDivision
)

const (
	OperatorEqual ComparisonOperator = iota
	OperatorNotEqual
	OperatorGraterThan
	OperatorGraterOrEqualThan
	OperatorLessThan
	OperatorLessOrEqualThan
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

func (r *Result) DoAddition() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return &ObjInteger{Value: v.Value + r.ObjectList[1].(*ObjInteger).Value}
	case *ObjFloat:
		return &ObjFloat{Value: v.Value + r.ObjectList[1].(*ObjFloat).Value}
	}
	return nil
}

func (r *Result) DoSubtraction() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return &ObjInteger{Value: v.Value - r.ObjectList[1].(*ObjInteger).Value}
	case *ObjFloat:
		return &ObjFloat{Value: v.Value - r.ObjectList[1].(*ObjFloat).Value}
	}
	return nil
}

func (r *Result) DoMultiplication() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return &ObjInteger{Value: v.Value * r.ObjectList[1].(*ObjInteger).Value}
	case *ObjFloat:
		return &ObjFloat{Value: v.Value * r.ObjectList[1].(*ObjFloat).Value}
	}
	return nil
}

func (r *Result) DoDivision() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return &ObjInteger{Value: v.Value / r.ObjectList[1].(*ObjInteger).Value}
	case *ObjFloat:
		return &ObjFloat{Value: v.Value / r.ObjectList[1].(*ObjFloat).Value}
	}
	return nil
}

func (r *Result) DoEqual() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return ToReservedBoolObj(v.Value == r.ObjectList[1].(*ObjInteger).Value)
	case *ObjFloat:
		return ToReservedBoolObj(v.Value == r.ObjectList[1].(*ObjFloat).Value)
	case *ObjBoolean:
		return ToReservedBoolObj(v == r.ObjectList[1])
	}
	return nil
}

func (r *Result) DoNotEqual() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return ToReservedBoolObj(v.Value != r.ObjectList[1].(*ObjInteger).Value)
	case *ObjFloat:
		return ToReservedBoolObj(v.Value != r.ObjectList[1].(*ObjFloat).Value)
	case *ObjBoolean:
		return ToReservedBoolObj(v != r.ObjectList[1])
	}
	return nil
}

func (r *Result) DoGraterThan() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return ToReservedBoolObj(v.Value > r.ObjectList[1].(*ObjInteger).Value)
	case *ObjFloat:
		return ToReservedBoolObj(v.Value > r.ObjectList[1].(*ObjFloat).Value)
	}
	return nil
}

func (r *Result) DoLessThan() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return ToReservedBoolObj(v.Value < r.ObjectList[1].(*ObjInteger).Value)
	case *ObjFloat:
		return ToReservedBoolObj(v.Value < r.ObjectList[1].(*ObjFloat).Value)
	}
	return nil
}

func (r *Result) DoGraterThanOrEqual() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return ToReservedBoolObj(v.Value >= r.ObjectList[1].(*ObjInteger).Value)
	case *ObjFloat:
		return ToReservedBoolObj(v.Value >= r.ObjectList[1].(*ObjFloat).Value)
	}
	return nil
}

func (r *Result) DoLessThanOrEqual() Object {
	switch v := r.ObjectList[0].(type) {
	case *ObjInteger:
		return ToReservedBoolObj(v.Value <= r.ObjectList[1].(*ObjInteger).Value)
	case *ObjFloat:
		return ToReservedBoolObj(v.Value <= r.ObjectList[1].(*ObjFloat).Value)
	}
	return nil
}

func NewNamedResult() *NamedResult {
	return &NamedResult{
		ObjectList: make(map[string]Object),
	}
}

type NamedResult struct {
	ObjectList map[string]Object
}

func ToReservedBoolObj(value bool) *ObjBoolean {
	if value {
		return ReservedObjTrue
	}
	return ReservedObjFalse
}
