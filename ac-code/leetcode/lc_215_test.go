package leetcode

import (
	"fmt"
	"testing"
)

/*
题目215：数组中的第K个最大元素
	给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
	请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
*/
func TestFindKthLargest(t *testing.T) {
	arr, k := []int{3, 2, 1, 5, 6, 4}, 2
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出：%+v\n", FindKthLargest(arr, k))

	arr, k = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出：%+v\n", FindKthLargest(arr, k))
}

func FindKthLargest(nums []int, k int) int {
	length := len(nums)
	// 初始化构建大顶堆
	for i := length / 2; i >= 0; i-- {
		AdjustHeap(&nums, i, length)
	}
	// 循环调整堆顶k-1次，
	for i := length - 1; i > length-1-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		AdjustHeap(&nums, 0, i)
	}
	return nums[length-k]
}

func AdjustHeap(arr *[]int, k, length int) {
	old_parent := (*arr)[k]
	for i := k * 2; i < length; i *= 2 {
		if i+1 < length && (*arr)[i] < (*arr)[i+1] {
			i++
		}
		if (*arr)[i] > old_parent {
			(*arr)[k] = (*arr)[i]
			k = i
			continue
		}
		break
	}
	(*arr)[k] = old_parent
}
