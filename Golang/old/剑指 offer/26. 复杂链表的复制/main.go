package main

type ComplexLinkedlist struct {
	Val    int
	Next   *ComplexLinkedlist
	Random *ComplexLinkedlist
}

func main() {

}

func CloneLinkedlist(root *ComplexLinkedlist) *ComplexLinkedlist {
	if root == nil {
		return nil
	}
	nodeMap := make(map[*ComplexLinkedlist]*ComplexLinkedlist)
	cur := root
	res := new(ComplexLinkedlist)
	tmp := res
	for cur != nil {
		tmp.Val = cur.Val
		tmp.Next = new(ComplexLinkedlist)
		nodeMap[cur] = tmp
		tmp = tmp.Next
		cur = cur.Next
	}
	cur = root
	tmp = res
	for cur != nil {
		if cur.Random != nil {
			tmp.Random = nodeMap[cur.Random]
		}
		cur = cur.Next
		tmp = tmp.Next
	}
	return res
}

func CloneNode(head *ComplexLinkedlist) {
	cur := root
	for cur != nil {
		node := new(ComplexLinkedlist)
		node.Val = cur.Val
		node.Next = cur.Next
		cur.Next = node
		cur = node.Next
	}
	return
}

func ConnectRandomNode(head *ComplexLinkedlist) {
	cur := root
	t := root.Next
	for cur != nil && t != nil {
		t.Random = cur.Random.Next
		cur = t.Next
		t = cur.Next
	}
	return
}

func ReconnectLinkedlist(head *ComplexLinkedlist) *ComplexLinkedlist {
	res := head.Next
	i := res
	for i != nil {
		t := i.Next
		if t == nil {
			return res
		}
		t = t.Next
		i.Next = t
		i = t
	}
	return res
}
