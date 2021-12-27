package sort

/*
 参考：https://www.runoob.com/w3cnote/bubble-sort.html
 */
//下面这个才是正确的冒泡排序
func bubbleSort(arr []int) []int {
	for i := 1; i <= len(arr)-1; i++ { //n个元素排序，最多执行n-1趟
		flag := false
		for j := 0; j < len(arr)-1 -i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return arr
}

//下面这种不是冒泡排序，冒泡排序的核心：
//1. 是每趟把最大的元素放到最后面
//2. 每次交换发生在相邻的元素之间
func bubbleSortV2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return arr
}
