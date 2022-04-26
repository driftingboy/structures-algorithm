package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	ints []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{name: "nil", args: args{nil}, want: nil},
	{name: "empty", args: args{[]int{}}, want: []int{}}, // 3 7 8 4 3 5 2  a[i]=4  a[j]=3 temp=a[i] a[j+2],a[j+1:i ] a[j+1] = temp
	{name: "small", args: args{[]int{8, 7, 3, 4, 3, 5, 2, 1}}, want: []int{1, 2, 3, 3, 4, 5, 7, 8}},
	{name: "sorted", args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8}},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.ints)
			assert.Equal(t, tt.want, tt.args.ints)
		})
	}
}

func TestInsertSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.ints)
			assert.Equal(t, tt.want, tt.args.ints)
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.ints)
			assert.Equal(t, tt.want, tt.args.ints)
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.ints)
			assert.Equal(t, tt.want, tt.args.ints)
		})
	}
}

func TestTheKLargest(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV := FindKthLargest(tt.args.ints, 4)
			if gotV != -1 {
				assert.Equal(t, tt.want[4], gotV)
			}
		})
	}
}
