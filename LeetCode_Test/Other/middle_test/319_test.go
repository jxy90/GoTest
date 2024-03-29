package middle_test

import (
	"fmt"
	"math"
	"testing"
)

func Test_bulbSwitch(t *testing.T) {
	fmt.Println(bulbSwitch(3))
	fmt.Println(bulbSwitch(2))
	fmt.Println(bulbSwitch(1))
	fmt.Println(bulbSwitch(0))
}

//首先思考：灯泡i会被按几次？这其实相当于求i有几个因子 比如灯泡8,一共会被按4次，分别是第一轮 第二轮 第四轮 第八轮
//一开始灯都是灭的，所以如果i有k个因子，且k为奇数，那么最终灯就亮，如果为偶数，灯就灭。 那么问题就转化成了，求1..n每个数分别有几个因子。 最直观的做法就是，枚举i，然后计算i的因子数。怎么计算呢？ 最直观的做法就是枚举j=0..i，count+= i%j==0 ? 1:0; 这里有个优化的点，假如x*y=z，显然z%y==0且z%x==0。 也就是说你只要需要枚举j从到1..根号i，count += i%j==0? 2:0;
func bulbSwitch0(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		temp := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				temp++
			}
		}
		if temp%2 == 1 {
			ans++
		}
	}
	return ans
}
func bulbSwitch1(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		temp := 0
		sqrti := int(math.Sqrt(float64(i)))
		for j := 1; j <= sqrti; j++ {
			if i%j == 0 {
				temp++
				if i/j != j {
					temp++
				}
			}
		}
		if temp%2 == 1 {
			ans++
		}
	}
	return ans
}

//在写完上面的式子之后你突然可以发现，在根号i的左边每发现一个j使得i%j==0，那么根号i的右边一定存在一个k同样满足i%k==0，一次枚举会把count+2。而我们关心的其实是最终count为奇数还是偶数！！通过枚举，count最终的结果都是偶数，当且仅当 i可以被开根号时，count才会是奇数！
//然后其实问题转换成了，求数字1..n中有几个数能开更开得尽（结果是整数）
func bulbSwitch2(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		//奇数
		isOdd := false
		sqrti := int(math.Sqrt(float64(i)))
		for j := 1; j <= sqrti; j++ {
			if i%j == 0 && i/j == j {
				isOdd = true
			}
		}
		if isOdd {
			ans++
		}
	}
	return ans
}
func bulbSwitch3(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		//奇数
		sqrti := math.Sqrt(float64(i))
		if sqrti-float64(int(sqrti)) == 0 {
			ans++
		}
	}
	return ans
}

//想到这里，你就可以一个O(n)的枚举来完成了吗？ 其实还可以优化！
//
//想想，在1..n中，假设n等于100，1*1 2*2 3*3 4*4 ... 10*10的结果都小于等于100，换句话说，1*1 ... 根号n*根号n 都<= n，所以求1..n里有几个开根号能开尽的数，其实就是求根号n向下取整等于几。
func bulbSwitch(n int) int {
	sqrti := int(math.Sqrt(float64(n)))
	return sqrti
}

//还能优化不？ 当然还可以！
//
//求n开根号怎么求，有大量的论文可以参考，它们收敛速度能比语言自带的sqrt还快！！
