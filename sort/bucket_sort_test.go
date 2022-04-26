package sort

import (
	"testing"
)

func TestBucketSort(t *testing.T) {
	dataSet := make([]interface{}, 0, 120)
	for i := 100; i > 0; i-- {
		dataSet = append(dataSet, i)
	}
	// for i := 20; i > 0; i-- {
	// 	dataSet = append(dataSet, rand.Int31())
	// }

	type args struct {
		data    []interface{}
		builder IBucketBuilder
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "small data",
			args: args{
				data:    dataSet,
				builder: NewIntBucketBuilder(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BucketSort(tt.args.data, tt.args.builder)
		})
	}

	t.Logf("%+v", dataSet)
}
