package gosql

import (
	"bytes"
	"fmt"

	"github.com/petar/GoLLRB/llrb"
)

type index struct {
	name       string
	exp        Expression
	unique     bool
	primaryKey bool
	tree       *llrb.LLRB
	typ        string
}

func (i *index) addRow(t *table, rowIndex uint) error {
	indexValue, _, _, err := t.evaluateCell(rowIndex, i.exp)
	if err != nil {
		return err
	}

	if indexValue == nil {
		return ErrViolatesNotNullConstraint
	}

	if i.unique && i.tree.Has(treeItem{value: indexValue}) {
		return ErrViolatesUniqueConstraint
	}

	i.tree.InsertNoReplace(treeItem{
		value: indexValue,
		index: rowIndex,
	})
	return nil
}

func (i *index) applicableValue(exp Expression) *Expression {
	if exp.Kind != BinaryKind {
		return nil
	}

	be := exp.Binary
	// Find the column and the value in the binary Expression
	columnExp := be.A
	valueExp := be.B
	if columnExp.GenerateCode() != i.exp.GenerateCode() {
		columnExp = be.B
		valueExp = be.A
	}

	// Neither side is applicable, return nil
	if columnExp.GenerateCode() != i.exp.GenerateCode() {
		return nil
	}

	supportedChecks := []Symbol{EqSymbol, NeqSymbol, GtSymbol, GteSymbol, LtSymbol, LteSymbol}
	supported := false
	for _, sym := range supportedChecks {
		if string(sym) == be.Op.Value {
			supported = true
			break
		}
	}
	if !supported {
		return nil
	}

	if valueExp.Kind != LiteralKind {
		fmt.Println("Only index checks on literals supported")
		return nil
	}

	return &valueExp
}

func (i *index) newTableFromSubset(t *table, exp Expression) *table {
	valueExp := i.applicableValue(exp)
	if valueExp == nil {
		return t
	}

	value, _, _, err := createTable().evaluateCell(0, *valueExp)
	if err != nil {
		fmt.Println(err)
		return t
	}

	tiValue := treeItem{value: value}

	indexes := []uint{}
	switch Symbol(exp.Binary.Op.Value) {
	case EqSymbol:
		i.tree.AscendGreaterOrEqual(tiValue, func(i llrb.Item) bool {
			ti := i.(treeItem)

			if !bytes.Equal(ti.value, value) {
				return false
			}

			indexes = append(indexes, ti.index)
			return true
		})
	case NeqSymbol:
		i.tree.AscendGreaterOrEqual(llrb.Inf(-1), func(i llrb.Item) bool {
			ti := i.(treeItem)
			if bytes.Equal(ti.value, value) {
				indexes = append(indexes, ti.index)
			}

			return true
		})
	case LtSymbol:
		i.tree.DescendLessOrEqual(tiValue, func(i llrb.Item) bool {
			ti := i.(treeItem)
			if bytes.Compare(ti.value, value) < 0 {
				indexes = append(indexes, ti.index)
			}

			return true
		})
	case LteSymbol:
		i.tree.DescendLessOrEqual(tiValue, func(i llrb.Item) bool {
			ti := i.(treeItem)
			if bytes.Compare(ti.value, value) <= 0 {
				indexes = append(indexes, ti.index)
			}

			return true
		})
	case GtSymbol:
		i.tree.AscendGreaterOrEqual(tiValue, func(i llrb.Item) bool {
			ti := i.(treeItem)
			if bytes.Compare(ti.value, value) > 0 {
				indexes = append(indexes, ti.index)
			}

			return true
		})
	case GteSymbol:
		i.tree.AscendGreaterOrEqual(tiValue, func(i llrb.Item) bool {
			ti := i.(treeItem)
			if bytes.Compare(ti.value, value) >= 0 {
				indexes = append(indexes, ti.index)
			}

			return true
		})
	}

	newT := createTable()
	newT.columns = t.columns
	newT.columnTypes = t.columnTypes
	newT.indexes = t.indexes
	newT.rows = [][]memoryCell{}

	for _, index := range indexes {
		newT.rows = append(newT.rows, t.rows[index])
	}

	return newT
}
