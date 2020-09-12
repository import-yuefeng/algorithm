package queue

import "errors"

type Queue struct {
	front, rare int
	data        []interface{}
	size        int
	cap         int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		cap:  cap,
		size: 0,
		data: make([]interface{}, cap),
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == q.cap
}

func (q *Queue) Push(x interface{}) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.data[q.rare] = x
	q.rare = (q.rare+1) % q.cap
	q.size++
	return nil
}

func (q *Queue) Pop() (x interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	x = q.data[q.front]
	q.front = (q.front+1) % q.cap
	q.size--
	return
}