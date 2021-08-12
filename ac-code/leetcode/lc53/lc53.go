package lc53

import "sync"

/*
题目215：数组中的第K个最大元素
	给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
	请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
*/

func MaxSubArray(nums []int) int {
	_, _, max := FindMaxSubArray(nums, 0, len(nums)-1)
	return max
}

// 2分查找最大子序列(不跨越mid点)
func FindMaxSubArray(nums []int, low, high int) (start, end, max int) {
	if low >= high {
		// 递归到递归树的叶子节点
		return low, high, nums[low]
	}
	mid := (low + high) / 2
	// 线性查找(跨越mid点)
	cross_low, cross_high, cross_max := FindMaxCrossSubArray(nums, low, mid, high)
	// 查找左半边
	left_low, left_high, left_max := FindMaxSubArray(nums, low, mid)
	// 查找右半边
	right_low, right_high, right_max := FindMaxSubArray(nums, mid+1, high)

	// 找出最大的子序列，返回
	switch {
	case cross_max >= left_max && cross_max >= right_max:
		return cross_low, cross_high, cross_max
	case left_max >= cross_max && left_max >= right_max:
		return left_low, left_high, left_max
	default:
		return right_low, right_high, right_max
	}
}

// 线性查找最大子序列(跨越mid点)
func FindMaxCrossSubArray(nums []int, low, mid, high int) (start, end, max int) {
	sum, left_i, right_i := 0, mid, mid+1
	left_max, right_max := nums[left_i], nums[right_i]
	for i := left_i; i >= low; i-- {
		sum += nums[i]
		if sum >= left_max {
			left_i = i
			left_max = sum
		}
	}
	sum = 0
	for i := right_i; i <= high; i++ {
		sum += nums[i]
		if sum >= right_max {
			right_i = i
			right_max = sum
		}
	}
	return left_i, right_i, left_max + right_max
}

// 线性查找最大子序列(跨越mid点)-使用goroutine，不如直接循环，可能是切换协程有损耗
func FindMaxCrossSubArray2(nums []int, low, mid, high int) (start, end, max int) {
	left_i, right_i := mid, mid+1
	left_max, right_max := nums[left_i], nums[right_i]
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		sum := 0
		for i := left_i; i >= low; i-- {
			sum += nums[i]
			if sum >= left_max {
				left_i = i
				left_max = sum
			}
		}
		wg.Done()
	}()
	go func() {
		sum := 0
		for i := right_i; i <= high; i++ {
			sum += nums[i]
			if sum >= right_max {
				right_i = i
				right_max = sum
			}
		}
		wg.Done()
	}()
	wg.Wait()
	return left_i, right_i, left_max + right_max
}

// 线性查找最大子序列(跨越mid点)-使用goroutine和chan，不如直接循环，可能是切换协程有损耗，以及chan的内存消耗
func FindMaxCrossSubArray3(nums []int, low, mid, high int) (start, end, max int) {
	ch_low, ch_high := make(chan int), make(chan int)
	go func() {
		left_i := mid
		if left_i < 0 {
			ch_low <- mid
			ch_low <- nums[mid]
			return
		}
		sum, max := 0, nums[left_i]
		for i := left_i; i >= low; i-- {
			sum += nums[i]
			if sum >= max {
				left_i = i
				max = sum
			}
		}
		ch_low <- left_i
		ch_low <- max
	}()
	go func() {
		right_i := mid + 1
		sum, max := 0, nums[right_i]
		for i := right_i; i <= high; i++ {
			sum += nums[i]
			if sum >= max {
				right_i = i
				max = sum
			}
		}
		ch_high <- right_i
		ch_high <- max
	}()
	return <-ch_low, <-ch_high, <-ch_low + <-ch_high
}
