package gosql

type ColumnType uint

const (
	TextType ColumnType = iota
	IntType
	BoolType
)

func (c ColumnType) String() string {
	switch c {
	case TextType:
		return "TextType"
	case IntType:
		return "IntType"
	case BoolType:
		return "BoolType"
	default:
		return "Error"
	}
}
