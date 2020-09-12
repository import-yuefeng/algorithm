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
	//res := ReverseLinkedlsitNodeByIterate(root)
	res := ReverseLinkedlistNodeByRecursive(root)
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
	var last *LinkedlistNode = ReverseLinkedlistNodeByRecursive(cur.Next)
	// 找到最后一个节点
	cur.Next.Next = cur
	// 将 head 的下一个节点的 next 指针指向自己, 用于反转自身
	cur.Next = nil
	// 将自身的 Next 转为空, 是为了防止出现死循环
	return last
}

func ReverseLinkedlistKtdNodeByRecursive(cur *LinkedlistNode, successor *LinkedlistNode, k int) *LinkedlistNode {
	if k == 1 {
		successor = cur.Next
		return cur
	}
	last := ReverseLinkedlistKtdNodeByRecursive(cur.Next, successor, k-1)
	cur.Next.Next = cur
	cur.Next = successor
	return last
}

//func ReverseLinkedlistM2NNodeByRecursive(cur *LinkedlistNode, m, n int) *LinkedlistNode {
//	if cur == nil {
//		return nil
//	}
//	if m == 1 {
//		return ReverseLinkedlistKtdNodeByRecursive(cur, m, n)
//	}
//	return ReverseLinkedlistM2NNodeByRecursive(cur.Next, m-1, n-1)
//}

func PrintLinkedlist(root *LinkedlistNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}
