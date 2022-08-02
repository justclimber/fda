package ast

func NewFunctionDefinition(
	name string,
	statementsBlock *StatementsBlock,
	args []*VarAndType,
	returns []*VarAndType,
) *FunctionDefinition {
	return &FunctionDefinition{
		key:             KeyFunctionDefinition,
		Name:            name,
		statementsBlock: statementsBlock,
		args:            args,
		returns:         returns,
	}
}

type FunctionDefinition struct {
	id              int64
	key             NodeKey
	Name            string
	statementsBlock *StatementsBlock
	args            []*VarAndType
	returns         []*VarAndType
}

func (fd *FunctionDefinition) ID() int64        { return fd.id }
func (fd *FunctionDefinition) NodeKey() NodeKey { return fd.key }
