package sort

import "errors"

// "container/heap"
type Heap struct {
	cap  int
	ints []int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		cap: cap,
	}
}

func (h *Heap) Init(ints []int) error {
	if len(ints) > h.cap {
		return errors.New("over cap")
	}
	h.ints = ints
	n := len(h.ints)
	for i := n/2 - 1; i >= 0; i-- {
		down(h.ints, i, n-1)
	}
	return nil
}

func (h *Heap) Len() int {
	return len(h.ints)
}

func (h *Heap) Head() int {
	if len(h.ints) > 0 {
		return h.ints[0]
	}
	return -1
}

// parentIndex = (subIndex-1)/2
// 自下而上构建堆
func (h *Heap) Push(data int) {
	if len(h.ints) >= h.cap {
		return
	}
	h.ints = append(h.ints, data)

	for i := len(h.ints) - 1; ; {
		pIndex := (i - 1) / 2
		if i <= 0 || h.ints[i] >= h.ints[pIndex] {
			break
		}

		h.ints[i], h.ints[pIndex] = h.ints[pIndex], h.ints[i]
		i = pIndex
	}
}

// 自上而下构建堆
func (h *Heap) Pop() int {
	if len(h.ints) == 0 {
		return -1
	}

	n := len(h.ints) - 1
	h.ints[0], h.ints[n] = h.ints[n], h.ints[0]

	down(h.ints, 0, n-1)

	result := h.ints[n]
	h.ints = h.ints[:n]
	return result
}

func (h *Heap) SortValues() []int {
	result := make([]int, len(h.ints))
	copy(result, h.ints)

	end, start := len(result)-1, 0
	for end > 0 {
		result[end], result[start] = result[start], result[end]
		end--
		down(result, 0, end)
	}

	return result
}

func HeapSort(ints []int, cap int) error {
	h := NewHeap(cap)
	if err := h.Init(ints); err != nil {
		return err
	}

	end, start := len(ints)-1, 0
	for end > 0 {
		h.ints[end], h.ints[start] = h.ints[start], h.ints[end]
		end--
		down(h.ints, 0, end)
	}
	return nil
}

func down(ints []int, i, n int) {
	for {
		subIdx := i*2 + 1
		if subIdx > n || i < 0 { // overflow
			break
		}

		if rIndex := subIdx + 1; rIndex <= n && ints[rIndex] < ints[subIdx] {
			subIdx = rIndex // set min value's index, right
		}

		if ints[subIdx] > ints[i] {
			break
		}
		ints[i], ints[subIdx] = ints[subIdx], ints[i]
		i = subIdx
	}
}
