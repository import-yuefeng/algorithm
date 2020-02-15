package main

import "fmt"

type LinkedlistNode struct {
	Val  int
	Next *LinkedlistNode
}

func main() {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := new(LinkedlistNode)
	root.Val = l[0]
	cur := root
	for i := 1; i < len(l); i++ {
		cur.Next = NewLinkedlistNode(l[i])
		cur = cur.Next
	}
	PrintLinkedlist(root)
	res := ReverseLinkedlsitNodeByIterate(root)
	res = ReverseLinkedlistNodeByRecursive(res)
	PrintLinkedlist(res)
}

func NewLinkedlistNode(x int) *LinkedlistNode {
	return &LinkedlistNode{
		Val:  x,
		Next: nil,
	}
}

func ReverseLinkedlsitNodeByIterate(root *LinkedlistNode) *LinkedlistNode {
	if root == nil || root.Next == nil {
		return root
	}
	front, cur, next := root, root.Next, root.Next.Next
	for cur != nil {
		cur.Next = front
		front = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}
	root.Next = nil
	return front
}

func ReverseLinkedlistNodeByRecursive(cur *LinkedlistNode) *LinkedlistNode {

	if cur == nil || cur.Next == nil {
		return cur
	}
	r := ReverseLinkedlistNodeByRecursive(cur.Next)
	cur.Next.Next = cur
	// 下一个节点的指针指向自己
	cur.Next = nil
	// 将自己指向空
	return r
}

func PrintLinkedlist(root *LinkedlistNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}
