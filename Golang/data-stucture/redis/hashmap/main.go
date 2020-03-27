package main

import (
	"fmt"
)

func main() {
	hm := NewHashmap()

	hm.Set("123", "123")
	hm.Set("234", "234")
	hm.Set("3", "3")
	hm.Set("4", "4")
	hm.Set("5", "5")
	hm.Set("6", "6")
	hm.Set("7", "7")

	for i := 0; i < 100; i++ {
		r, e := hm.Get("2")
		if e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r)
		}
	}
	fmt.Println(hm.size, hm.used, hm.oldBucket, hm.newBucket)
}
