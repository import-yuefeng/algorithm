package main

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

func main() {

}

func isMirror(root *BinaryTreeNode) bool {
	return isMirrorUtil(root, root)
}

func isMirrorUtil(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isMirrorUtil(root1.Left, root2.Right) && isMirrorUtil(root1.Right, root2.Left)
}
