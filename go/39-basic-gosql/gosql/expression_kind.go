package gosql

type ExpressionKind uint

const (
	LiteralKind ExpressionKind = iota
	BinaryKind
)
