package gosql

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

type memoryCell []byte

func (mc memoryCell) AsInt() *int32 {
	if len(mc) == 0 {
		return nil
	}

	var i int32
	err := binary.Read(bytes.NewBuffer(mc), binary.BigEndian, &i)
	if err != nil {
		fmt.Printf("Corrupted data [%s]: %s\n", mc, err)
		return nil
	}

	return &i
}

func (mc memoryCell) AsText() *string {
	if len(mc) == 0 {
		return nil
	}

	s := string(mc)
	return &s
}

func (mc memoryCell) AsBool() *bool {
	if len(mc) == 0 {
		return nil
	}

	b := mc[0] == 1
	return &b
}

func (mc memoryCell) equals(b memoryCell) bool {
	// Seems verbose but need to make sure if one is nil, the
	// comparison still fails quickly
	if mc == nil || b == nil {
		return mc == nil && b == nil
	}

	return bytes.Equal(mc, b)
}

func literalToMemoryCell(t *Token) memoryCell {
	if t.Kind == NumericKind {
		buf := new(bytes.Buffer)
		i, err := strconv.Atoi(t.Value)
		if err != nil {
			fmt.Printf("Corrupted data [%s]: %s\n", t.Value, err)
			return nil
		}

		// TODO: handle bigint
		err = binary.Write(buf, binary.BigEndian, int32(i))
		if err != nil {
			fmt.Printf("Corrupted data [%s]: %s\n", buf.String(), err)
			return nil
		}
		return buf.Bytes()
	}

	if t.Kind == StringKind {
		return memoryCell(t.Value)
	}

	if t.Kind == BoolKind {
		if t.Value == "true" {
			return []byte{1}
		}

		return []byte{0}
	}

	return nil
}
