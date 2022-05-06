package offer

/*
leetcode-剑指 Offer II 116
https://leetcode.cn/problems/bLyHh0/
 */


func findCircleNum(isConnected [][]int) int {
	parent := make(map[int]int)
	for i:=0;i<len(isConnected);i++ {
		parent[i] = i
	}
	numOfProvince := len(isConnected)

	for i:=0;i<len(isConnected);i++ {
		for j:=i+1;j<len(isConnected);j++ {
			pi := find(i,parent)
			pj := find(j,parent)
			if pi == pj {
				continue
			} else {
				if isConnected[i][j] == 1 {
					//注意：查并集的精髓是把一个节点的代表节点（父节点）的父节点置为另一个节点的代表节点（父节点）
					parent[pi] = pj
					numOfProvince --
				}
			}
		}
	}
	return numOfProvince
}

func find(i int, parent map[int]int) int {
	if parent[i] != i {
		return find(parent[i], parent)
	}
	return parent[i]
}