package main

type BST struct {
	Val         int
	Left, Right *BST
}

func main() {

}

func BST2Linkedlist(root *BST, head, end *BST) *BST {
	if root == nil || root.Next = nil {
		return root
	}
	BST2Linkedlist(root.Left, head, end)
	if end == nil {
		end = root
		head = root
	} else {
		end.Right = root
		// 链表当前最后一个元素 的右节点连接新的节点
		root.Left = end
		end = root
	}
	return head
}
