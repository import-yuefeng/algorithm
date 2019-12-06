package main

// package dynamicMemoryAllocation

import (
	"errors"
	"fmt"
)

type MemoryTable struct {
	Root          *MemoryNode
	AllMemorySize int
}

type MemoryNode struct {
	StartAddress int
	Size         int
	IsUsed       bool
	Next, Front  *MemoryNode
}

func NewMemoryTable(memorySize int) (MemoryTable, error) {
	if memorySize <= 0 {
		return MemoryTable{}, errors.New("memorySize is invaild!")
	}

	newMemoryBlock := &MemoryNode{
		StartAddress: 0,
		Size:         memorySize,
		IsUsed:       false,
		Next:         nil,
		Front:        nil,
	}
	// dummy node...
	root := &MemoryNode{
		StartAddress: 0,
		Size:         0,
		IsUsed:       false,
		Next:         newMemoryBlock,
		Front:        nil,
	}
	newMemoryBlock.Front = root

	return MemoryTable{
		Root:          root,
		AllMemorySize: memorySize,
	}, nil
}

func NewMemoryNode(allocateMemory int) (*MemoryNode, error) {
	if allocateMemory <= 0 {
		return new(MemoryNode), errors.New("allocateMemory is invaild!")
	}
	newMemoryBlock := MemoryNode{
		StartAddress: 0,
		Size:         allocateMemory,
		IsUsed:       true,
		Next:         nil,
		Front:        nil,
	}
	return &newMemoryBlock, nil
}

// func DeleteMemoryBlock() error {}

func FindMemoryBlockByStartAddree(startAddress int) *MemoryNode {
	// TODO
	return nil
}

func (mt *MemoryTable) Allocate(allocateMemory int) error {
	// Default first...
	if allocateMemory <= 0 {
		return errors.New("allocateMemory is invaild!")
	}
	tmp := mt.Root.Next
	// dummy.Next
	for tmp != nil {
		if !tmp.IsUsed && tmp.Size > allocateMemory {
			newMemoryBlock, err := NewMemoryNode(allocateMemory)
			if err != nil {
				fmt.Println(err)
				return err
			}
			newMemoryBlock.StartAddress = tmp.StartAddress

			frontMemoryBlock := tmp.Front

			frontMemoryBlock.Next = newMemoryBlock
			newMemoryBlock.Front = frontMemoryBlock
			newMemoryBlock.Next = tmp
			tmp.Front = newMemoryBlock
			/*
				mock data:
				[0, 500] allocated
				[501, 600] allocated
				[601, 800] allocated
				[801, 2500] free
			**/

			tmp.StartAddress += allocateMemory + 1
			return nil
		} else if !tmp.IsUsed && tmp.Size == allocateMemory {
			tmp.IsUsed = true
			return nil
		}
	}
	// TODO call Recovery func
	return errors.New("MemoryError")
}

// TODO
// func (mt *MemoryTable) ChoiceAllocate() {
// }

func (mt *MemoryTable) RecoveryBySize(memorySize int, number int) error {
	// TODO
	if memorySize == 0 || number == 0 {
		return nil
	}
	// number == -1 recovery all memory
	// number > 0 recovery number memory

	//dummy Node...
	tmp := mt.Root.Next
	for tmp.Next != nil {

	}
	return nil
}

func (mt *MemoryTable) RecoveryByAddress(startAddress int) error {
	res := FindMemoryBlockByStartAddree(startAddress)
	if res != nil {
		/*
			1. 上下都已分配 res.Front.IsUsed && res.Next.IsUsed (true)
			2. 上分配, 下未分配 res.Front.IsUsed && !res.Next.IsUsed (true)
			3. 下分配, 上未分配 !res.Front.IsUsed && res.Next.IsUsed (true)
			4. 上下均未分配 !res.Front.IsUsed && !res.Next.IsUsed (true)
		**/

		if res.Front.IsUsed && res.Next.IsUsed {
			res.IsUsed = false
		} else if res.Front.IsUsed && !res.Next.IsUsed {
			appendMemory := res.Next
			res.Size += appendMemory.Size
			res.Next = appendMemory.Next

		} else if !res.Front.IsUsed && res.Next.IsUsed {

		} else {

		}

	} else {
		errors.New("not found!")
	}
	return nil
}

func main() {

}
