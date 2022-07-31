package executor

func NewAssignment(left []*Identifier, value Expr) *Assignment {
	return &Assignment{
		key:   KeyAssignment,
		left:  left,
		value: value,
	}
}

type Assignment struct {
	id    int64
	key   NodeKey
	left  []*Identifier
	value Expr
}

func (a *Assignment) NodeKey() NodeKey { return a.key }
func (a *Assignment) ID() int64        { return a.id }

func (a *Assignment) Exec(env *Environment, result *Result, executor execManager) error {
	executor.AddNextExec(a.value, func() error {
		return a.value.Exec(env, result, executor)
	})
	for i := range a.left {
		ii := i
		executor.AddNextExec(a.left[ii], func() error {
			varName := a.left[ii].value
			env.Set(varName, result.objectList[ii])
			return nil
		})
	}
	return nil
}
