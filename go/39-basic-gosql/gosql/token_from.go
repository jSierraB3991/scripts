package gosql

func tokenFromKeyword(k Keyword) Token {
	return Token{
		Kind:  KeywordKind,
		Value: string(k),
	}
}

func tokenFromSymbol(s Symbol) Token {
	return Token{
		Kind:  SymbolKind,
		Value: string(s),
	}
}
