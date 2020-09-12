package main

import (
	"hash/adler32"
	"io"
)

func Hash(key string) uint32 {
	hs := adler32.New()
	io.WriteString(hs, key)
	return hs.Sum32()
}
