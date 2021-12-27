package common

import "math/rand"

/*
 lc-470: https://leetcode.com/problems/implement-rand10-using-rand7/
 */
//Runtime: 5 ms, faster than 5.00%
func rand10() int {
	for {
		num := (rand7() -1) * 7 + rand7()
		if num > 10 {
			continue
		}
		return num
	}
}

//Runtime: 16 ms, faster than 28.57%
func rand10V2() int {
	for {
		num := ((rand7() -1) * 7 + rand7())*3
		if num > 100 || num==30 || num==60 || num==90 {
			continue
		}
		return num/10 + 1
	}
}

//Runtime: 8 ms, faster than 100.00%
func rand10V3() int {
	for {
		num := (rand7() -1) * 7 + rand7()
		if num-1 >= 40 {
			continue
		}
		return (num-1)/4 + 1
	}
}

func rand7() int {
	random := rand.Intn(7)
	return random + 1
}