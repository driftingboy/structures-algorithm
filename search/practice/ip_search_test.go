package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 判读是否在某一个ip范围内
func TestIPTable_ToPlace(t *testing.T) {
	table, err := NewIPTable(map[string]string{
		"110.0.0.1,110.0.1.0":      "上海",
		"110.0.10.1,110.0.20.1":    "苏州",
		"110.10.0.1,110.100.0.1":   "四川",
		"168.192.0.1,168.192.10.0": "杭州",
	})
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	assert.Equal(t, "notfound", table.ToPlace("110.0.0.0"))
	assert.Equal(t, "notfound", table.ToPlace("110.0.1.125"))
	assert.Equal(t, "notfound", table.ToPlace("168.192.10.1"))
	assert.Equal(t, "上海", table.ToPlace("110.0.0.125"))
	assert.Equal(t, "上海", table.ToPlace("110.0.1.0"))
	assert.Equal(t, "杭州", table.ToPlace("168.192.0.1"))
	assert.Equal(t, "杭州", table.ToPlace("168.192.1.0"))
}

func TestIPToInt(t *testing.T) {

	got := IPToInt("110.0.0.1")
	assert.Equal(t, 1845493761, got)

}
