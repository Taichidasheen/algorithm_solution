package contest268

import "fmt"


//找到value,或者比value小的最大的index
func binarySearchLowBound(arr []int, value int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right)/2
		fmt.Printf("first left:%d, right:%d, mid:%d\n", left, right, mid)
		if arr[mid] < value {
			left = mid
		} else if arr[mid] > value {
			right = mid - 1
		} else {
			//arr[mid] == value
			return mid
		}
		fmt.Printf("second left:%d, right:%d\n", left, right)
	}
	return left
}

//找到value，或者比value大的最小的index
func binarySearchUpperBound(arr []int, value int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right)/2
		fmt.Printf("first left:%d, right:%d, mid:%d\n", left, right, mid)
		if arr[mid] < value {
			left = mid + 1
		} else if arr[mid] > value {
			right = mid - 1
		} else {
			//arr[mid] == value
			return mid
		}
		fmt.Printf("second left:%d, right:%d\n", left, right)
	}
	return right
}
