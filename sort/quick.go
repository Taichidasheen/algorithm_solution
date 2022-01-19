package sort

import "fmt"

/*
参考链接：
 1. https://wiki.jikexueyuan.com/project/easy-learn-algorithm/fast-sort.html 算法思路
 2. https://www.runoob.com/w3cnote/quick-sort-2.html C++版最简洁，对应partitionV2

关于快排的时间复杂度。当对一个已经是正序的数组,或者完全是逆序的数组进行排序时，快排的时间复杂度是O(n*n)
快排的平均时间复杂度是O(n*logn)

快排为什么要从右边开始移动？考虑对于一个已经是正序的数组执行快排算法，如果从左边开始移动，结果是会交换第一个元素和第二个元素的位置

快排是稳定排序吗？不是的。考虑9,9,8,7,6序列。
在第一次的交换操作中，就把第一个9交换到了数组最后面，破坏了稳定性。
 */
func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		mid := partitionV2(arr, left, right)
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
		//case: arr := []int{5,4,3,2,1,8,9,4,10,2}
		if arr[right] >= arr[pivot]{
			right--
			continue
		}
		for left < right {
			if arr[left] <= arr[pivot] {
				left++
				continue
			}
			//swap right left（第一波替换操作）
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
		//case: arr := []int{5,4,3,2,1,8,9,4,10,2}
		//会一直卡在状态arr:[1 4 3 2 2 8 9 4 10 5], pivot:1, left:1, right:7（因为left和right指向的数值相等）
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
		fmt.Printf("mid arr:%+v, pivot:%d, left:%d, right:%d\n", arr, pivot, left, right)
	}
	//swap pivot
	temp := arr[left]
	arr[left] = arr[pivot]
	arr[pivot] = temp
	fmt.Printf("output arr:%+v, left:%d, right:%d\n", arr, left, right)
	return left
}