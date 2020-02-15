package main

import (
	"errors"
	"fmt"
)

type LinkListNode struct {
	Val  int
	Next *LinkListNode
}

func main() {
	root := new(LinkListNode)
	cur := root
	// front := cur
	for i := 0; i < 10; i++ {
		cur.Next = new(LinkListNode)
		cur.Val = i
		// front = cur
		cur = cur.Next
	}
	// front.Next = nil
	cur.Next = root.Next.Next
	res := isCycleLinkedlist(root)
	fmt.Println(res)
	if res {
		r := getCycleLinkedlistEntrance(root)
		fmt.Println(r.Val)
	}
	// res, err := kNumInLinkedlist(root, 100)
	// if err == nil {
	// 	fmt.Println(res.Val)
	// }
}

func kNumInLinkedlist(root *LinkListNode, k int) (res *LinkListNode, err error) {
	if root == nil {
		return nil, errors.New("invaild")
	}
	fast := root
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	if k > 0 || k < 0 {
		return nil, errors.New("invaild")
	}
	slow := root
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow, nil
}

func isCycleLinkedlist(root *LinkListNode) bool {
	if root == nil {
		return false
	}
	slow, fast := root, root
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func getCycleLinkedlistEntrance(root *LinkListNode) *LinkListNode {
	if root == nil {
		return nil
	}
	slow, fast := root, root
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
		if slow == fast {
			fast = slow
			slow = root
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
