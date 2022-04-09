package common

import "fmt"

/*
https://leetcode-cn.com/problems/sqrtx/
 */
func mySqrt(x int) int {
	if x==0 {
		return 0
	}
	var left, right int
	left = 1
	right = x
	//注意这里的条件不是left<right，而是left<right-1。
	//如果是left<right，则当输入x=24时，会执行到left=4，right=5然后一直循环下去
	for left < right - 1 {
		mid := (left + right)/2
		fmt.Println("mid:", mid, "left:", left, "right:", right)
		//这里用除法而不是乘法是为了防止溢出
		tmp := x/mid
		if tmp > mid {
			left = mid //注意这里不要+1
		} else if tmp == mid {
			return mid
		} else {
			right = mid //注意这里不要-1
		}
	}
	return left
}
