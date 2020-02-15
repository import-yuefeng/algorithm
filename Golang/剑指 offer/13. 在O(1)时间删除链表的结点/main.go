package main

type LinkListNode struct {
	Val  int
	Next *LinkListNode
}

func main() {

}

func deleteLinkedlistNode(root, deletedNode *LinkListNode) bool {
	if root == nil {
		return true
	}
	if root.Next == nil {
		deletedNode = nil
		root = nil
		return true
	} else if deletedNode.Next == nil {
		return traverseDelete(root, deletedNode)
	} else {
		deletedNode.Val = deletedNode.Next.Val
		deletedNode.Next = deletedNode.Next.Next
	}
	return true
}

func traverseDelete(root, deleteNode *LinkListNode) bool {
	front := root
	cur := root
	for cur != nil {
		if cur == deleteNode {
			front.Next = nil
		}
		front = cur
		cur = cur.Next
	}
	return true
}
