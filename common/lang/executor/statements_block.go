package executor

func NewStatementsBlock(stmts []Stmt) *StatementsBlock {
	return &StatementsBlock{
		key:        KeyStatementsBlock,
		statements: stmts,
	}
}

type StatementsBlock struct {
	id         int64
	key        NodeKey
	statements []Stmt
}

func (sb *StatementsBlock) ID() int64        { return sb.id }
func (sb *StatementsBlock) NodeKey() NodeKey { return sb.key }

func (sb *StatementsBlock) Exec(env *Environment, executor execManager) error {
	for _, statement := range sb.statements {
		err := statement.Exec(env, executor)
		if err != nil {
			return err
		}
	}
	return nil
}
