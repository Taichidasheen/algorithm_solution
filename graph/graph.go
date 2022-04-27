package graph

/*
leetcode-797:
https://leetcode.cn/problems/all-paths-from-source-to-target/
 */
//在for循环外部加入path
var res [][]int
func allPathsSourceTarget(graph [][]int) [][]int {
	var path []int
	res = [][]int{}//这里为了通过leetcode的多用例，需要清空res
	traverse(graph, path, 0)
	return res
}


func traverse(graph [][]int, path []int, n int) {
	path = append(path, n)
	length := len(graph)
	if n == length - 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)//注意：这里一定要copy一份再append，不然之前append的内容会被修改掉（因为只是append的一个引用，底层数组是共享的）
		path = path[0:len(path)-1]
		return
	}
	childs := graph[n]
	for _, node := range childs {
		traverse(graph, path, node)
	}
	path = path[0:len(path)-1]

}


//在for循环内部加入path
func allPathsSourceTargetV2(graph [][]int) [][]int {
	return nil
}


