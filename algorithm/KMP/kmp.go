package main

// package kmp

import (
	"errors"
	"fmt"
)

type kmp struct {
	pat  string
	data [][256]int
}

func (k *kmp) Init(pat string) error {
	if len(pat) <= 0 {
		return errors.New("pat is invaild")
	}
	k.pat = pat
	M := len(pat)
	X := 0
	k.data = make([][256]int, M)
	k.data[0][pat[0]] = 1

	for i := 1; i < M; i++ {
		tmpIndex := k.pat[i]
		for a := 0; a < 256; a++ {
			if int(tmpIndex) == a {
				k.data[i][a] = i + 1
			} else {
				k.data[i][a] = k.data[X][a]
			}
		}
		X = k.data[X][tmpIndex]
	}
	return nil
}

func (k *kmp) search(text string) int {
	if len(text) < len(k.pat) {
		return -1
	} else if (len(text) > len(k.pat) && len(k.pat) == 0) || (len(text) == len(k.pat) && len(text) == 0) {
		return 0
	} else {
		j := 0
		for i := 0; i < len(text); i++ {
			tmpIndex := text[i]
			j = k.data[j][tmpIndex]
			if j == len(k.pat) {
				return i - j + 1
			}
		}
	}
	return -1
}

func main() {
	kmper := new(kmp)
	kmper.Init("issip")
	fmt.Println(kmper.search("mississippi"))
}
