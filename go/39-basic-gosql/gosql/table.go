package gosql

import "strconv"

type table struct {
	name        string
	columns     []string
	columnTypes []ColumnType
	rows        [][]memoryCell
	indexes     []*index
}

func createTable() *table {
	return &table{
		name:        "?tmp?",
		columns:     nil,
		columnTypes: nil,
		rows:        nil,
		indexes:     []*index{},
	}
}

func (t *table) evaluateLiteralCell(rowIndex uint, exp Expression) (memoryCell, string, ColumnType, error) {
	if exp.Kind != LiteralKind {
		return nil, "", 0, ErrInvalidCell
	}

	lit := exp.Literal
	if lit.Kind == IdentifierKind {
		for i, tableCol := range t.columns {
			if tableCol == lit.Value {
				return t.rows[rowIndex][i], tableCol, t.columnTypes[i], nil
			}
		}

		return nil, "", 0, ErrColumnDoesNotExist
	}

	columnType := IntType
	if lit.Kind == StringKind {
		columnType = TextType
	} else if lit.Kind == BoolKind {
		columnType = BoolType
	}

	return literalToMemoryCell(lit), "?column?", columnType, nil
}

func (t *table) evaluateBinaryCell(rowIndex uint, exp Expression) (memoryCell, string, ColumnType, error) {
	if exp.Kind != BinaryKind {
		return nil, "", 0, ErrInvalidCell
	}

	bexp := exp.Binary

	l, _, lt, err := t.evaluateCell(rowIndex, bexp.A)
	if err != nil {
		return nil, "", 0, err
	}

	r, _, rt, err := t.evaluateCell(rowIndex, bexp.B)
	if err != nil {
		return nil, "", 0, err
	}

	switch bexp.Op.Kind {
	case SymbolKind:
		switch Symbol(bexp.Op.Value) {
		case EqSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", BoolType, nil
			}

			eq := l.equals(r)
			if lt == TextType && rt == TextType && eq {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			if lt == IntType && rt == IntType && eq {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			if lt == BoolType && rt == BoolType && eq {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			return falseMemoryCell, "?column?", BoolType, nil
		case NeqSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", BoolType, nil
			}

			if lt != rt || !l.equals(r) {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			return falseMemoryCell, "?column?", BoolType, nil
		case ConcatSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", TextType, nil
			}

			if lt != TextType || rt != TextType {
				return nil, "", 0, ErrInvalidOperands
			}

			return literalToMemoryCell(&Token{Kind: StringKind, Value: *l.AsText() + *r.AsText()}), "?column?", TextType, nil
		case PlusSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", IntType, nil
			}

			if lt != IntType || rt != IntType {
				return nil, "", 0, ErrInvalidOperands
			}

			iValue := int(*l.AsInt() + *r.AsInt())
			return literalToMemoryCell(&Token{Kind: NumericKind, Value: strconv.Itoa(iValue)}), "?column?", IntType, nil
		case LtSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", BoolType, nil
			}

			if lt != IntType || rt != IntType {
				return nil, "", 0, ErrInvalidOperands
			}

			if *l.AsInt() < *r.AsInt() {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			return falseMemoryCell, "?column?", BoolType, nil
		case LteSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", BoolType, nil
			}

			if lt != IntType || rt != IntType {
				return nil, "", 0, ErrInvalidOperands
			}

			if *l.AsInt() <= *r.AsInt() {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			return falseMemoryCell, "?column?", BoolType, nil
		case GtSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", BoolType, nil
			}

			if lt != IntType || rt != IntType {
				return nil, "", 0, ErrInvalidOperands
			}

			if *l.AsInt() > *r.AsInt() {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			return falseMemoryCell, "?column?", BoolType, nil
		case GteSymbol:
			if len(l) == 0 || len(r) == 0 {
				return nullMemoryCell, "?column?", BoolType, nil
			}

			if lt != IntType || rt != IntType {
				return nil, "", 0, ErrInvalidOperands
			}

			if *l.AsInt() >= *r.AsInt() {
				return trueMemoryCell, "?column?", BoolType, nil
			}

			return falseMemoryCell, "?column?", BoolType, nil
		default:
			// TODO
			break
		}
	case KeywordKind:
		switch Keyword(bexp.Op.Value) {
		case AndKeyword:
			res := falseMemoryCell
			if lt != BoolType || rt != BoolType {
				return nil, "", 0, ErrInvalidOperands
			}

			if len(l) == 0 || len(r) == 0 {
				res = nullMemoryCell
			} else if *l.AsBool() && *r.AsBool() {
				res = trueMemoryCell
			}

			return res, "?column?", BoolType, nil
		case OrKeyword:
			res := falseMemoryCell
			if lt != BoolType || rt != BoolType {
				return nil, "", 0, ErrInvalidOperands
			}

			if len(l) == 0 || len(r) == 0 {
				res = nullMemoryCell
			} else if *l.AsBool() || *r.AsBool() {
				res = trueMemoryCell
			}

			return res, "?column?", BoolType, nil
		default:
			// TODO
			break
		}
	}

	return nil, "", 0, ErrInvalidCell
}

func (t *table) evaluateCell(rowIndex uint, exp Expression) (memoryCell, string, ColumnType, error) {
	switch exp.Kind {
	case LiteralKind:
		return t.evaluateLiteralCell(rowIndex, exp)
	case BinaryKind:
		return t.evaluateBinaryCell(rowIndex, exp)
	default:
		return nil, "", 0, ErrInvalidCell
	}
}

type indexAndExpression struct {
	i *index
	e Expression
}

func (t *table) getApplicableIndexes(where *Expression) []indexAndExpression {
	var linearizeExpressions func(where *Expression, exps []Expression) []Expression
	linearizeExpressions = func(where *Expression, exps []Expression) []Expression {
		if where == nil || where.Kind != BinaryKind {
			return exps
		}

		if where.Binary.Op.Value == string(OrKeyword) {
			return exps
		}

		if where.Binary.Op.Value == string(AndKeyword) {
			exps := linearizeExpressions(&where.Binary.A, exps)
			return linearizeExpressions(&where.Binary.B, exps)
		}

		return append(exps, *where)
	}

	exps := linearizeExpressions(where, []Expression{})

	iAndE := []indexAndExpression{}
	for _, exp := range exps {
		for _, index := range t.indexes {
			if index.applicableValue(exp) != nil {
				iAndE = append(iAndE, indexAndExpression{
					i: index,
					e: exp,
				})
			}
		}
	}

	return iAndE
}
