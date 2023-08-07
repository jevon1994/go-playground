package heap

import "fmt"

type Heap struct {
	Data     []int
	Size     int
	Capacity int
}

func NewHeap(cap int) *Heap {
	var i = make([]int, cap+1)
	//sentinel
	i[0] = 1000
	return &Heap{i, 0, cap}
}

// arr[1] as root
func (h *Heap) parent(index int) int {
	return index / 2
}

func (h *Heap) left(index int) int {
	return index * 2
}

func (h *Heap) right(index int) int {
	return index*2 + 1
}

func (h *Heap) up(index int) {
	for index > 1 && h.less(h.parent(index), index) {
		h.swap(h.parent(index), index)
	}
}

func (h *Heap) swap(index1, index2 int) {
	h.Data[index1] ^= h.Data[index2]
	h.Data[index2] ^= h.Data[index1]
	h.Data[index1] ^= h.Data[index2]
}

func (h *Heap) down(index int) {
	for h.left(index) <= h.Size {
		max := h.left(index)
		if h.right(index) <= h.Size && h.less(max, h.right(index)) {
			max = h.right(index)
		}
		if h.less(max, index) {
			break
		}
		h.swap(index, max)
		index = max
	}
}

func (h *Heap) less(index1, index2 int) bool {
	return h.Data[index1] < h.Data[index2]
}

func (h *Heap) InsertWithUp(val int) {
	h.Size += 1
	h.Data[h.Size] = val
	h.up(h.Size)
}

func (h *Heap) DeleteWithDown(index int) int {
	max := h.Data[1]
	h.swap(index, 1)
	h.Data[h.Size] = -1
	h.Size--
	h.down(1)
	return max
}

func (h *Heap) Insert(val int) {
	h.Size += 1
	i := h.Size
	//向上找比 val 小的节点
	for ; h.Data[i/2] < val; i /= 2 {
		h.Data[i] = h.Data[i/2] // 比 val 小的放到下面
	}
	h.Data[i] = val // 找到 data[i/2] 比 val 大的替换
}

func (h *Heap) IsFull() bool {
	return h.Size == h.Capacity
}

func (h *Heap) IsEmpty() bool {
	return h.Size == 0
}

func (h *Heap) DeleteMax() int {
	var parent, child, Max, tmp int
	Max = h.Data[1]
	tmp = h.Data[h.Size]
	h.Size -= 1
	//是否有左子树
	for parent = 1; parent*2 <= h.Size; parent = child {
		child = parent * 2
		//child 不为当前最后一个结点，即 parent 有右孩子结点
		if child != h.Size && h.Data[child+1] > h.Data[child] {
			child += 1 // 返回左右孩子较大的
		}
		// 最后元素都比左右孩子小, 不需要调整
		if tmp >= h.Data[child] {
			break
		} else {
			//否则把大的孩子提回来
			h.Data[parent] = h.Data[child]
		}
	}
	h.Data[parent] = tmp
	fmt.Print("delete max : ", Max, "\n")
	return Max
}

//levelorder
func (h *Heap) BFSTravesal() {
	for i := 1; i <= h.Size; i++ {
		if h.Data[i] == 0 {
			break
		}
		fmt.Print(h.Data[i], "\n")
	}
}
