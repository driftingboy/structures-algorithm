package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	a, want := []int{7, 8, 9, 1, 2, 3, 5, 6, 7}, []int{9, 8, 7, 7, 6, 5, 3, 2, 1}
	err := HeapSort(a, 10)
	assert.NoError(t, err)

	assert.Equal(t, want, a)
	// h := NewHeap(10)
	// a := []int{7, 8, 9, 1, 2, 3, 5, 6, 7}
	// err := h.Init(a)
	// assert.NoError(t, err)
	// t.Log(a)

}
