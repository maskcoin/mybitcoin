package main

import (
	"bytes"
	"encoding/binary"
)

func Uint64ToByteSlice(num uint64) (ret []byte) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	ret = buf.Bytes()

	return
}
