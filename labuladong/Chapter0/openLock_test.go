package Chapter0_test

func openLock(deadends []string, target string) int {
	queue := []string
	queue = append(queue, "0000")
	depth := 0
	for len(queue) > 0 {
		length := len(queue)
		for j := 0; j < length; j++ {
			if queue[j] == target {
				return depth
			} else if queue[j] in deadends{
				continue
			}
			for i := 0; i < len(v); i++ {
				queue = append(queue, openLockUp(queue[j], i))
				queue = append(queue, openLockDown(queue[j], i))
			}
			depth++
		}
	}
	return depth
}

func openLockUp(str string, index int) string {
	temp := string(str[index])
	if temp == "9" {
		str[index] = "0"
	} else {
		str[index] -= 1
	}
	return str
}

func openLockDown(str string, index int) string {
	temp := string(str[index])
	if temp == "0" {
		str[index] = '9'
	} else {
		str[index] += 1
	}
	return str
}
