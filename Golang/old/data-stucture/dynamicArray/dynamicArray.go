// package dynamicArray
package main

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	dynamicArray []int
	size         int
	capacity     int
	arrayType    string
}

func main() {
	a := DynamicArray{}
	a.Init(10, "int")
	a.PrintInfo()
	for i := 0; i < 9; i++ {
		a.Insert(10, i)
	}
	a.PrintInfo()
	a.PrintValue()
	for i := 8; i > 0; i-- {
		a.Delete(1)
	}
	a.PrintInfo()

}

func (da *DynamicArray) Init(capacity int, arrayType string) error {
	if capacity > 0 && arrayType != "" {
		da.size = 0
		da.capacity = capacity
		da.dynamicArray = make([]int, capacity)
		da.arrayType = arrayType
		return nil
	}
	return errors.New("siteValue is invaild")
}

func (da *DynamicArray) Insert(value int, site int) error {
	// if use over 80% memory space, will call dynamicUp func
	if site < 0 {
		return errors.New("siteValue is invaild")
	}
	if da.size >= int(0.8*float64(da.capacity)) {
		da.dynamicUp()
	}
	da.dynamicArray[site] = value
	da.size++
	return nil
}

func (da *DynamicArray) Delete(site int) error {
	// if use lower 10% memory space, will call dynamicDown func
	if da.size <= int(0.1*float64(da.capacity)) {
		da.dynamicDown()
	}
	da.dynamicArray[site], da.dynamicArray[len(da.dynamicArray)-1] = da.dynamicArray[len(da.dynamicArray)-1], da.dynamicArray[site]
	da.dynamicArray = da.dynamicArray[:len(da.dynamicArray)-1]
	da.size--
	return nil

}

func (da *DynamicArray) dynamicUp() {
	if da.size >= int(0.8*float64(da.capacity)) {
		tmpArray := make([]int, da.capacity*2)
		for i, v := range da.dynamicArray {
			tmpArray[i] = v
		}
		da.dynamicArray = tmpArray
		da.capacity = 2 * da.capacity
	}
}

func (da *DynamicArray) dynamicDown() {
	if da.size <= int(0.1*float64(da.capacity)) {
		tmpArray := make([]int, da.capacity/2)
		for i, v := range da.dynamicArray {
			if i > len(tmpArray)-1 {
				break
			}
			tmpArray[i] = v
		}
		da.dynamicArray = tmpArray
		da.capacity = da.capacity / 2
	}
}

func (da *DynamicArray) Find(value int) bool {
	for _, v := range da.dynamicArray {
		if v == value {
			return true
		}
	}
	return false
}

func (da *DynamicArray) Update(oldValue, newValue int) {
	for i, v := range da.dynamicArray {
		if v == oldValue {
			da.dynamicArray[i] = newValue
		}
	}
	return
}

func (da *DynamicArray) PrintInfo() {
	fmt.Printf("Now array size: %d\n", da.size)
	fmt.Printf("Now array capacity: %d\n", da.capacity)
	return
}

func (da *DynamicArray) PrintValue() {
	for i, v := range da.dynamicArray {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	fmt.Println()
}
