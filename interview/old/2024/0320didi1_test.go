package _024

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func Test_024(t *testing.T) {
	main1()

	// 1.0 > 0.1
	// 1.1 > 1.0
	// 1.01 == 1.1
	// 1.0.0 == 1.0
	//println(compareVersion("1.0", "0.1"))
	//println(compareVersion("1.1", "1.0"))
	//println(compareVersion("1.01", "1.1"))
	//println(compareVersion("1.0.0", "1.0"))
	//println(compareVersion("0.1", "1.0"))
}

// 要求：
// 1. 只能编辑 foo 函数
// 2. foo 必须要调用 slow 函数
// 3. foo 函数在 ctx 超时后必须立刻返回
// 4. 【加分项】如果 slow 结束的比 ctx 快，也立刻返回
func main1() {
	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	foo(ctx)
}

// 您需要实现 foo 函数，要求 foo 在 ctx 超时后立即返回
// foo 必须调用 slow 函数
func foo(ctx context.Context) {
	ch := make(chan struct{})
	go func() {
		slow()
		close(ch)
		ch = nil
	}()
	select {
	case <-ctx.Done():
		fmt.Println("time out")
		return
	case v, ok := <-ch:
		fmt.Println("slow", v, ok)
		return
	}
}

// 您不能修改 slow 函数
func slow() {
	n := rand.Intn(3)
	fmt.Printf("sleep %ds\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}

// 我们有一张学生的分数表 scores:
// name, subject, score
// 张三 语文   81
// 张三 数学   75
// 李四 语文   85
// 李四 数学   90

// 要求：用一条 SQL 查出每门课都大于 80 分的学生的姓名

//select name form scores group by name having min(score)>80;

// 给你两个版本号 `version1` 和 `version2` ，请你比较它们。
// 版本号由多个部分组成，每个部分由一个 '.' 连接。每个部分都由数字组成，可能包含前导零。
// 返回规则如下：
// 如果 version1 > version2 返回 1，
// 如果 version1 < version2 返回 -1，
// 除此之外返回 0

// 示例：
// 1.0 > 0.1
// 1.1 > 1.0
// 1.01 == 1.1
// 1.0.0 == 1.0

func compareVersion(v1, v2 string) int {
	num1 := versionToNums(v1)
	num2 := versionToNums(v2)
	//双指针比较
	p1, p2 := 0, 0
	for p1 < len(num1) && p2 < len(num2) {
		val1, val2 := num1[p1], num2[p2]
		if val1 > val2 {
			return 1
		} else if val1 < val2 {
			return -1
		}
		p1++
		p2++
	}

	//剩余元素判断
	if p1 != len(num1) {
		for p1 < len(num1) {
			if num1[p1] != 0 {
				return 1
			}
			p1++
		}
	}
	if p2 != len(num2) {
		for p2 < len(num2) {
			if num2[p2] != 0 {
				return -1
			}
			p2++
		}
	}
	return 0
}

func versionToNums(version string) []int {
	strs := strings.Split(version, ".")
	nums := make([]int, len(strs))
	for i, str := range strs {
		for _, b := range str {
			nums[i] = nums[i]*10 + int(b-'0')
		}
	}
	return nums
}
