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
