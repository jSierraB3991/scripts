package gosql

import (
	"bytes"

	"github.com/petar/GoLLRB/llrb"
)

type treeItem struct {
	value memoryCell
	index uint
}

func (te treeItem) Less(than llrb.Item) bool {
	return bytes.Compare(te.value, than.(treeItem).value) < 0
}
