package middle_test

import "testing"

func Test_openLock(t *testing.T) {
	deadlines := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0000"
	deadlines = []string{"0201", "0101", "0102", "1212", "2002"}
	target = "0202"
	deadlines = []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	target = "8888"
	//["8887","8889","8878","8898","8788","8988","7888","9888"]
	println(openLock(deadlines, target))
}

func openLock(deadends []string, target string) int {
	start := "0000"
	dead := map[string]bool{}
	for i := range deadends {
		dead[deadends[i]] = true
	}
	if dead[start] {
		return -1
	}
	depth := 0
	options := []rune{1, 9}

	visited := map[string]int{}
	visited[start] = 0
	queue := make([]string, 0)
	queue = append(queue, start)

	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			now := queue[i]
			nowSlice := []rune(now)
			if now == target {
				return depth
			}
			for j := 0; j < len(nowSlice); j++ {
				for k := 0; k < 2; k++ {
					option := options[k]
					origin := nowSlice[j]
					nowSlice[j] = (nowSlice[j]-'0'+option)%10 + '0'
					temp := string(nowSlice)
					if _, ok := visited[temp]; !ok && !dead[temp] {
						queue = append(queue, temp)
						visited[temp] = visited[now] + 1
					}
					nowSlice[j] = origin
				}
			}
		}
		queue = queue[length:]
		depth++
	}
	return -1
}
