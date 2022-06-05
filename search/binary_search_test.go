package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {

	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 13, 14}

	target := 12
	gotIndex := BinarySearch(sorted, target)
	assert.Equal(t, 11, gotIndex)

	target = 13
	gotIndex = BinarySearch(sorted, target)
	assert.Equal(t, 13, gotIndex)

}
