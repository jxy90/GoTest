package Chapter0_test

import (
	"fmt"
	"testing"
)

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool, len(deadends))
	for _, v := range deadends {
		dead[v] = true
	}
	visited := map[string]int{}
	//visited["0000"] = 0
	depth := 0
	queue := []string{"0000"}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[i]
			CurSlice := []rune(cur)
			//if _,ok:=visited[cur];ok {
			//	continue
			//}
			if target == cur {
				return depth
			}
			for j := range CurSlice {
				origin := CurSlice[j]
				CurSlice[j] = (CurSlice[j]-'0'+1)%10 + '0'
				temp := string(CurSlice)
				if _, ok := visited[temp]; !ok && !dead[temp] {
					queue = append(queue, temp)
					visited[temp] = visited[cur] + 1
				}
				CurSlice[j] = origin
				CurSlice[j] = (CurSlice[j]-'0'+9)%10 + '0'
				temp = string(CurSlice)
				if _, ok := visited[temp]; !ok && !dead[temp] {
					queue = append(queue, temp)
					visited[temp] = visited[cur] + 1
				}
				CurSlice[j] = origin
			}
		}
		queue = queue[length:]
		depth++
	}
	return -1
}

func queueContains(deadends []string, target string) bool {
	for _, b := range deadends {
		if b == target {
			return true
		}
	}
	return false
}

func openLockUp(str string, index int) string {
	ab := []rune(str)
	if ab[index] == '9' {
		ab[index] = '0'
	} else {
		ab[index] += 1
	}
	return string(ab)
}

func openLockDown(str string, index int) string {
	ab := []rune(str)
	if ab[index] == '0' {
		ab[index] = '9'
	} else {
		ab[index] -= 1
	}
	return string(ab)
}

func openLockBFSD(deadends []string, target string) int {
	queueHead := []string{}
	queueHead = append(queueHead, "0000")
	queueTail := []string{}
	queueTail = append(queueTail, target)
	visited := []string{}
	//visited = append(visited, "0000")
	//visited = append(visited, target)
	depth := 0
	for len(queueHead) > 0 && len(queueTail) > 0 {
		if len(queueHead) > len(queueTail) {
			queueHead, queueTail = queueTail, queueHead
		}
		length := len(queueHead)
		for i := 0; i < length; i++ {
			cur := queueHead[i]
			if queueContains(queueTail, cur) {
				return depth
			} else if queueContains(deadends, cur) {
				continue
			}
			visited = append(visited, cur)
			for j := 0; j < len(cur); j++ {
				up := openLockUp(cur, j)
				if !queueContains(visited, up) {
					queueHead = append(queueHead, up)
					//visited = append(visited, up)
				}
				down := openLockDown(cur, j)
				if !queueContains(visited, down) {
					queueHead = append(queueHead, down)
					//visited = append(visited, down)
				}
			}
		}
		queueHead = queueHead[length:]
		depth++
	}
	return -1
}

func Test_openLock(t *testing.T) {
	//fmt.Println(openLockUp("0000", 2))
	//fmt.Println(openLockDown("0000", 2))

	deadlines := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0000"
	fmt.Println(openLock(deadlines, target))
	fmt.Println(openLockBFSD(deadlines, target))
}
