package main

import "fmt"

// "fmt"
// "strconv"

func main() {
	root := new(BST)
	root.Val = 10
	root.Left = &BST{Val: 6, Left: &BST{Val: 3}, Right: &BST{Val: 9}}
	root.Right = &BST{Val: 15, Left: &BST{Val: 12}, Right: &BST{Val: 19}}

	var head, end *BST
	head, end = new(BST), new(BST)
	head = findLastLeftNode(root)

	BST2Linkedlist(root, end)
	// res, _ := strconv.Atoi("12")
	// fmt.Println(res, strconv.Itoa(12))
	for head != nil && head.Right != nil {
		fmt.Println(head.Val)
		head = head.Right
	}
}

func findLastLeftNode(root *BST) *BST {
	if root == nil || root.Left == nil {
		return root
	}
	return findLastLeftNode(root.Left)
}

func BST2Linkedlist(root *BST, end *BST) {
	if root == nil {
		return
	}
	BST2Linkedlist(root.Left, end)

	if end == nil {
		end = root
		// fmt.Println(root.Val)
	} else {
		root.Left = end
		end.Right = root
		// 链表当前最后一个元素 的右节点连接新的节点
		end = root
	}
	BST2Linkedlist(root.Right, end)
}

type BST struct {
	Val         int
	Left, Right *BST
}

func (b *BST) Insert(val int) {

}

func (b *BST) Delete(val int) *BST {

}

func (b *BST) Search(val int) (res *BST, err error) {

}
