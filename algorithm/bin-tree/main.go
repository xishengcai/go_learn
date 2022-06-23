package main

import (
	"fmt"
	"math"
)

// 堆是一个完全二叉树，必须满足根节点大于两个叶子节点
// 从一个数组中构建堆：1.转二叉树 2.对半 3.siftDown

// ElementType 堆元素
type ElementType int

type Heap struct {
	List []ElementType
}

// Push 入列, 从堆尾插入，然后向上移动
func (h *Heap) Push(data ElementType) {
	j := len(h.List)
	h.List = append(h.List, data)
	k := (j - 1) / 2
	for k > -1 {
		if h.List[k] > h.List[j] {
			h.List[k], h.List[j] = h.List[j], h.List[k]
			j = k
			k = (j - 1) / 2
		} else {
			break
		}
	}
}

// Pop 出列
func (h *Heap) Pop() ElementType {
	minData := h.List[0]
	j := len(h.List)
	if j == 1 {
		h.List = []ElementType{}
		return minData
	}
	lastData := h.List[j-1]
	h.List = h.List[0 : j-1]
	h.ShiftDown(lastData, 0)
	return minData
}

// ShiftDown 向下移动
func (h *Heap) ShiftDown(data ElementType, begin int) {
	// 从堆顶构建最小堆, 在堆顶放入一个元素
	//循环条件 begin < len(heap)
	list := h.List
	if len(h.List) == 1 {
		h.List[0] = data
		return
	}
	for 2*begin+1 < len(h.List) {
		minIndex := 2*begin + 1
		if ((minIndex + 1) < len(list)) && (list[minIndex+1] < list[minIndex]) {
			minIndex += 1
		}
		if data < list[minIndex] {
			list[begin] = data
			break //当不发生上下交换的时候,说明已经满足堆这种数据结构的性质
		} else {
			list[begin] = list[minIndex]
			list[minIndex] = data
			begin = minIndex
		}
	}
}

func main() {
	// 无序数组
	list := []ElementType{9, 3, 12, 100, 1000, -11, -34}
	fmt.Printf("初始无序列表:%v\r\n", list)
	var heap Heap
	heap.List = list

	// 将一个无序数组构建成堆，假设后半部分的数据就是一个堆的叶子节点，然后在上面一层开始从左往右构建二叉树，
	// 这里的二叉树是一个抽象的感念，真实的情况是：这个堆仍然是个数组，它的的数据是分层存储的
	half := len(list) / 2
	for root := half; root >= 0; root-- {
		heap.ShiftDown(list[root], root)
	}

	fmt.Println("往堆中插入新的数据: -100, 30, -15, -95")
	heap.Push(-100)
	heap.Push(30)
	heap.Push(-15)
	heap.Push(-95)

	fmt.Println("打印堆")
	level := 1
	for i, data := range heap.List {
		maxIndex := 1*math.Pow(2, float64(level)) - 1
		if i+1 == int(maxIndex) {
			level += 1
			fmt.Printf("%v\r\n", data)
		} else {
			fmt.Printf("%v ", data)
		}
	}

	fmt.Println("\r\n 从堆中取数据")
	for i := len(heap.List); i > 0; i-- {
		fmt.Println(len(heap.List), " ", heap.Pop())
	}
}
