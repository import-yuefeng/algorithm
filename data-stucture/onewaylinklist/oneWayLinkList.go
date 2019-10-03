package main

import "fmt"

//one-way link-list
// golang

type node struct {
	data int
	next *node
	size int
	use  int
}

func (head *node) NewLink(size int) *node {
	head.next = nil
	head.size = size
	return head
}

func (head *node) IsLoop() bool {
	slow, fast := head, head
	if head.use <= 1 {
		return false
	}
	slow = slow.next
	fast = fast.next.next
	for fast != nil && slow != nil {
		slow = slow.next
		fast = fast.next
		if fast == slow {
			return false
		}
		if fast != nil {
			fast = fast.next
		}
	}
	return true
}

func (head *node) Reverse() (bool, *node) {
	if head.use < 1 {
		return true, head
	}
	front := head
	handPoint := front.next
	tail := handPoint.next
	newHead := new(node)
	for {
		if front != head {
			handPoint.next = front
		} else {
			handPoint.next = nil
		}
		front = handPoint
		if tail == nil {
			newHead.next = handPoint
			newHead.use = head.use
			newHead.size = head.size
			break
		}
		handPoint = tail
		tail = tail.next
	}
	head.next = nil
	// import!!!!!!!!
	return true, newHead
}

func (head *node) FindAll(data int) int {
	count, tmp := 0, head.next
	for {
		if tmp == nil {
			return count
		}
		if tmp.data == data {
			count++
		}
		tmp = tmp.next
	}
}

func (head *node) Insert(data int, seat int) bool {
	if seat > head.size || seat > head.use || seat < 0 {
		return false
	}
	count, tmp := 1, head
	for {
		if count == seat {
			newNode := new(node)
			newNode.data = data
			newNode.next = tmp.next
			tmp.next = newNode
			head.use++
			break
		}
		count++
		tmp = tmp.next
	}
	return true
}

func (head *node) InsertHead(data int) bool {
	if head.use+1 <= head.size {
		tmp := new(node)
		tmp.next = head.next
		head.next = tmp
		tmp.data = data
		head.use++
	} else {
		return false
	}
	return true
}

func (head *node) InsertTail(data int) bool {
	if head.use+1 <= head.size {
		tmp := new(node)
		tmp = head
		newNode := new(node)
		newNode.data, newNode.next = data, nil
		for {
			tmp = tmp.next
			if tmp.next == nil {
				tmp.next = newNode
				head.use++
				break
			}
		}
	} else {
		return false
	}
	return true
}

func (head *node) DeleteHead() bool {
	// tmp := head.next
	head.next = head.next.next
	head.use--
	// free(tmp)
	return true
}

func (head *node) DeleteTail() bool {
	// two point
	tmp := head
	tmp2 := tmp.next
	for {
		if tmp2.next == nil {
			tmp.next = nil
			head.use--
			// free(tmp2)
			break
		}
	}
	return true
}

func (head *node) FindElem(data int) (exist bool, seat int) {
	tmp := head.next
	count := 1
	for {
		if tmp == nil {
			return false, -1
		}
		if tmp.data == data {
			return true, count
		}
		tmp = tmp.next
		count++
	}
}

func (head *node) PrintLinkList() {
	tmp := head
	for {
		fmt.Print(tmp.data)
		if tmp.next != nil {
			fmt.Print("->")
		}
		tmp = tmp.next
		if tmp == nil {
			break
		}
	}
	fmt.Printf("\n")
}

func main() {
	head := new(node)
	head.NewLink(10)
	fmt.Print("New link list\n")
	head.PrintLinkList()
	head.InsertHead(1)
	fmt.Print("Insert head 1\n")
	head.PrintLinkList()
	head.InsertHead(1)
	fmt.Print("Insert head 1\n")
	head.PrintLinkList()
	head.InsertHead(1)
	fmt.Print("Insert head 1\n")
	head.Insert(2, 2)
	fmt.Print("Insert 2 seat: 2\n")
	head.InsertTail(5)
	fmt.Print("Insert tail 5\n")
	head.PrintLinkList()
	if result, seat := head.FindElem(2); result {
		fmt.Printf("%v seat: %d\n", result, seat)
	} else {
		fmt.Print("Not found elem\n")
	}
	all := head.FindAll(0)
	fmt.Printf("linklist has: 1(%d)\n", all)
	_, newHead := head.Reverse()

	newHead.PrintLinkList()

}
