package search

/*
 lc-704: https://leetcode.com/problems/binary-search/
 */

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

//递归解法
func binarySearch(nums []int, left, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right)/2
	if target > nums[mid] {
		return binarySearch(nums, mid+1, right, target)
	} else if target < nums[mid] {
		return binarySearch(nums, left, mid-1, target)
	} else {
		//target == nums[mid]
		return mid
	}
}

//非递归解法
func binarySearchV2(nums []int, left, right int, target int) int {
	for left <= right {
		mid := (left + right)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid -1
		} else {
			return mid
		}
	}
	return -1
}

//注意寻找下界的方法和二分查找很像，注意最后返回的是right的下标
//分析好两种边界情况就可以了
//第一种：三个数依次为target，left，right
//第二种：三个数依次为left，right，target
func findLowBoundIndex(nums []int, left, right int, target int) int {
	//注意：这里left不能和right相等
	for left <= right {
		mid := (left + right)/2
		if target > mid {
			left = mid + 1 //注意这里有可能导致nums[left]比target大
		} else if target < mid {
			right = mid -1 //注意这里有可能导致nums[right]比target小
		} else {
			return mid
		}
	}
	//for循环结束时，right < left
	return right
}

