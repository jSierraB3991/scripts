package gosql

var (
	trueToken  = Token{Kind: BoolKind, Value: "true"}
	falseToken = Token{Kind: BoolKind, Value: "false"}

	trueMemoryCell  = literalToMemoryCell(&trueToken)
	falseMemoryCell = literalToMemoryCell(&falseToken)
	nullMemoryCell  = literalToMemoryCell(&Token{Kind: NullKind})
)
