package main

type BST struct {
	Val         int
	Left, Right *BST
}

func main() {
	root := new(BST)
	root.Val = 10
}

func BST2Linkedlist(root *BST, head, end *BST) *BST {
	if root == nil {
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
	BST2Linkedlist(root.Right, head, end)

	return head
}

func (b *BST) insert(x int) {
	node := new(BST)
	node.Val = x
	cur := b
	for cur != nil {
		if cur.Val > node.Val {
			if cur.Left == nil {
				cur.Left = node
				break
			} else {
				cur = cur.Left
			}
		} else {
			if cur.Right == nil {
				cur.Right = node
				break
			} else {
				cur = cur.Right
			}
		}
	}
}

func (b *BST) delete(x int) {
	cur := b
	prev := b
	for cur != nil {
		if cur.Val > x {
			prev = cur
			cur = cur.Left
		} else if cur.Val == x {
			break
		} else {
			prev = cur
			cur = cur.Right
		}
	}
	if cur == nil {
		// not found...
		return
	}
	if cur.Left == nil && cur.Right == nil {
		if prev.Left == cur {
			prev.Left = nil
			cur = nil
		} else if prev.Right == cur {
			prev.Right = nil
			cur = nil
		}
	} else if cur.Right != nil {
		r := cur.Right
		prev := cur
		for r.Left != nil {
			prev = r
			r = r.Left
		}
		cur.Val = r.Val
		prev.Left = nil
		r = nil
	} else if cur.Left != nil {
		l := cur.Left
		prev := cur
		for l.Right != nil {
			prev = l
			l = l.Right
		}
		cur.Val = l.Val
		prev.Right = nil
		l = nil
	}
	return
}
