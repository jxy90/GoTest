package Chapter1

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"math"
	"strconv"
	"testing"
)

//经典动态规划：高楼扔鸡蛋
//若干层楼，若干个鸡蛋，让你算出最少的尝试次数，找到鸡蛋恰好摔不碎的那层楼。

//当前状态是K个鸡蛋，N层楼
//返回这个状态下最优的结果
/*
def dp(K,N)
	if K==1 return N
	if N=0 return 0

	int res
	for 1<=i<=N
		res = min(res,
				  max(
						dp(K-1,i-1), 	//碎了
						dp(K,N-i)		//没碎
					 )+1 #在第i层扔
				 )
	return res

def superEggDrop(K: int, N: int):

    memo = dict()
    def dp(K, N) -> int:
        # base case
        if K == 1: return N
        if N == 0: return 0
        # 避免重复计算
        if (K, N) in memo:
            return memo[(K, N)]

        res = float('INF')
        # 穷举所有可能的选择
        for i in range(1, N + 1):
            res = min(res,
                      max(
                            dp(K, N - i),
                            dp(K - 1, i - 1)
                         ) + 1
                  )
        # 记入备忘录
        memo[(K, N)] = res
        return res

    return dp(K, N)
*/
var memo [][]int

func Test_superEggDrop(t *testing.T) {
	res := superEggDrop(2, 6)
	println("结果：" + strconv.Itoa(res))
	res2 := superEggDropDp2(2, 6)
	println("结果：" + strconv.Itoa(res2))
}

func superEggDrop(k, n int) int {
	memo = make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	return superEggDropDP(k, n)
}

func superEggDropDP(K, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	//避免重复计算
	if memo[K][N] != 0 {
		return memo[K][N]
	}
	res := math.MaxInt32
	//穷举所有可能的选择
	for i := 1; i <= N; i++ {
		res = CommonUtil.Min(res,
			CommonUtil.Max(superEggDropDP(K-1, i-1), superEggDropDP(K, N-i))+1)
	}
	//记入备忘录
	memo[K][N] = res
	return res
}

func superEggDropDp2(K, N int) int {
	// m 最多不会超过 N 次（线性扫描）
	//int[][] dp = new int[K + 1][N + 1];
	dp := make([][]int, K+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	// base case:
	// dp[0][..] = 0
	// dp[..][0] = 0
	// Java 默认初始化数组都为 0
	m := 0
	for dp[K][m] < N {
		m++
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
		}
	}
	return m
}
