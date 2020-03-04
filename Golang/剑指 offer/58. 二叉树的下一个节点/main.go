package main

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
	Parent      *BinaryTreeNode
}

func main() {

}

func GetNextNode(root, cur *BinaryTreeNode) *BinaryTreeNode {
	if root == nil || cur == nil || cur.Parent == nil {
		return nil
	}
	var prev *BinaryTreeNode = nil
	if cur.Right != nil {
		res := cur.Right
		for res.Left != nil {
			res = res.Left
		}
		return res
	} else if cur.Parent.Left == cur {
		return cur.Parent
	} else if cur.Parent.Right == cur {
		prev = cur.Parent
		for prev != nil && cur != prev.Left {
			cur = prev
			prev = prev.Parent
		}
	}
	return prev
}
