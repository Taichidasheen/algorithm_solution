package sort

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 { //注意这个条件要加上，不然递归不会停止，最后报栈溢出
		return arr
	}
	mid := len(arr)/2
	fmt.Println("mid:", mid)
	arr1 := mergeSort(arr[0:mid])
	arr2 := mergeSort(arr[mid:])
	res := merge(arr1, arr2)
	return res
}

func merge(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}
	index1 := 0
	index2 := 0
	var res []int
	for {
		//如果一个数组已经全部处理完毕，则把另一个数组剩下的数据全部添加到结果数组中
		if index1 >= len(arr1) {
			res = append(res, arr2[index2:]...)
			break
		}
		if index2 >= len(arr2) {
			res = append(res, arr1[index1:]...)
			break
		}
		if arr1[index1] <= arr2[index2] {
			res = append(res, arr1[index1])
			index1++
		} else {
			res = append(res, arr2[index2])
			index2++
		}
	}
	return res
}
