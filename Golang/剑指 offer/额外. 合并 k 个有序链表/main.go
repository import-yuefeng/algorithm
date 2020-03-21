package main

import "fmt"

type Linkedlist struct {
	Val  int
	Next *Linkedlist
}

func main() {
	test := make([]*Linkedlist, 5)
	a := &Linkedlist{Val: 1, Next: &Linkedlist{Val: 2, Next: &Linkedlist{Val: 4, Next: &Linkedlist{Val: 6}}}}
	b := &Linkedlist{Val: 8, Next: &Linkedlist{Val: 9, Next: &Linkedlist{Val: 10, Next: &Linkedlist{Val: 11}}}}
	c := &Linkedlist{Val: 13, Next: &Linkedlist{Val: 16, Next: &Linkedlist{Val: 19, Next: &Linkedlist{Val: 20}}}}
	d := &Linkedlist{Val: 33, Next: &Linkedlist{Val: 35, Next: &Linkedlist{Val: 38, Next: &Linkedlist{Val: 40}}}}
	e := &Linkedlist{Val: 50, Next: &Linkedlist{Val: 56, Next: &Linkedlist{Val: 58, Next: &Linkedlist{Val: 80}}}}
	test[0] = a
	test[1] = b
	test[2] = d
	test[3] = e
	test[4] = c
	res := mergeSort(test, 0, len(test)-1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func mergeSort(list []*Linkedlist, left, right int) *Linkedlist {
	if left < right {
		mid := left + ((right - left) >> 1)
		left := mergeSort(list, left, mid)
		right := mergeSort(list, mid+1, right)
		res := merge(left, right)
		return res
	}
	return list[left]
}

func merge(left, right *Linkedlist) *Linkedlist {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	if left.Val < right.Val {
		left.Next = merge(left.Next, right)
		return left
	}
	right.Next = merge(left, right.Next)
	return right
}
