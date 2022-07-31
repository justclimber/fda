package executor

func NewUnaryMinus(value Expr) *UnaryMinus {
	return &UnaryMinus{
		key:   KeyUnaryMinus,
		value: value,
	}
}

type UnaryMinus struct {
	id    int64
	key   NodeKey
	value Expr
}

func (um *UnaryMinus) ID() int64        { return um.id }
func (um *UnaryMinus) NodeKey() NodeKey { return um.key }

func (um *UnaryMinus) Exec(env *Environment, result *Result, executor execManager) error {
	executor.AddNextExec(um.value, func() error {
		return um.value.Exec(env, result, executor)
	})
	executor.AddNextExec(um, func() error {
		switch obj := result.objectList[0].(type) {
		case *ObjInteger:
			obj.Value = -obj.Value
		}
		return nil
	})

	return nil
}
