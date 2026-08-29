package gosql

import "fmt"

type BinaryExpression struct {
	A  Expression
	B  Expression
	Op Token
}

func (be BinaryExpression) GenerateCode() string {
	return fmt.Sprintf("(%s %s %s)", be.A.GenerateCode(), be.Op.Value, be.B.GenerateCode())
}

type Expression struct {
	Literal *Token
	Binary  *BinaryExpression
	Kind    ExpressionKind
}

func (e Expression) GenerateCode() string {
	switch e.Kind {
	case LiteralKind:
		switch e.Literal.Kind {
		case IdentifierKind:
			return fmt.Sprintf("\"%s\"", e.Literal.Value)
		case StringKind:
			return fmt.Sprintf("'%s'", e.Literal.Value)
		default:
			return fmt.Sprintf(e.Literal.Value)
		}

	case BinaryKind:
		return e.Binary.GenerateCode()
	}

	return ""
}
