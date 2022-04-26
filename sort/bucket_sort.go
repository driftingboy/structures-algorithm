package sort

import "sort"

type IBucket interface {
	sort.Interface
	GetData(index int) interface{}
}

// 将一个大桶的数据拆分为多个桶
type IBucketBuilder interface {
	// 获取bucket的数量
	BucketCount() int
	// 添加到指定桶
	AddToBucket(data interface{})
	// 获取指定桶
	Bucket(index int) IBucket
	Init(data []interface{})
}

func BucketSort(data []interface{}, builder IBucketBuilder) {
	builder.Init(data)

	for i := 0; i < len(data); i++ {
		d := data[i]
		builder.AddToBucket(d)
	}

	curIndex := 0

	for i := 0; i < builder.BucketCount(); i++ { //还可以启用多个协程去并发排序
		bucket := builder.Bucket(i)
		if bucket.Len() == 0 {
			continue
		}
		sort.Sort(bucket)
		for j := 0; j < bucket.Len(); j++ {
			data[curIndex] = bucket.GetData(j)
			curIndex++
		}
	}
}

type IntBucket struct {
	sort.IntSlice
}

func NewIntBucket(data []int) *IntBucket {
	return &IntBucket{
		data,
	}
}

func (ib IntBucket) GetData(index int) interface{} {
	return ib.IntSlice[index]
}

type IntBucketBuilder struct {
	max int
	min int
	// 一个桶的数据范围，比如 100 表示 [0,99][100, 199]...
	dataRange int
	data      [][]int
}

func NewIntBucketBuilder(dataRange int) *IntBucketBuilder {
	return &IntBucketBuilder{
		dataRange: dataRange,
	}
}

func (ibb *IntBucketBuilder) Init(data []interface{}) {
	var (
		max int
		min int
	)
	for i := 0; i < len(data); i++ {
		dint := data[i].(int)
		if max < dint {
			max = dint
		}
		if dint < min {
			min = dint
		}
	}

	if ibb.data == nil {
		ibb.data = make([][]int, ibb.dataRange+1)
	}
}

// 100的数据范围一个桶
func (ibb IntBucketBuilder) bucketIndex(data interface{}) (index int) {
	intData := data.(int)
	return (intData - ibb.min) / ibb.dataRange
}

// 获取bucket的数量
func (ibb IntBucketBuilder) BucketCount() int {
	return len(ibb.data)
}

// 添加到指定桶
func (ibb IntBucketBuilder) AddToBucket(data interface{}) {
	index := ibb.bucketIndex(data)
	ibb.data[index] = append(ibb.data[index], data.(int))
}

// 获取指定桶
func (ibb IntBucketBuilder) Bucket(index int) IBucket {
	return NewIntBucket(ibb.data[index])
}
