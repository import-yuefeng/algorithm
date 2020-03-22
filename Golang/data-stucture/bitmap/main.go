package main

import (
	"fmt"
)

type bitmap struct {
	d      []byte
	length int
}

func NewBitMap(length int) *bitmap {
	if length <= 0 {
		return nil
	}
	return &bitmap{
		d:      make([]byte, length/8+1),
		length: length/8 + 1,
	}
}

func (b *bitmap) Set(val int) {
	byteIndex := val / 8
	if byteIndex > len(b.d)-1 {
		return
	}
	bitIndex := uint(val % 8)
	b.d[byteIndex] |= (1 << bitIndex)
}

func (b *bitmap) Get(val int) bool {
	byteIndex := val / 8
	if byteIndex >= len(b.d) {
		return false
	}
	bitIndex := uint(val % 8)
	if b.d[byteIndex]>>bitIndex == 1 {
		return true
	}
	return false
}

func main() {
	bm := NewBitMap(100)
	bm.Set(10)
	r := bm.Get(2)
	fmt.Println(r)
	r = bm.Get(10)
	fmt.Println(r)
}
