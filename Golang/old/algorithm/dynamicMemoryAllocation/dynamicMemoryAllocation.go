package main

// package dynamicMemoryAllocation

import (
	"errors"
	"fmt"
)

type MemoryTable struct {
	Root           *MemoryNode
	AllMemorySize  int
	FreeMemorySize int
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
		Root:           root,
		AllMemorySize:  memorySize,
		FreeMemorySize: memorySize,
	}, nil
}

func NewMemoryNode(allocateMemory int) (*MemoryNode, error) {
	if allocateMemory <= 0 {
		// default allocateMemory > 0
		return nil, errors.New("allocateMemory is invaild!")
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

func (mt *MemoryTable) FindMemoryBlockByStartAddree(startAddress int) *MemoryNode {
	tmp := mt.Root.Next
	for tmp != nil {
		if tmp.StartAddress == startAddress {
			return tmp
		}
		tmp = tmp.Next
	}
	return nil
}

func (mt *MemoryTable) printMemoryTable() {
	tmp := mt.Root.Next
	fmt.Println("IsUsed, StartAddress, Size")
	for tmp != nil {
		fmt.Println(tmp.IsUsed, tmp.StartAddress, tmp.Size)
		tmp = tmp.Next
	}
	return
}

func (mt *MemoryTable) AllocateArrayMemory(allocateMemoryArray []int) error {
	sum := 0
	for _, v := range allocateMemoryArray {
		sum += v
	}
	if sum > mt.FreeMemorySize {
		return errors.New("memory error, try again...")
	}
	// TODO
	for _, v := range allocateMemoryArray {
		if err := mt.Allocate(v); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (mt *MemoryTable) Allocate(allocateMemory int) error {
	// Default first...
	if allocateMemory <= 0 {
		// default allocateMemory > 0
		return errors.New("allocateMemory is invaild!")
	}
	// RootNode -> dummy
	tmp := mt.Root.Next

	for tmp != nil {
		if !tmp.IsUsed && tmp.Size > allocateMemory {
			// insert head
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
				[0, 500) allocated
				[500, 600) allocated
				[600, 800) allocated
				[800, 2500) free
			**/

			tmp.StartAddress += allocateMemory
			tmp.Size -= allocateMemory
			mt.FreeMemorySize -= allocateMemory
			return nil
		} else if !tmp.IsUsed && tmp.Size == allocateMemory {
			tmp.IsUsed = true
			mt.FreeMemorySize -= allocateMemory
			return nil
		}
		tmp = tmp.Next
	}
	// TODO call Recovery func
	if err := mt.RecoveryBySize(allocateMemory, 1); err != nil {
		return errors.New("MemoryError")
	}
	return mt.Allocate(allocateMemory)
}

// TODO
// func (mt *MemoryTable) ChoiceAllocate() {
// }

func (mt *MemoryTable) RecoveryBySize(memorySize int, number int) error {
	// number: int Recovery numver (default: -1 complete recovery)
	// memorySize: int Recovery target
	// number == -1 recovery all memory
	// number > 0 recovery number memory

	if memorySize == 0 || number == 0 {
		return nil
	}

	//dummyNode...
	tmp := mt.Root.Next
	if tmp == nil {
		return nil
	}
	for tmp.Next != nil {
		if tmp.Size == memorySize && tmp.IsUsed {
			if err := mt.JudgeRecoveryMode(tmp); err != nil {
				fmt.Println(err)
				return err
			}
			number--
			if number == 0 || tmp == nil || tmp.Next == nil {
				return nil
			}
		}
		tmp = tmp.Next
	}
	return nil
}

func (mt *MemoryTable) JudgeRecoveryMode(node *MemoryNode) error {
	/*
		1. 上下都已分配 res.Front.IsUsed && res.Next.IsUsed (true)
		2. 上分配, 下未分配 res.Front.IsUsed && !res.Next.IsUsed (true)
		3. 下分配, 上未分配 !res.Front.IsUsed && res.Next.IsUsed (true)
		4. 上下均未分配 !res.Front.IsUsed && !res.Next.IsUsed (true)
	**/
	res := node
	if res != nil {
		if res.Front.IsUsed && res.Next.IsUsed {
			// 上下都已分配 res.Front.IsUsed && res.Next.IsUsed (true)

			res.IsUsed = false
			mt.FreeMemorySize += res.Size
		} else if res.Front.IsUsed && !res.Next.IsUsed {
			// 上分配, 下未分配 res.Front.IsUsed && !res.Next.IsUsed (true)
			res.IsUsed = false
			appendMemory := res.Next
			nextNode := appendMemory.Next
			res.Size += appendMemory.Size
			res.Next = appendMemory.Next
			if nextNode == nil {
				res.Next = nil
			} else {
				nextNode.Front = res
			}
			appendMemory = nil
			mt.FreeMemorySize += res.Size
		} else if !res.Front.IsUsed && res.Next.IsUsed {
			// 下分配, 上未分配 !res.Front.IsUsed && res.Next.IsUsed (true)
			frontNode := res.Front
			frontNode.Size += res.Size
			nextNode := res.Next
			frontNode.Next = nextNode
			if nextNode != nil {
				nextNode.Front = frontNode
			}
			mt.FreeMemorySize += frontNode.Size
			res = nil
		} else {
			// 上下均未分配 !res.Front.IsUsed && !res.Next.IsUsed (true)
			frontRecoveryNode := res.Front
			nextRecoveryNode := res.Next
			frontRecoveryNode.Size += res.Size
			if nextRecoveryNode != nil {
				frontRecoveryNode.Size += nextRecoveryNode.Size
				frontRecoveryNode.Next = nextRecoveryNode.Next
			}
			nextRecoveryNode = nil
			res = nil
			mt.FreeMemorySize += frontRecoveryNode.Size
		}
	} else {
		errors.New("not found!")
	}
	return nil
}

func (mt *MemoryTable) RecoveryByAddress(startAddress int) error {
	res := mt.FindMemoryBlockByStartAddree(startAddress)
	return mt.JudgeRecoveryMode(res)
}

func main() {
	mt, err := NewMemoryTable(2000)
	if err != nil {
		fmt.Println(err)
		return
	}
	allocateArray := []int{500, 300, 300, 300, 100}
	if err := mt.AllocateArrayMemory(allocateArray); err != nil {
		fmt.Println(err)
		return
	}

	mt.printMemoryTable()
	mt.RecoveryByAddress(500)
	fmt.Println()
	mt.printMemoryTable()
	// fmt.Println()
	// mt.RecoveryBySize(300, -1)
	// mt.printMemoryTable()
	// fmt.Println()
	// mt.RecoveryBySize(100, -1)
	// mt.printMemoryTable()
	// fmt.Println()
}
