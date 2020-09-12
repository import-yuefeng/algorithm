package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	prev, cur, next := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)
	prev <- true
	wg.Add(3)
	go printCharByGoroutine("a", prev, cur, wg, 3)
	go printCharByGoroutine("b", cur, next, wg, 3)
	go printCharByGoroutine("c", next, prev, wg, 3)
	wg.Wait()
}

func printCharByGoroutine(char string, cur, next chan bool, wg *sync.WaitGroup, num int) {
	count := 0
	time.Sleep(1)
	for range cur {
		fmt.Printf("%s ", char)
		next <- true
		count++
		if count == num {
			wg.Done()
			return
		}
	}
}
