package Backpack

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

/*描述
你总共有n 万元，希望申请国外的大学，要申请的话需要交一定的申请费用，给出每个大学的申请费用以及你得到这个大学offer的成功概率，大学的数量是 m。如果经济条件允许，你可以申请多所大学。找到获得至少一份工作的最高可能性。

0<=n<=10000,0<=m<=10000

样例 1:
	输入:
		n = 10
		prices = [4,4,5]
		probability = [0.1,0.2,0.3]
	输出:  0.440

	解释：
	选择第2和第3个学校。

样例 2:
	输入:
		n = 10
		prices = [4,5,6]
		probability = [0.1,0.2,0.3]
	输出:  0.370

	解释:
	选择第1和第3个学校。
*/

//

func Test_backpackIX(t *testing.T) {
	fmt.Println(backpackIX(10, []int{4, 4, 5}, []float64{0.1, 0.2, 0.3}))
}

func backpackIX(n int, prices []int, probability []float64) float64 {
	f := make([]float64, n+1)
	for i := range f {
		f[i] = 1
	}
	m := len(prices)
	for i := 0; i < m; i++ {
		for j := n; j >= prices[i]; j-- {
			f[j] = CommonUtil.MinFloat64(f[j], f[j-prices[i]]*(1-probability[i]))
		}
	}
	return 1 - f[n]
}
