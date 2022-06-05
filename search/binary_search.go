package search

// 如果有重复值只能获取到某一个值的下标
// 所以只能处理不重复的数据集合，或重复的数据集判断是否某数据存在
func BinarySearch(sorted []int, target int) (index int) {
	if len(sorted) == 0 {
		return
	}

	left, right := 0, len(sorted)-1
	mid := 0
	for left <= right { // 注意等于
		mid = left + (right-left)>>1 // 防止溢出

		if target < sorted[mid] {
			right = mid - 1
		} else if sorted[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
