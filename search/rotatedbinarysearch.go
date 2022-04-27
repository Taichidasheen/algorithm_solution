package search

/*
leetcode-33:
https://leetcode.cn/problems/search-in-rotated-sorted-array/
 */
func rotatedBinarySearch(nums []int, target int) int {
	var left, right int
	left = 0
	right = len(nums)-1
	for left <=right {
		mid := (left + right)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { //说明左侧是有序的,注：等于可能不成立，因为题目说数组中的值互不相同
			if target <= nums[mid] && target >= nums[left] {//target在有序这一侧
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //说明左侧是无序的，右侧是有序的
			if target >= nums[mid] && target <= nums[right] {//target在有序这一侧
				left = mid + 1
			} else {
				right = mid -1
			}
		}
	}
	return -1
}