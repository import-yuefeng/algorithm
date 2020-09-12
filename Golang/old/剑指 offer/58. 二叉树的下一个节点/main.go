package main

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
	Parent      *BinaryTreeNode
}

func main() {

}

func GetNextNode(root, cur *BinaryTreeNode) *BinaryTreeNode {
	if cur.Right != nil {
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		return tmp
	} else if cur.Parent != nil && cur.Parent.Left == cur {
		return cur.Parent
	} else if cur.Parent != nil && cur.Parent.Right == cur {
		tmp := cur.Parent
		for tmp.Parent != nil && tmp.Parent.Left != tmp {
			tmp = tmp.Parent
		}
	}
}
