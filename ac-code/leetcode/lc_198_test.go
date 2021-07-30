package leetcode

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	arr := []int{1, 2, 3, 1}
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出：%+v\n", Rob(arr))
	arr = []int{2, 7, 9, 3, 1}
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出：%+v\n", Rob(arr))
}

// 循环写法(正解，加result缓存，使用空间换取时间)
func Rob(nums []int) int {
	length := len(nums)
	result := make([]int, length)
	copy(result, nums)
	max := 0
	for i := 0; i < length; i++ {
		for j := i + 2; j < length; j++ {
			sum := nums[j] + result[i]
			if sum > result[j] {
				result[j] = sum
			}
		}
		if result[i] > max {
			max = result[i]
		}
	}
	fmt.Println(result)
	return max
}

// 递归写法(超时，未使用缓存，导致重复计算，耗费时间)
func Rob1(nums []int) int {
	length := len(nums)
	switch length {
	case 0:
		return 0
	case 1:
		return nums[0]
	}
	max := 0
	for i := 0; i < length; i++ {
		temp := 0
		for j := i + 2; j < length; j++ {
			p := Rob(nums[j:])
			if p > temp {
				temp = p
			}
		}
		temp += nums[i]
		if temp > max {
			max = temp
		}
	}
	return max
}
