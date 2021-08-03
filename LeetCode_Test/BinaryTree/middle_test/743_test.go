package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"testing"
)

func Test_networkDelayTime(t *testing.T) {
	//println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	//println(networkDelayTime([][]int{{1, 2, 1}, {2, 1, 3}}, 2, 2))
	//{{4,2,76},{1,3,79},{3,1,81},{4,3,30},{2,1,47},{1,5,61},{1,4,99},{3,4,68},{3,5,46},{4,1,6},{5,4,7},{5,3,44},{4,5,19},{2,3,13},{3,2,18},{1,2,0},{5,1,25},{2,5,58},{2,4,77},{5,2,74}}
	println(networkDelayTime([][]int{{4, 2, 76}, {1, 3, 79}, {3, 1, 81}, {4, 3, 30}, {2, 1, 47}, {1, 5, 61}, {1, 4, 99}, {3, 4, 68}, {3, 5, 46}, {4, 1, 6}, {5, 4, 7}, {5, 3, 44}, {4, 5, 19}, {2, 3, 13}, {3, 2, 18}, {1, 2, 0}, {5, 1, 25}, {2, 5, 58}, {2, 4, 77}, {5, 2, 74}}, 5, 3))

}

func networkDelayTime(times [][]int, n int, k int) int {
	//N, M := 1001, 6001
	w := make([][]int, n)
	for i := range w {
		w[i] = make([]int, n)
		for j := range w[i] {
			if i == j {
				w[i][j] = 0
			} else {
				w[i][j] = math.MaxInt32
			}
		}
	}
	for i := range times {
		source, target, weight := times[i][0]-1, times[i][1]-1, times[i][2]
		w[source][target] = weight
	}
	for p := 0; p < n; p++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				w[i][j] = CommonUtil.Min(w[i][j], w[i][p]+w[p][j])
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = CommonUtil.Max(ans, w[k-1][i])
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
