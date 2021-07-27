package attempt

import "fmt"

/*
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/

// 数组截取切片
func SliceArray() {
	var array [10]int
	// array := make([]int, 10)
	fmt.Printf("array: len(%d), cap(%d)\n", len(array), cap(array))
	slice := array[5:6]
	fmt.Printf("slice: len(%d), cap(%d)\n", len(slice), cap(slice))
	fmt.Println(&slice[0] == &array[5])
}

// 切片append
func SliceAppend() {
	var slice []int
	fmt.Printf("before: len(%d), cap(%d), pointer(%p)\n", len(slice), cap(slice), &slice)
	slice = append(slice, 1, 2, 3)
	fmt.Printf("append: len(%d), cap(%d), pointer(%p)\n", len(slice), cap(slice), &slice)
	newSlice := AddElemnt(slice, 4)
	fmt.Printf("AddElemnt: len(%d), cap(%d), pointer(%p)\n", len(newSlice), cap(newSlice), &newSlice)
	fmt.Println(&slice[0] == &newSlice[0])
}
func AddElemnt(slice []int, e int) []int {
	fmt.Printf("append1: len(%d), cap(%d), pointer(%p)\n", len(slice), cap(slice), &slice)
	return append(slice, e)
}

// 切片截取
func Slice() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}
