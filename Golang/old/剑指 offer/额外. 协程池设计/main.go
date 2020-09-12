package main

import (
	"fmt"
	"sync"
)

func main() {
	producter(100)
}

func producter(workerNum int) {
	if workerNum <= 0 {
		return
	}

	var wg sync.WaitGroup

	task := make(chan *string)
	for i := 0; i < workerNum; i++ {
		go worker(task, &wg)
		wg.Add(1)
	}
	tsk := "hello world"
	for i := 0; i < workerNum; i++ {
		task <- &tsk
	}
	close(task)
	wg.Wait()
	fmt.Println("end...")
}

func worker(task chan *string, wg *sync.WaitGroup) {
	for v := range task {
		fmt.Println(*v)
	}
	fmt.Println("task channel is closed")
	wg.Done()
}
