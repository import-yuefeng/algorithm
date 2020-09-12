package main

import "fmt"

type LinkListNode struct {
	Val  int
	Next *LinkListNode
}

func main() {
	list1 := new(LinkListNode)
	list2 := new(LinkListNode)
	l1 := []int{1, 3, 5, 7, 9, 11, 12}
	l2 := []int{2, 4, 5, 6, 7, 8}
	cur := list1
	front := cur
	for _, v := range l1 {
		cur.Val = v
		cur.Next = new(LinkListNode)
		front = cur
		cur = cur.Next
	}
	front.Next = nil
	cur = list2
	for _, v := range l2 {
		cur.Val = v
		cur.Next = new(LinkListNode)
		front = cur
		cur = cur.Next
	}
	front.Next = nil

	// res := MergeLinkList(list1, list2)
	res := MergeLinkList_2(list1, list2)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}

func MergeLinkList(a, b *LinkListNode) *LinkListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	root := new(LinkListNode)
	if a.Val < b.Val {
		root.Val = a.Val
		root.Next = MergeLinkList(a.Next, b)
	} else {
		root.Val = b.Val
		root.Next = MergeLinkList(a, b.Next)
	}
	return root
}

func MergeLinkList_2(a, b *LinkListNode) *LinkListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	root := new(LinkListNode)
	cur := root
	front := cur
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Val = a.Val
			a = a.Next
		} else {
			cur.Val = b.Val
			b = b.Next
		}
		front = cur
		cur.Next = new(LinkListNode)
		cur = cur.Next
	}

	if a != nil {
		front.Next = a
	}
	if b != nil {
		front.Next = b
	}
	return root
}
