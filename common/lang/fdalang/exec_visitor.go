package fdalang

type ExecAstVisitor struct {
	execCallback ExecCallback
	builtins     map[string]*ObjBuiltin
}

const (
	_ OperationType = iota
	OperationAssignment
	OperationStructFieldAssignment
	OperationReturn
	OperationIfStmt
	OperationSwitch
	OperationUnary
	OperationQuestion
	OperationBinExpr
	OperationStruct
	OperationStructFieldCall
	OperationNumInt
	OperationNumFloat
	OperationBoolean
	OperationArray
	OperationArrayIndex
	OperationIdentifier
	OperationFunction
	OperationFunctionCall
	OperationEnumElementCall
	OperationBuiltin
)

type OperationType int

type Operation struct {
	Type     OperationType
	FuncName string
}

type ExecCallback func(Operation)

func NewExecAstVisitor() *ExecAstVisitor {
	e := &ExecAstVisitor{
		execCallback: func(operation Operation) {},
		builtins:     make(map[string]*ObjBuiltin),
	}
	e.setupBasicBuiltinFunctions()
	return e
}

func (e *ExecAstVisitor) SetExecCallback(callback ExecCallback) {
	e.execCallback = callback
}

func (e *ExecAstVisitor) ExecAst(ast *AstStatementsBlock, env *Environment) error {
	_, err := e.execStatementsBlock(ast, env)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExecAstVisitor) execStatementsBlock(node *AstStatementsBlock, env *Environment) (*ObjReturnValue, error) {
	for _, statement := range node.Statements {
		returnValue, err := e.execStatement(statement, env)
		if returnValue != nil || err != nil {
			return returnValue, err
		}
	}

	return nil, nil
}

func (e *ExecAstVisitor) execStatement(node AstStatement, env *Environment) (*ObjReturnValue, error) {
	switch astNode := node.(type) {
	case *AstStatementWithVoidedExpression:
		_, err := e.execExpression(astNode.Expr, env)
		return nil, err
	case *AstReturn:
		return e.execReturn(astNode, env)
	case *AstIf:
		return e.execIfStatement(astNode, env)
	case *AstSwitch:
		return e.execSwitch(astNode, env)
	case *AstStructDefinition:
		return nil, env.RegisterStructDefinition(astNode)
	case *AstEnumDefinition:
		return nil, env.RegisterEnumDefinition(astNode)
	default:
		return nil, runtimeError(node, "Unexpected node for statement: %T", node)
	}
}

func (e *ExecAstVisitor) execExpression(node AstExpression, env *Environment) (Object, error) {
	switch astNode := node.(type) {
	case *AstAssignment:
		return e.execAssignment(astNode, env)
	case *AstStructFieldAssignment:
		return e.execStructFieldAssignment(astNode, env)
	case *AstUnary:
		return e.execUnaryExpression(astNode, env)
	case *AstEmptier:
		return e.execEmptierExpression(astNode, env)
	case *AstBinOperation:
		return e.execBinExpression(astNode, env)
	case *AstStruct:
		return e.execStruct(astNode, env)
	case *AstStructFieldCall:
		return e.execStructFieldCall(astNode, env)
	case *AstEnumElementCall:
		return e.execEnumElementCall(astNode, env)
	case *AstNumInt:
		return e.execNumInt(astNode)
	case *AstNumFloat:
		return e.execNumFloat(astNode)
	case *AstBoolean:
		return e.execBoolean(astNode)
	case *AstArray:
		return e.execArray(astNode, env)
	case *AstArrayIndexCall:
		return e.execArrayIndexCall(astNode, env)
	case *AstIdentifier:
		return e.execIdentifier(astNode, env)
	case *AstFunction:
		return e.execFunction(astNode, env)
	case *AstFunctionCall:
		return e.execFunctionCall(astNode, env)
	default:
		return nil, runtimeError(node, "Unexpected node for expression: %T", node)
	}
}

func (e *ExecAstVisitor) execAssignment(node *AstAssignment, env *Environment) (Object, error) {
	varName := node.Left.Value
	if _, exists := e.builtins[varName]; exists {
		return nil, runtimeError(node.Left, "Builtins are immutable")
	}
	e.execCallback(Operation{Type: OperationAssignment})
	value, err := e.execExpression(node.Value, env)
	if err != nil {
		return nil, err
	}

	if oldVar, isVarExist := env.Get(varName); isVarExist && oldVar.Type() != value.Type() {
		return nil, runtimeError(node.Value, "type mismatch on assignment: var type is %s and value type is %s",
			oldVar.Type(), value.Type())
	}

	env.Set(varName, value)
	return value, nil
}

func (e *ExecAstVisitor) execStructFieldAssignment(
	node *AstStructFieldAssignment,
	env *Environment,
) (Object, error) {
	e.execCallback(Operation{Type: OperationStructFieldAssignment})
	value, err := e.execExpression(node.Value, env)
	if err != nil {
		return nil, err
	}

	left, err := e.execExpression(node.Left.StructExpr, env)
	if err != nil {
		return nil, err
	}

	structObj, ok := left.(*ObjStruct)
	if !ok {
		return nil, runtimeError(node, "Field access can be only on struct but '%s' given", left.Type())
	}

	if _, ok = structObj.Fields[node.Left.Field.Value]; !ok {
		return nil, runtimeError(node,
			"Struct '%s' doesn't have field '%s'", structObj.Definition.Name, node.Left.Field.Value)
	}
	structObj.Fields[node.Left.Field.Value] = value
	return value, nil
}

func (e *ExecAstVisitor) execUnaryExpression(node *AstUnary, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationUnary})
	right, err := e.execExpression(node.Right, env)
	if err != nil {
		return nil, err
	}
	switch node.Operator {
	case TokenNot:
		boolObj, ok := right.(*ObjBoolean)
		if !ok {
			return nil, runtimeError(node, "Operator '!' could be applied only on bool, '%s' given", right.Type())
		}
		return nativeBooleanToBoolean(!boolObj.Value), nil
	case TokenMinus:
		switch right.Type() {
		case TypeInt:
			value := right.(*ObjInteger).Value
			return &ObjInteger{Value: -value}, nil
		case TypeFloat:
			value := right.(*ObjFloat).Value
			return &ObjFloat{Value: -value}, nil
		default:
			return nil, runtimeError(node, "unknown operator: -%s", right.Type())
		}
	default:
		return nil, runtimeError(node, "unknown operator: %s%s", node.Operator, right.Type())
	}
}

func (e *ExecAstVisitor) execEmptierExpression(node *AstEmptier, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationQuestion})
	if node.IsArray {
		if node.Type == TypeInt || node.Type == TypeFloat {
			return &ObjArray{Emptier: Emptier{Empty: true}, ElementsType: node.Type}, nil
		} else if _, ok := env.StructDefinition(node.Type); ok {
			return &ObjArray{Emptier: Emptier{Empty: true}, ElementsType: node.Type}, nil
		} else {
			return nil, runtimeError(node, "? is not supported on type: '%s[]'", node.Type)
		}
	} else if node.Type == TypeInt {
		return &ObjInteger{Emptier: Emptier{Empty: true}}, nil
	} else if node.Type == TypeFloat {
		return &ObjFloat{Emptier: Emptier{Empty: true}}, nil
	} else if def, ok := env.StructDefinition(node.Type); ok {
		return &ObjStruct{
			Emptier:    Emptier{Empty: true},
			Definition: def,
			Fields:     make(map[string]Object),
		}, nil
	} else {
		return nil, runtimeError(node, "? is not supported on type: '%s'", node.Type)
	}
}

func (e *ExecAstVisitor) execBinExpression(node *AstBinOperation, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationBinExpr})
	left, err := e.execExpression(node.Left, env)
	if err != nil {
		return nil, err
	}
	right, err := e.execExpression(node.Right, env)
	if err != nil {
		return nil, err
	}

	if left.Type() != right.Type() {
		return nil, runtimeError(node, "forbidden operation on different types: %s and %s",
			left.Type(), right.Type())
	}

	result, err := execScalarBinOperation(left, right, node.Operator)
	return result, err
}

func (e *ExecAstVisitor) execIdentifier(node *AstIdentifier, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationIdentifier})
	if builtin, ok := e.builtins[node.Value]; ok {
		return builtin, nil
	}

	if ed, ok := env.EnumDefinition(node.Value); ok {
		return &ObjEnum{Definition: ed}, nil
	}

	if val, ok := env.Get(node.Value); ok {
		return val, nil
	}

	return nil, runtimeError(node, "identifier not found: "+node.Value)
}

func (e *ExecAstVisitor) execReturn(node *AstReturn, env *Environment) (*ObjReturnValue, error) {
	e.execCallback(Operation{Type: OperationReturn})
	value, err := e.execExpression(node.ReturnValue, env)
	return &ObjReturnValue{Value: value}, err
}

func (e *ExecAstVisitor) execFunction(node *AstFunction, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationFunction})
	return &ObjFunction{
		Arguments:  node.Arguments,
		Statements: node.StatementsBlock,
		ReturnType: node.ReturnType,
		Env:        env,
	}, nil
}

func (e *ExecAstVisitor) execFunctionCall(node *AstFunctionCall, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationFunctionCall})
	functionObj, err := e.execExpression(node.Function, env)
	if err != nil {
		return nil, err
	}

	args, err := e.execExpressionList(node.Arguments, env)
	if err != nil {
		return nil, err
	}

	switch fn := functionObj.(type) {
	case *ObjFunction:
		err = functionCallArgumentsCheck(node, fn.Arguments, args)
		if err != nil {
			return nil, err
		}

		// todo: what is fn.Env?
		functionEnv := transferArgsToNewEnv(fn, args)
		statementsBlockResult, err := e.execStatementsBlock(fn.Statements, functionEnv)
		if err != nil {
			return nil, err
		}

		var result Object
		if statementsBlockResult == nil {
			result = &ObjVoid{}
		} else {
			result = statementsBlockResult.Value
		}

		if err = functionReturnTypeCheck(node, result, fn.ReturnType); err != nil {
			return nil, err
		}

		return result, nil

	case *ObjBuiltin:
		e.execCallback(Operation{Type: OperationBuiltin, FuncName: fn.Name})
		if err := e.checkArgs(fn, args); err != nil {
			return nil, err
		}
		result, err := fn.Fn(env, args)
		if err != nil {
			return nil, err
		}

		if err = functionReturnTypeCheck(node, result, fn.ReturnType); err != nil {
			return nil, err
		}

		return result, nil

	default:
		return nil, runtimeError(node, "not a function: %s", fn.Type())
	}
}
func (e *ExecAstVisitor) execExpressionList(expressions []AstExpression, env *Environment) ([]Object, error) {
	var result []Object

	for _, expr := range expressions {
		evaluated, err := e.execExpression(expr, env)
		if err != nil {
			return nil, err
		}
		result = append(result, evaluated)
	}

	return result, nil
}

func (e *ExecAstVisitor) execIfStatement(node *AstIf, env *Environment) (*ObjReturnValue, error) {
	e.execCallback(Operation{Type: OperationIfStmt})
	condition, err := e.execExpression(node.Condition, env)
	if err != nil {
		return nil, err
	}
	if condition.Type() != TypeBool {
		return nil, runtimeError(node, "Condition should be boolean type but %s in fact", condition.Type())
	}

	if condition == ReservedObjTrue {
		return e.execStatementsBlock(node.PositiveBranch, env)
	} else if node.ElseBranch != nil {
		return e.execStatementsBlock(node.ElseBranch, env)
	} else {
		return nil, nil
	}
}

func (e *ExecAstVisitor) execArray(node *AstArray, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationArray})
	elements, err := e.execExpressionList(node.Elements, env)
	if err != nil {
		return nil, err
	}
	if err = arrayElementsTypeCheck(node, node.ElementsType, elements); err != nil {
		return nil, err
	}

	return &ObjArray{
		ElementsType: node.ElementsType,
		Elements:     elements,
	}, nil
}

func (e *ExecAstVisitor) execArrayIndexCall(node *AstArrayIndexCall, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationArrayIndex})
	left, err := e.execExpression(node.Left, env)
	if err != nil {
		return nil, err
	}

	index, err := e.execExpression(node.Index, env)
	if err != nil {
		return nil, err
	}

	arrayObj, ok := left.(*ObjArray)
	if !ok {
		return nil, runtimeError(node, "Array access can be only on arrays but '%s' given", left.Type())
	}

	indexObj, ok := index.(*ObjInteger)
	if !ok {
		return nil, runtimeError(node, "Array access can be only by 'int' type but '%s' given", index.Type())
	}

	i := indexObj.Value
	if i < 0 || int(i) > len(arrayObj.Elements)-1 {
		return nil, runtimeError(node, "Array access out of bounds: '%d'", i)
	}

	return arrayObj.Elements[i], nil
}

func (e *ExecAstVisitor) execStruct(node *AstStruct, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationStruct})
	definition, ok := env.StructDefinition(node.Ident.Value)
	if !ok {
		return nil, runtimeError(node, "Struct '%s' is not defined", node.Ident.Value)
	}
	fields := make(map[string]Object)
	for _, n := range node.Fields {
		result, err := e.execExpression(n.Value, env)
		if err != nil {
			return nil, err
		}

		if err = structTypeAndVarsChecks(n, definition, result); err != nil {
			return nil, err
		}

		fields[n.Left.Value] = result
	}
	if len(fields) != len(definition.Fields) {
		return nil, runtimeError(node,
			"Var of struct '%s' should have %d fields filled but in fact only %d",
			definition.Name,
			len(definition.Fields),
			len(fields))
	}
	obj := &ObjStruct{
		Definition: definition,
		Fields:     fields,
	}

	return obj, nil
}

func (e *ExecAstVisitor) execStructFieldCall(node *AstStructFieldCall, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationStructFieldCall})
	left, err := e.execExpression(node.StructExpr, env)
	if err != nil {
		return nil, err
	}

	structObj, ok := left.(*ObjStruct)
	if !ok {
		return nil, runtimeError(node, "Field access can be only on struct but '%s' given", left.Type())
	}

	fieldObj, ok := structObj.Fields[node.Field.Value]
	if !ok {
		return nil, runtimeError(node,
			"Struct '%s' doesn't have field '%s'", structObj.Definition.Name, node.Field.Value)
	}

	return fieldObj, nil
}

func (e *ExecAstVisitor) execEnumElementCall(node *AstEnumElementCall, env *Environment) (Object, error) {
	e.execCallback(Operation{Type: OperationEnumElementCall})
	left, err := e.execExpression(node.EnumExpr, env)
	if err != nil {
		return nil, err
	}

	enumObj, ok := left.(*ObjEnum)
	if !ok {
		return nil, runtimeError(node, "Expected enum, got '%s'", left.Type())
	}

	found := false
	for value, str := range enumObj.Definition.Elements {
		if node.Element.Value == str {
			found = true
			enumObj.Value = int8(value)
			break
		}
	}
	if !found {
		return nil, runtimeError(node,
			"Enum '%s' doesn't have element '%s'", enumObj.Definition.Name, node.Element.Value)
	}

	return enumObj, nil
}

func (e *ExecAstVisitor) execSwitch(node *AstSwitch, env *Environment) (*ObjReturnValue, error) {
	e.execCallback(Operation{Type: OperationSwitch})
	for _, c := range node.Cases {
		condition, err := e.execExpression(c.Condition, env)
		if err != nil {
			return nil, err
		}
		if condition.Type() != TypeBool {
			return nil, runtimeError(c.Condition,
				"Result of case condition should be 'boolean' but '%s' given", condition.Type())
		}
		conditionResult, _ := condition.(*ObjBoolean)
		if conditionResult.Value {
			return e.execStatementsBlock(c.PositiveBranch, env)
		}
	}
	if node.DefaultBranch != nil {
		return e.execStatementsBlock(node.DefaultBranch, env)
	}
	return nil, nil
}

func (e *ExecAstVisitor) execNumInt(node *AstNumInt) (Object, error) {
	e.execCallback(Operation{Type: OperationNumInt})
	return &ObjInteger{Value: node.Value}, nil
}

func (e *ExecAstVisitor) execNumFloat(node *AstNumFloat) (Object, error) {
	e.execCallback(Operation{Type: OperationNumFloat})
	return &ObjFloat{Value: node.Value}, nil
}

func (e *ExecAstVisitor) execBoolean(node *AstBoolean) (Object, error) {
	e.execCallback(Operation{Type: OperationBoolean})
	return nativeBooleanToBoolean(node.Value), nil
}
