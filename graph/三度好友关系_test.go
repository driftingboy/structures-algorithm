package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_Recommend(t *testing.T) {
	g := NewGraph()
	g.Follow("1", "10")
	g.Follow("1", "11")
	g.Follow("1", "12")
	g.Follow("2", "12")
	g.Follow("12", "1")

	g.Follow("10", "100")
	g.Follow("11", "110")

	g.Follow("100", "1000")
	g.Follow("100", "1001")

	assert.Equal(t, []string{"10", "11", "12"}, g.ListFollowers("1"))
	assert.Equal(t, []string{"1", "2"}, g.ListFans("12"))
	// assert.True(t, g.IsFriend("1", "2"))
	assert.Equal(t, []string{"100", "110"}, g.Recommend("1", 2))
	assert.Equal(t, []string{"1000", "1001"}, g.Recommend("1", 3))

	assert.Equal(t, []string{"10", "100", "1000"}, g.Link("1", "1000"))
	assert.Equal(t, []string{"10", "100", "1000", "1001"}, g.Link("1", "1001"))
}
