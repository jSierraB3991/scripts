package gosql

import (
	"fmt"

	"github.com/petar/GoLLRB/llrb"
)

type MemoryBackend struct {
	tables map[string]*table
}

func (mb *MemoryBackend) Select(slct *SelectStatement) (*Results, error) {
	t := createTable()

	if slct.From != nil {
		var ok bool
		t, ok = mb.tables[slct.From.Value]
		if !ok {
			return nil, ErrTableDoesNotExist
		}
	}

	if slct.Item == nil || len(*slct.Item) == 0 {
		return &Results{}, nil
	}

	results := [][]Cell{}
	columns := []ResultColumn{}

	if slct.From == nil {
		t = createTable()
		t.rows = [][]memoryCell{{}}
	}

	for _, iAndE := range t.getApplicableIndexes(slct.Where) {
		index := iAndE.i
		exp := iAndE.e
		t = index.newTableFromSubset(t, exp)
	}

	// Expand SELECT * at the AST level into a SELECT on all columns
	finalItems := []*SelectItem{}
	for _, item := range *slct.Item {
		if item.Asterisk {
			newItems := []*SelectItem{}
			for j := 0; j < len(t.columns); j++ {
				newSelectItem := &SelectItem{
					Exp: &Expression{
						Literal: &Token{
							Value: t.columns[j],
							Kind:  IdentifierKind,
							Loc:   Location{0, uint(len("SELECT") + 1)},
						},
						Binary: nil,
						Kind:   LiteralKind,
					},
					Asterisk: false,
					As:       nil,
				}
				newItems = append(newItems, newSelectItem)
			}
			finalItems = append(finalItems, newItems...)
		} else {
			finalItems = append(finalItems, item)
		}
	}

	limit := len(t.rows)
	if slct.Limit != nil {
		v, _, _, err := t.evaluateCell(0, *slct.Limit)
		if err != nil {
			return nil, err
		}

		limit = int(*v.AsInt())
	}
	if limit < 0 {
		return nil, fmt.Errorf("Invalid, negative limit")
	}

	offset := 0
	if slct.Offset != nil {
		v, _, _, err := t.evaluateCell(0, *slct.Offset)
		if err != nil {
			return nil, err
		}

		offset = int(*v.AsInt())
	}
	if offset < 0 {
		return nil, fmt.Errorf("Invalid, negative limit")
	}

	rowIndex := -1
	for i := range t.rows {
		result := []Cell{}
		isFirstRow := len(results) == 0

		if slct.Where != nil {
			val, _, _, err := t.evaluateCell(uint(i), *slct.Where)
			if err != nil {
				return nil, err
			}

			if !*val.AsBool() {
				continue
			}
		}

		rowIndex++
		if rowIndex < offset {
			continue
		} else if rowIndex > offset+limit-1 {
			break
		}

		for _, col := range finalItems {
			value, columnName, columnType, err := t.evaluateCell(uint(i), *col.Exp)
			if err != nil {
				return nil, err
			}

			if isFirstRow {
				columns = append(columns, ResultColumn{
					Type: columnType,
					Name: columnName,
				})
			}

			result = append(result, value)
		}

		results = append(results, result)
	}

	return &Results{
		Columns: columns,
		Rows:    results,
	}, nil
}

func (mb *MemoryBackend) Insert(inst *InsertStatement) error {
	t, ok := mb.tables[inst.Table.Value]
	if !ok {
		return ErrTableDoesNotExist
	}

	if inst.Values == nil {
		return nil
	}

	if len(*inst.Values) != len(t.columns) {
		return ErrMissingValues
	}

	row := []memoryCell{}
	for _, valueNode := range *inst.Values {
		if valueNode.Kind != LiteralKind {
			fmt.Println("Skipping non-literal.")
			continue
		}

		emptyTable := createTable()
		value, _, _, err := emptyTable.evaluateCell(0, *valueNode)
		if err != nil {
			return err
		}

		row = append(row, value)
	}

	t.rows = append(t.rows, row)

	for _, index := range t.indexes {
		err := index.addRow(t, uint(len(t.rows)-1))
		if err != nil {
			// Drop the row on failure
			t.rows = t.rows[:len(t.rows)-1]
			return err
		}
	}

	return nil
}

func (mb *MemoryBackend) CreateTable(crt *CreateTableStatement) error {
	if _, ok := mb.tables[crt.Name.Value]; ok {
		return ErrTableAlreadyExists
	}

	t := createTable()
	t.name = crt.Name.Value
	mb.tables[t.name] = t
	if crt.Cols == nil {
		return nil
	}

	var primaryKey *Expression = nil
	for _, col := range *crt.Cols {
		t.columns = append(t.columns, col.Name.Value)

		var dt ColumnType
		switch col.Datatype.Value {
		case "int":
			dt = IntType
		case "text":
			dt = TextType
		case "boolean":
			dt = BoolType
		default:
			delete(mb.tables, t.name)
			return ErrInvalidDatatype
		}

		if col.PrimaryKey {
			if primaryKey != nil {
				delete(mb.tables, t.name)
				return ErrPrimaryKeyAlreadyExists
			}

			primaryKey = &Expression{
				Literal: &col.Name,
				Kind:    LiteralKind,
			}
		}

		t.columnTypes = append(t.columnTypes, dt)
	}

	if primaryKey != nil {
		err := mb.CreateIndex(&CreateIndexStatement{
			Table:      crt.Name,
			Name:       Token{Value: t.name + "_pkey"},
			Unique:     true,
			PrimaryKey: true,
			Exp:        *primaryKey,
		})
		if err != nil {
			delete(mb.tables, t.name)
			return err
		}
	}

	return nil
}

func (mb *MemoryBackend) CreateIndex(ci *CreateIndexStatement) error {
	table, ok := mb.tables[ci.Table.Value]
	if !ok {
		return ErrTableDoesNotExist
	}

	for _, index := range table.indexes {
		if index.name == ci.Name.Value {
			return ErrIndexAlreadyExists
		}
	}

	index := &index{
		exp:        ci.Exp,
		unique:     ci.Unique,
		primaryKey: ci.PrimaryKey,
		name:       ci.Name.Value,
		tree:       llrb.New(),
		typ:        "rbtree",
	}
	table.indexes = append(table.indexes, index)

	for i := range table.rows {
		err := index.addRow(table, uint(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func (mb *MemoryBackend) DropTable(dt *DropTableStatement) error {
	if _, ok := mb.tables[dt.Name.Value]; ok {
		delete(mb.tables, dt.Name.Value)
		return nil
	}
	return ErrTableDoesNotExist
}

func (mb *MemoryBackend) GetTables() []TableMetadata {
	tms := []TableMetadata{}
	for name, t := range mb.tables {
		tm := TableMetadata{}
		tm.Name = name

		pkeyColumn := ""
		for _, i := range t.indexes {
			if i.primaryKey {
				pkeyColumn = i.exp.GenerateCode()
			}

			tm.Indexes = append(tm.Indexes, Index{
				Name:       i.name,
				Type:       i.typ,
				Unique:     i.unique,
				PrimaryKey: i.primaryKey,
				Exp:        i.exp.GenerateCode(),
			})
		}

		for i, column := range t.columns {
			tm.Columns = append(tm.Columns, ResultColumn{
				Type:    t.columnTypes[i],
				Name:    column,
				NotNull: pkeyColumn == `"`+column+`"`,
			})
		}

		tms = append(tms, tm)
	}

	return tms
}

func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{
		tables: map[string]*table{},
	}
}
