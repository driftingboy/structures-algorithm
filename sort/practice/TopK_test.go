package practice

import (
	"github/driftingboy/structures-algorithm/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TopK(t *testing.T) {
	// build heap
	h := sort.NewHeap(5)
	data := []int{7, 8, 9, 1, 2, 3, 5, 6, 7, 11, 17, 18, 29, 1, 21, 8, 5, 6, 71}
	for _, d := range data {
		if d < h.Head() && h.Len() >= 5 {
			continue
		}
		if h.Len() == 5 {
			_ = h.Pop()
		}
		h.Push(d)
	}

	// 前五的值
	gotData, wantData := h.SortValues(), []int{71, 29, 21, 18, 17}
	assert.Equal(t, wantData, gotData)

	// 动态 top5，实际使用要时注意， h.SortValues() 和 h.Push(d)、h.Pop() 需要加 读写锁
	h.Pop()
	h.Push(100)
	gotData, wantData = h.SortValues(), []int{100, 71, 29, 21, 18}
	assert.Equal(t, wantData, gotData)

}

// 优先级队列
