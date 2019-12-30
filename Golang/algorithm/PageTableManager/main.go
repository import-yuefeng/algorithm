package main

type Process struct {
	totalMemoryRequirement int
	PT                     *PageTable
	requestAddressList     []int
}

type Memory struct {
	allSize       int // unmutable
	block         []bool
	blockSize     int // unmutable
	blockFrequent []int
}

type PageTable struct {
	pageNum      int // const value == 4
	page2Block   []int
	isWrite      []bool
	flag         []bool
	diskAddress  []int
	pageFrequent []int
	inPageTable  []bool
}
