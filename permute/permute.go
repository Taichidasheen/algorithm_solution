package permute

import "fmt"

//输入一个数组，计算出所有的排列方式
func permute(selects []int, path []int, pathMap map[int]int) {
	//如果已经满足要求，加入到结果集中
	if len(path) == len(selects) {
		var tmp []int
		tmp = append(tmp, path...)
		results = append(results, tmp)
		return
	}
	for _, num := range selects {
		//如果已经在路径中，则不再选择
		if pathMap[num]> 0 {
			continue
		}
		//进行选择操作
		path = append(path, num)
		pathMap[num] = 1
		permute(selects, path, pathMap)
		//撤销选择操作
		path = path[0:len(path)-1]
		pathMap[num] = 0
	}
}

//计算全排列
var results [][]int
func calPermute(selects []int) {
	pathMap := make(map[int]int)
	var path []int
	permute(selects, path, pathMap)
	fmt.Println("len(results):", len(results))
	for _, res := range results {
		fmt.Printf("res:%+v\n", res)
	}
}