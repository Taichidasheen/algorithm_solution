package graph

import (
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	//[[4,3,1],[3,2,4],[3],[4],[]]
	/*graph := [][]int{
		[]int{4,3,1},
		[]int{3,2,4},
		[]int{3},
		[]int{4},
		[]int{},
	}*/
	//[[3,1],[4,6,7,2,5],[4,6,3],[6,4],[7,6,5],[6],[7],[]]
	graph := [][]int{
		[]int{3,1},
		[]int{4,6,7,2,5},
		[]int{4,6,3},
		[]int{6,4},
		[]int{7,6,5},
		[]int{6},
		[]int{7},
		[]int{},
	}

	allPathsSourceTarget(graph)

}