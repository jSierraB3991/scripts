package gosql

type Token struct {
	Value string
	Kind  TokenKind
	Loc   Location
}

func (t Token) bindingPower() uint {
	switch t.Kind {
	case KeywordKind:
		switch Keyword(t.Value) {
		case AndKeyword:
			fallthrough
		case OrKeyword:
			return 1
		}
	case SymbolKind:
		switch Symbol(t.Value) {
		case EqSymbol:
			fallthrough
		case NeqSymbol:
			return 2

		case LtSymbol:
			fallthrough
		case GtSymbol:
			return 3

		// For some reason these are grouped separately
		case LteSymbol:
			fallthrough
		case GteSymbol:
			return 4

		case ConcatSymbol:
			fallthrough
		case PlusSymbol:
			return 5
		}
	}

	return 0
}

func (t *Token) equals(other *Token) bool {
	return t.Value == other.Value && t.Kind == other.Kind
}
