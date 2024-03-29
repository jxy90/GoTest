package unionFindDisjointSets

import (
	"fmt"
	"testing"
)

func Test_findRedundantConnection(t *testing.T) {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
}

func findRedundantConnection(edges [][]int) []int {
	f := make([]int, 1001)
	for i := range f {
		f[i] = i
	}
	for i := range edges {
		if find684(&f, edges[i][0]) == find684(&f, edges[i][1]) {
			return edges[i]
		} else {
			join684(&f, edges[i][0], edges[i][1])
		}
	}
	return nil
}

func find684(f *[]int, value int) int {
	root := (*f)[value]
	for root != (*f)[root] {
		root = (*f)[root]
	}
	//路径压缩
	i := value
	for (*f)[i] != root {
		temp := (*f)[i]
		(*f)[i] = root
		i = (*f)[temp]
	}
	return root
}

func join684(f *[]int, x, y int) {
	xp := find684(f, x)
	yp := find684(f, y)
	if xp != yp {
		(*f)[xp] = yp
	}
}
