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

