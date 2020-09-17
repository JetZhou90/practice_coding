package sort

import (
	"fmt"
)

func Test() {
	var nums []int = []int{2, 32, 12, 34, 3, 1, 8, 55}
	fmt.Println(nums)
	// fmt.Println("quick sort", QuickSort(nums))
	// fmt.Println("merge sort", MergeSort(nums))
	fmt.Println("heap sort", HeapSort(nums))

}

/*
稳定：如果a原本在b前面，而a=b，排序之后a仍然在b的前面。
不稳定：如果a原本在b的前面，而a=b，排序之后 a 可能会出现在 b 的后面。
时间复杂度：对排序数据的总的操作次数。反映当n变化时，操作次数呈现什么规律。
空间复杂度：是指算法在计算机内执行时所需存储空间的度量，它也是数据规模n的函数
*/

/*
快速排序
时间复杂度：nlogn
空间复杂度：nlogn
算法稳定性：不稳定
*/

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums

}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end] // 基准比较值
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

/*
归并排序
时间复杂度：nlogn
空间复杂度：n
算法稳定性：稳定
*/
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// 分治法：divide 分为两段
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// 合并两段数据
	result := merge(left, right)
	return result
}
func merge(left, right []int) (result []int) {
	// 两边数组合并游标
	l := 0
	r := 0
	// 注意不能越界
	for l < len(left) && r < len(right) {
		// 谁小合并谁
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 剩余部分合并
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

/*
堆排序
时间复杂度：nlogn
空间复杂度：1
算法稳定性：不稳定
*/
func HeapSort(nums []int) []int {
	// 1、无序数组nums
	// 2、将无序数组nums构建为一个大根堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i, len(nums))
	}
	// 3、交换a[0]和a[len(a)-1]
	// 4、然后把前面这段数组继续下沉保持堆结构，如此循环即可
	for i := len(nums) - 1; i >= 1; i-- {
		// 从后往前填充值
		swap(nums, 0, i)
		// 前面的长度也减一
		sink(nums, 0, i)
	}
	return nums
}

func sink(nums []int, i, length int) {
	for {
		left := i*2 + 1
		right := i*2 + 2
		index := i
		if left < length && nums[left] > nums[index] {
			index = left
		}
		if right < length && nums[right] > nums[index] {
			index = right
		}
		if index == i {
			break
		}
		swap(nums, i, index)
		i = index
	}
}
