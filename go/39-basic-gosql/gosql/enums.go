package gosql

// for storing SQL syntax
type Symbol string

const (
	SemicolonSymbol  Symbol = ";"
	AsteriskSymbol   Symbol = "*"
	CommaSymbol      Symbol = ","
	LeftParenSymbol  Symbol = "("
	RightParenSymbol Symbol = ")"
	EqSymbol         Symbol = "="
	NeqSymbol        Symbol = "<>"
	NeqSymbol2       Symbol = "!="
	ConcatSymbol     Symbol = "||"
	PlusSymbol       Symbol = "+"
	LtSymbol         Symbol = "<"
	LteSymbol        Symbol = "<="
	GtSymbol         Symbol = ">"
	GteSymbol        Symbol = ">="
)

type TokenKind uint

const (
	KeywordKind TokenKind = iota
	SymbolKind
	IdentifierKind
	StringKind
	NumericKind
	BoolKind
	NullKind
)
