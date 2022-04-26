package sort

import (
	"math/rand"
	"time"
)

// sort.Ints() // 标准库 快排、二分查找

func BubbleSort(ints []int) {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-i-1; j++ {
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
}

func InsertSort(ints []int) {
	for i := 0; i < len(ints); i++ {
		temp := ints[i]

		j := i - 1
		for ; j >= 0 && temp < ints[j]; j-- {
			ints[j+1] = ints[j]
		}

		ints[j+1] = temp
	}
}

// MergeSort(i,j) = Merge(MergeSort(i,mid), MergeSort(mid+1,j))
// return when i >= j
func MergeSort(ints []int) {
	if len(ints) == 0 {
		return
	}

	mergeCore(ints, 0, len(ints)-1)
}

func mergeCore(ints []int, n, m int) {
	if n >= m {
		return
	}

	mid := (n + m) / 2
	mergeCore(ints, n, mid)
	mergeCore(ints, mid+1, m)
	// note, slice [n:m+1)
	merge(ints[n:m+1], ints[n:mid+1], ints[mid+1:m+1])
}

func merge(result []int, p1, p2 []int) {
	temp := make([]int, 0, len(result))

	// compare and set min value
	i, j := 0, 0
	for i < len(p1) && j < len(p2) {
		min := 0
		if p1[i] < p2[j] {
			min = p1[i]
			i++
		} else {
			min = p2[j]
			j++
		}
		temp = append(temp, min)
	}

	// set remain value in p1 or p2
	if i < len(p1) {
		temp = append(temp, p1[i:]...)
	} else if j < len(p2) {
		temp = append(temp, p2[j:]...)
	}

	// cover in result
	result = result[:0]
	_ = append(result, temp...)
}

func QuickSort(ints []int) {
	if len(ints) == 0 {
		return
	}

	quickSortCore(ints, 0, len(ints)-1)
}

func quickSortCore(ints []int, n, m int) {
	if n >= m {
		return
	}

	pivot := partition(ints, n, m)
	quickSortCore(ints, n, pivot-1)
	quickSortCore(ints, pivot+1, m)
}

func partition(ints []int, n, m int) (pivot int) {
	pivotV := ints[m]

	i, j := n, n
	for i <= m && j <= m {
		if ints[j] < pivotV {
			j++
		} else {
			ints[j], ints[i] = ints[i], ints[j]
			i++
			j++
		}
	}

	if i == n {
		return n
	}
	return i - 1
}

// 快速排序O(n) 时间复杂度内求无序数组中的第 K 大元素。比如，4， 2， 5， 12， 3 这样一组数据，第 3 大元素就是 4。
func FindKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	rand.Seed(time.Now().UnixNano())
	return findKthLargestCore(nums, 0, len(nums)-1, k)
}

func findKthLargestCore(nums []int, n, m int, k int) int {
	// 找到中间点
	pivot := partitionDescRandmon(nums, n, m)
	if k == pivot+1 {
		return nums[pivot]
	} else if k < pivot+1 {
		return findKthLargestCore(nums, n, pivot-1, k)
	} else {
		return findKthLargestCore(nums, pivot+1, m, k)
	}
}

func partitionDescRandmon(nums []int, n, m int) int {
	// 取随机 pivot
	i := rand.Int()%(m-n+1) + n
	nums[i], nums[m] = nums[m], nums[i]
	return partitionDesc(nums, n, m)
}

// partition [n,m]获取第x大元素，返回其下标
func partitionDesc(nums []int, n, m int) int {
	pivotV := nums[m]
	i, j := n, n

	for ; j <= m; j++ {
		if nums[j] >= pivotV { // desc
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	if i == n {
		return n
	}
	return i - 1
}
