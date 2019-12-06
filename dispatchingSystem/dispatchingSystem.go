package main

// package dispatchingSystem

import (
	"errors"
	"fmt"
)

type JCB struct {
	/*
		JCB 结构体定义
	**/
	JID         int
	StartTime   float64
	RequestTime float64
	ArriveTime  float64
	FinaishTime float64
	WaitTime    float64
	RevolveTime float64
}

type PCBList struct {
	/*
		PCB任务队列 结构体
	**/
	job []*JCB
}

type MinHeap struct {
	/*
		最小堆 结构体定义
	**/
	job      []*JCB
	size     int
	capacity int
}

func NewMinHeap(capacity int) *MinHeap {
	/*
		初始化堆
	**/
	root := new(MinHeap)
	root.job = make([]*JCB, capacity)
	root.capacity = capacity
	root.size = 0
	return root
}

func (h *MinHeap) BuildMinHeap() {
	/*
		构造最小堆
	**/
	if h.size == 0 {
		return
	}
	for parent := h.size / 2; parent >= 0; parent-- {
		h.downHeapNode(parent)
	}
}

func (h *MinHeap) IsFull() bool {
	return h.size == h.capacity
}

func (h *MinHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MinHeap) Insert(val *JCB) error {
	/*
		向堆内插入一个 新任务并动态调整
	**/
	if h.IsFull() {
		return errors.New("minHeap is full!")
	}

	h.job[h.size] = val
	h.size++
	h.upHeapNode(h.size - 1)
	return nil
}

func (h *MinHeap) Delete() (val *JCB, err error) {
	/*
		从堆中删除一个元素(根节点)
	**/
	if h.IsEmpty() {
		return nil, errors.New("heap is empty!")
	}
	val = h.job[0]
	err = nil
	h.size--
	h.job[0] = h.job[h.size]
	if !h.IsEmpty() {
		h.downHeapNode(0)
	}
	return
}

func (h *MinHeap) downHeapNode(i int) {
	/*
		堆中进行建堆和删除元素进行的 堆动态下沉操作
	**/

	parent := i
	child := parent*2 + 1

	for child < h.size {
		if child+1 < h.size && h.job[child+1].RequestTime < h.job[child].RequestTime {
			child++
		}
		if h.job[child].RequestTime >= h.job[parent].RequestTime {
			break
		}
		h.job[child], h.job[parent] = h.job[parent], h.job[child]
		parent = child
		child = parent*2 + 1
	}
}

func (h *MinHeap) upHeapNode(i int) {
	/*
		堆中插入新元素时进行的堆动态上浮操作
	**/

	child := i
	parent := (child - 1) >> 1
	for child > 0 {
		if child+1 < h.size && h.job[child+1].RequestTime < h.job[child].RequestTime {
			child++
		}
		if h.job[parent].RequestTime <= h.job[child].RequestTime {
			break
		}
		h.job[parent], h.job[child] = h.job[child], h.job[parent]
		child = parent
		parent = (child - 1) >> 1
	}
}

func (h *MinHeap) updataWaitTime(waitTime float64) {
	for _, v := range h.job {
		v.WaitTime += waitTime
	}
	return
}

func qSort(jcb []*JCB, left, right int) {
	/*
		根据任务到达时间进行的快速排序算法
	**/

	if left < right {
		l, r := left, right
		middle := jcb[(l+r)>>1]
		for l < r {
			for l < r && middle.ArriveTime > jcb[l].ArriveTime {
				l++
			}
			for l < r && middle.ArriveTime < jcb[r].ArriveTime {
				r--
			}
			if l >= r {
				break
			}
			if jcb[r].ArriveTime == jcb[l].ArriveTime && middle.ArriveTime == jcb[r].ArriveTime {
				r--
			} else {
				jcb[r], jcb[l] = jcb[l], jcb[r]
			}
		}
		qSort(jcb, left, l-1)
		qSort(jcb, r+1, right)
	}
	return
}

func PrintJob(data *JCB) {
	/*
		打印具体任务运行情况 函数
	**/

	data.FinaishTime = data.StartTime + data.RequestTime
	data.WaitTime = data.StartTime - data.ArriveTime
	data.RevolveTime = data.FinaishTime - data.ArriveTime

	fmt.Printf(
		"JID: %d, ArriveTime: %.1f, StartTime: %.1f, FinaishTime: %.1f, RevolveTime: %.1f, WaitTime: %.1f, RequestTime: %.1f \n",
		data.JID, data.ArriveTime, data.StartTime, data.FinaishTime, data.RevolveTime, data.WaitTime, data.RequestTime,
	)
	return
}

type Stack struct {
	/*
		自己实现的简单栈
		注意 data 的类型是 interface{}
		可以存储任何类型数据, 在使用时需要进行类型转换(interface-> spacial type)
	**/

	top      int
	size     int
	data     []interface{}
	capacity int
}

func (s *Stack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	s.data = make([]interface{}, capacity)
	s.capacity = capacity
	s.size = 0
	s.top = -1
	return nil
}

func (s *Stack) Push(val interface{}) error {
	if s.IsFull() {
		return errors.New("stack is full!")
	}
	s.top++
	s.data[s.top] = val
	s.size++
	return nil
}

func (s *Stack) Pop() (val interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty!")
	}
	val = s.data[s.top]
	s.top--
	s.size--
	err = nil
	return
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == s.capacity
}

func ScanJob() *PCBList {
	/*
		输入任务数据
		并初始化任务数组
	**/
	var (
		jobNumber   int
		RequestTime float64
		ArriveTime  float64
	)
	fmt.Printf("Please input your job number: ")
	fmt.Scanln(&jobNumber)
	jobList := new(PCBList)

	for i := 0; i < jobNumber; i++ {
		newJob := new(JCB)
		newJob.JID = i
		fmt.Printf("Please input id.%d job request, arrive time: \n", i)
		fmt.Scanln(&RequestTime, &ArriveTime)
		newJob.StartTime = 0
		newJob.RequestTime = RequestTime
		newJob.ArriveTime = ArriveTime
		newJob.FinaishTime = 0
		newJob.WaitTime = 0
		newJob.RevolveTime = 0

		jobList.job = append(jobList.job, newJob)
	}
	return jobList
}

func SchedulingChoice(flag int, jobList *PCBList, nowTime float64) {
	/*
		flag    记录选择的调度模式进行进程模拟调度
		jobList 存储已经进入可调度状态的进程组(由进程数组构成)
		nowTime 用于记录任务运行所花费的时间
	**/
	taskStack := new(Stack)
	// init task stack
	if err := taskStack.Init(len(jobList.job)); err != nil {
		fmt.Println(err)
		return
	}
	heap := NewMinHeap(len(jobList.job))
	// init task heap
	heap.BuildMinHeap()

	switch flag {

	case 1:
		// FIFO 先来先服务模式
		// 任务根据到达时间进行快速排序, 时间复杂度: O(nlogn), 升序排序
		// 空间复杂度 O(1), 除临时变量外, 未利用额外空间
		// 再按顺序打印即可
		qSort(jobList.job, 0, len(jobList.job)-1)
		for _, v := range jobList.job {
			v.StartTime = nowTime
			PrintJob(v)
			nowTime += v.RequestTime
		}
	case 2:
		// SJF 短任务优先 调度模式
		// 先根据任务到达时间进行一次快速排序, 同 FIFO
		// 再进行入栈操作, 相同到达时间的任务, 均入栈, 直到碰到下一个不是相同时间到达的任务
		// 开始出栈. 并开始入最小堆, 当栈内所有任务均出栈完成, 开始出堆元素, 按出堆顺序打印即可
		// 其中 堆按照 任务所需运行时间动态构造最小堆
		// 时间复杂度最高在 1.建堆, O(nlogn), 2. 快排 O(nlogn)
		// 故时间复杂度是 O(nlogn), 空间复杂度是 O(n+n) -> O(2n) 分别为堆, 栈所占用的额外空间
		qSort(jobList.job, 0, len(jobList.job)-1)
		// sort task by ArriveTime
		jobList.job = append(jobList.job, &JCB{})
		// append flag
		for _, v := range jobList.job {
			// push value to stack that stack is empty
			if taskStack.IsEmpty() {
				if err := taskStack.Push(v); err != nil {
					fmt.Println(err)
					return
				}
			} else {
				if taskStack.Top().(*JCB).ArriveTime == v.ArriveTime {
					if err := taskStack.Push(v); err != nil {
						fmt.Println(err)
						return
					}
				} else {
					for !taskStack.IsEmpty() {
						if currentTask, err := taskStack.Pop(); err != nil {
							fmt.Println(err)
							return
						} else {
							heap.Insert(currentTask.(*JCB))
						}
					}
					if err := taskStack.Push(v); err != nil {
						fmt.Println(err)
						return
					}
					for !heap.IsEmpty() {
						currentTask, err := heap.Delete()
						if err != nil {
							fmt.Println(err)
							return
						}
						currentTask.StartTime = nowTime
						nowTime += currentTask.RequestTime
						PrintJob(currentTask)
					}
				}
			}
		}
	case 3:
		// TODO add HRN func

	default:
		fmt.Println("nil choice!")
	}
}

func main() {

	var (
		flag    int
		nowTime float64
		jobList *PCBList
	)

	jobList = ScanJob()

	fmt.Println("Choice your function: \n 1. FCFS \n 2. SJF \n 3. HRN")
	fmt.Scanln(&flag)

	SchedulingChoice(flag, jobList, nowTime)

	return
}
