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
