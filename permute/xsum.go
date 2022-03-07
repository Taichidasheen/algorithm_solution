package permute

import (
	"fmt"
	"sort"
)

/*
https://leetcode.com/problems/3sum/
 3sum问题思路：
1. 回溯法穷举出符合条件的组合，然后计算每个组合是否符合要求
参考链接：https://mp.weixin.qq.com/s/nrTpZ9b9RvfNsaEkJoHMvg
该算法的时间复杂度为O(n^3)，在leetcode上会报time limit exceeded
2. 采用三指针
参考链接：https://www.code-recipe.com/post/three-sum
该算法的时间复杂度为O(n^2)

扩展：
3sum问题可以扩展到ksum问题,解决思路为把ksum问题转变为k-1 sum问题，直到2sum问题。
ksum问题的解法参考https://leetcode.com/problems/4sum/discuss/8609/My-solution-generalized-for-kSums-in-JAVA
2sum问题可参考https://leetcode.com/problems/two-sum/
 */
const LengthOfSelects = 3
const SumOfSelects = 0

var res [][]int
//采用穷举遍历方法
func threeSum1(nums []int) [][]int {
	var path []int
	sort.Ints(nums)
	fmt.Printf("nums:%+v\n", nums)
	combinationX(nums, path, 0, 0)
	return res
}

func combinationX(nums []int, path []int, pathSum int, startI int) {
	if len(path) > LengthOfSelects || pathSum > 0 {
		return
	}
	if len(path) == LengthOfSelects && pathSum == SumOfSelects {
		fmt.Printf("path:%+v\n", path)
		var tmp []int
		tmp = append(tmp, path...)
		res = append(res, tmp)
		return
	}
	for i:= startI;i<len(nums);i++ {
		//进行选择操作
		//如果当前选择的num和前一个元素相同，则可以直接跳过，不用选择
		if i > startI && nums[i] == nums[i-1] {
			fmt.Printf("i:%d, nums[i]:%d\n", i, nums[i])
			continue
		}
		path = append(path, nums[i])
		pathSum = pathSum + nums[i]
		combinationX(nums, path, pathSum, i+1)
		//撤销选择操作
		path = path[0:len(path)-1]
		pathSum = pathSum - nums[i]
	}
}

//采用三指针的解法
func threeSum2(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for idx1:=0;idx1<len(nums)-2;idx1++ {
		//去重逻辑，防止出现相同nums[idx1]的情况
		if idx1 >0 && nums[idx1] == nums[idx1-1] {
			continue
		}
		idx2 := idx1 + 1
		idx3 := len(nums) - 1
		for idx2 < idx3 {
			sum := nums[idx1] + nums[idx2] + nums[idx3]
			if sum == 0 {
				result = append(result, []int{nums[idx1], nums[idx2], nums[idx3]})
				idx3 --
				//去重逻辑，防止出现相同nums[idx3]的情况
				for nums[idx3] == nums[idx3+1] {
					idx3 --
				}
			} else if sum < 0 {
				idx2 ++
			} else {
				idx3 --
			}
		}
	}
	return result
}


