package middle_test

import (
	"math/rand"
	"sort"
	"testing"
)

func Test_PickIndex(t *testing.T) {
	//println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

type Solution struct {
	w []int
}

func ConstructorPickIndex(w []int) Solution {
	n := len(w)
	for i := 1; i < n; i++ {
		w[i] += w[i-1]
	}
	return Solution{w}
}

func (this *Solution) PickIndex() int {
	index := rand.Intn(this.w[len(this.w)]) + 1
	return sort.SearchInts(this.w, index)
}
