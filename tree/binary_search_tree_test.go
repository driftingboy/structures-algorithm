package tree

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTree_Insert(t *testing.T) {
	bsTree := NewBSTree()
	nodeValues := []int{3, 4, 2, 6, 5, 9, 1}

	for _, v := range nodeValues {
		bsTree.Insert(v)
	}

	sort.Ints(nodeValues)
	result := bsTree.ModPrint()
	assert.Equal(t, nodeValues, result)
}

func TestBSTree_Delete(t *testing.T) {
	bsTree := NewBSTree()
	nodeValues := []int{3, 4, 2, 6, 5, 9, 1}

	for _, v := range nodeValues {
		bsTree.Insert(v)
	}

	// single child node
	bsTree.Delete(2)
	assert.Equal(t, []int{1, 3, 4, 5, 6, 9}, bsTree.ModPrint())
	// no child node
	bsTree.Delete(9)
	assert.Equal(t, []int{1, 3, 4, 5, 6}, bsTree.ModPrint())
	// double child node
	bsTree.Delete(4)
	assert.Equal(t, []int{1, 3, 5, 6}, bsTree.ModPrint())

}
