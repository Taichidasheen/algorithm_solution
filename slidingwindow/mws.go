package slidingwindow

import "fmt"

/*
 lc-76: https://leetcode.com/problems/minimum-window-substring/
 */

func minWindow(s string, t string) string {
	var tmap map[uint8]int = make(map[uint8]int)
	var window map[uint8]int = make(map[uint8]int)
	for i:=0;i<len(t);i++ {
		tmap[t[i]] ++
	}
	var left, right int
	var valid int
	var start,end int
	end = -1
	shortest := 1000000
	//扩大窗口，移动right
	for right < len(s) {
		c := s[right]
		fmt.Printf("c:%+v\n", string(c))
		if  _, ok := tmap[c]; ok {
			window[c]++
			if window[c] == tmap[c]{
				valid++
			}
		}
		//缩减窗口，移动left
		for valid == len(tmap) {
			if right - left < shortest {
				shortest = right - left
				end = right
				start = left
			}
			h := s[left]
			fmt.Printf("h:%+v\n", string(h))

			if _, ok := tmap[h]; ok {
				if window[h] == tmap[h] {
					valid--
				}
				window[h]--
			}
			left++
		}
		right++
	}
	fmt.Printf("start:%d, end:%d\n", start, end)
	return s[start:end+1]
}