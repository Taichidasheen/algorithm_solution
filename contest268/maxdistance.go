package contest268

/*
 https://leetcode.com/contest/weekly-contest-268/problems/two-furthest-houses-with-different-colors/
 */
func maxDistance(colors []int) int {
	var maxdistance int
	for i:=0;i<len(colors);i++ {
		var distance int
		for j:=len(colors)-1;j>i;j-- {
			if colors[j] != colors[i] {
				distance = j - i
				break
			}
		}
		if distance > maxdistance {
			maxdistance = distance
		}
	}
	return maxdistance
}
