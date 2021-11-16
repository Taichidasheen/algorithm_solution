package sort

import "fmt"

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		mid := partition(arr, left, right)
		_quickSort(arr, left, mid-1)
		_quickSort(arr, mid+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	fmt.Printf("input arr:%+v, left:%d, right:%d\n", arr, left, right)
	pivot := left
	for right > left {
		//注意：这里需要加上等于的条件，不然可能会陷入死循环，防止基准值在第一次替换操作中被替换
		if arr[right] >= arr[pivot]{
			right--
			continue
		}
		for left < right {
			if arr[left] <= arr[pivot] {
				left++
				continue
			}
			//swap right left（第一次替换操作）
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
			fmt.Printf("mid arr:%+v, left:%d, right:%d\n", arr, left, right)
			//下面这行是不必要的
			//right--
			break
		}
	}
	//swap pivot（第二次替换操作）
	temp := arr[left]
	arr[left] = arr[pivot]
	arr[pivot] = temp
	fmt.Printf("output arr:%+v, left:%d, right:%d\n", arr, left, right)
	return left
}

func partitionV2(arr []int, left, right int) int {
	fmt.Printf("input arr:%+v, left:%d, right:%d\n", arr, left, right)
	pivot := left
	for right > left {
		//注意：这里需要加上等于的条件，不然可能会陷入死循环
		for right > left && arr[right] >= arr[pivot] {
			right--
		}
		for right > left && arr[left] <= arr[pivot] {
			left++
		}
		//swap right left
		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp
		fmt.Printf("mid arr:%+v, left:%d, right:%d\n", arr, left, right)
	}
	//swap pivot
	temp := arr[left]
	arr[left] = arr[pivot]
	arr[pivot] = temp
	fmt.Printf("output arr:%+v, left:%d, right:%d\n", arr, left, right)
	return left
}