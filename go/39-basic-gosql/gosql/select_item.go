package gosql

import (
	"fmt"
	"strings"
)

type SelectItem struct {
	Exp      *Expression
	Asterisk bool // for *
	As       *Token
}

type SelectStatement struct {
	Item   *[]*SelectItem
	From   *Token
	Where  *Expression
	Limit  *Expression
	Offset *Expression
}

func (ss SelectStatement) GenerateCode() string {
	item := []string{}
	for _, i := range *ss.Item {
		s := "\t*"
		if !i.Asterisk {
			s = "\t" + i.Exp.GenerateCode()

			if i.As != nil {
				s = fmt.Sprintf("\t%s AS \"%s\"", s, i.As.Value)
			}
		}
		item = append(item, s)
	}

	code := "SELECT\n" + strings.Join(item, ",\n")
	if ss.From != nil {
		code += fmt.Sprintf("\nFROM\n\t\"%s\"", ss.From.Value)
	}

	if ss.Where != nil {
		code += "\nWHERE\n\t" + ss.Where.GenerateCode()
	}

	if ss.Limit != nil {
		code += "\nLIMIT\n\t" + ss.Limit.GenerateCode()
	}

	if ss.Offset != nil {
		code += "\nOFFSET\n\t" + ss.Limit.GenerateCode()
	}

	return code + ";"
}
