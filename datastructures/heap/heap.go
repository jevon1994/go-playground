package heap

import "fmt"

type Heap struct {
	data     []int
	Size     int
	Capacity int
}

func NewHeap(cap int) *Heap {
	var i = make([]int, cap+1)
	//sentinel
	i[0] = 1000
	return &Heap{i, 0, cap}
}

func (h *Heap) Insert(val int) {
	h.Size += 1
	i := h.Size
	//i  数组下标
	for ; h.data[i/2] < val; i /= 2 {
		h.data[i] = h.data[i/2]
	}
	h.data[i] = val
}

func (h *Heap) IsFull() bool {
	return h.Size == h.Capacity
}

func (h *Heap) IsEmpty() bool {
	return h.Size == 0
}

func (h *Heap) DeleteMax() int {
	var parent, child, Max, tmp int
	Max = h.data[1]
	tmp = h.data[h.Size]
	h.Size -= 1
	//是否有左子树
	for parent = 1; parent*2 <= h.Size; parent = child {
		child = parent * 2
		//child 不为当前最后一个结点，即 parent 有右孩子结点
		if child != h.Size && h.data[child+1] > h.data[child] {
			child += 1 // 返回左右孩子较大的
		}
		// 最后元素都比左右孩子小, 不需要调整
		if tmp >= h.data[child] {
			break
		} else {
			//否则把大的孩子提回来
			h.data[parent] = h.data[child]
		}
	}
	h.data[parent] = tmp
	fmt.Print("delete max : ", Max, "\n")
	return Max
}

func (h *Heap) BFSTravesal() {
	for i := 1; i <= h.Size; i++ {
		if h.data[i] == 0 {
			break
		}
		fmt.Print(h.data[i], "\n")
	}
}
