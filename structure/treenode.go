package structure

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenTree(x []int) *TreeNode {
	index := 0
	return GenTreeHelp(x, &index)
}

func GenTreeHelp(x []int, index *int) *TreeNode {
	if len(x) == *index {
		return nil
	}
	fmt.Println(x, x[*index])
	if x[*index] == -1 {
		*index = *index + 1
		return nil
	}
	//fmt.Print(*index, x[*index])
	(*index) = *index + 1
	node := &TreeNode{Val: x[*index]}
	node.Left = GenTreeHelp(x, index)
	return node
}

//层序遍历
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue)> 0 {
		n := len(queue)
		var level []int
		for i:=0;i<n;i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		res = append(res, level)
	}
	return res
}

//Z字形层序遍历
func ZigLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	deep := 1
	for len(queue)> 0 {
		n := len(queue)
		var level []int
		for i:=0;i<n;i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		if deep % 2 == 0 {
			level = reverseArr(level)
		}
		res = append(res, level)
		deep ++
	}
	return res
}

func reverseArr(input []int) []int {
	var output []int
	len := len(input)
	for i:= len-1;i>=0;i-- {
		output = append(output, input[i])
	}
	return output
}

//	序列化
func dumpTreeToString(tree *TreeNode) string {
	var str string = ""
	dumpTreeToStringHelper(tree, &str)
	return str
}

func dumpTreeToStringHelper(tree *TreeNode, str *string) {
	if tree == nil {
		*str += "#"
		return
	}
	*str += strconv.Itoa(tree.Val)
	dumpTreeToStringHelper(tree.Left, str)
	dumpTreeToStringHelper(tree.Right, str)
}

//	反序列化
func conFromPreStr(str string) *TreeNode {
	var strIndex int = 0
	return conFromPreStrHelper(str, &strIndex)
}

func conFromPreStrHelper(str string, index *int) *TreeNode {
	if len(str) == *index {
		return nil
	}
	if (str)[*index] == '#' {
		(*index) = *index + 1
		return nil
	}
	v, _ := strconv.Atoi(string((str)[*index]))
	root := &TreeNode{Val: v}
	(*index) = *index + 1
	root.Left = conFromPreStrHelper(str, index)
	root.Right = conFromPreStrHelper(str, index)

	return root
}
