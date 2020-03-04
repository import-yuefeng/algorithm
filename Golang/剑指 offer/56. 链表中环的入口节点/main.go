package main

import "fmt"

type Linkedlist struct {
	Val  int
	Next *Linkedlist
}

func main() {
	root := new(Linkedlist)
	root.Val = 1
	root.Next = &Linkedlist{Val: 2, Next: &Linkedlist{Val: 3, Next: &Linkedlist{Val: 4, Next: &Linkedlist{Val: 5}}}}
	root.Next.Next.Next.Next = root.Next
	// PrintLinkedlist(root)
	res := findEntry(root)
	fmt.Println(res)
}

func findEntry(root *Linkedlist) *Linkedlist {
	if root == nil {
		return nil
	}
	fast, slow := root, root
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast.Next == nil || fast != slow {
		// not find ring
		return nil
	}

	p1 := root
	for p1 != slow {
		p1 = p1.Next
		slow = slow.Next
	}
	count := 1
	cur := fast.Next
	for cur != fast {
		cur = cur.Next
		count++
	}
	fmt.Println(count)
	return p1
}

func PrintLinkedlist(root *Linkedlist) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	return
}
