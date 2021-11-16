package hard_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func Test_findMinStep(t *testing.T) {
	fmt.Println(findMinStep("WRRBBW", "RB"))
}

func findMinStep(board string, hand string) int {
	handByte := []byte(hand)
	sort.Slice(handByte, func(i, j int) bool {
		return handByte[i] > handByte[j]
	})

	//用队列维护的状态队列:其中的三个元素分别为桌面球状态,手中球状态 和 回合数
	queue := make([]State, 0)
	queue = append(queue, State{board: board, hand: hand, step: 0})

	//记录访问过的状态
	visited := map[string]bool{}
	visited[hand+" "+board] = true

	for len(queue) > 0 {
		curState := queue[0]
		for i := 0; i <= len(curState.board); i++ {
			for j := 0; j < len(curState.hand); j++ {
				//当前球的颜色和上一个相同
				if j > 0 && curState.hand[j] == curState.hand[j-1] {
					continue
				}
				//只在连续相同颜色求的开头插入
				if i > 0 && curState.board[i-1] == curState.hand[j] {
					continue
				}
				//只在一下两种情况,放置新球
				choose := false
				//1 当前球的颜色和手中球颜色相同
				if i < len(curState.board) && curState.board[i] == curState.hand[j] {
					choose = true
				}
				//2 前后颜色相同和手中颜色不同
				if i > 0 && i < len(curState.board) && curState.board[i-1] == curState.board[i] && curState.board[i-1] != curState.hand[j] {
					choose = true
				}

				if choose {
					newBoard := findMinStepHelper(curState.board[:i] + curState.hand[j:j+1] + curState.board[i:])
					newHand := curState.hand[:j] + curState.hand[j+1:]
					if len(newBoard) == 0 {
						return curState.step + 1
					}
					str := newBoard + " " + newHand
					if _, ok := visited[str]; !ok {
						visited[str] = true
						queue = append(queue, State{
							board: newBoard,
							hand:  newHand,
							step:  curState.step + 1,
						})
					}
				}
			}
		}

	}
	return -1
}

func findMinStepHelper(s string) string {
	sb := strings.Builder{}
	letterStack := make([]byte, 0)
	countStack := make([]int, 0)

	for i := 0; i < len(s); i++ {
		c := s[i]
		for len(letterStack) > 0 && c != letterStack[len(letterStack)-1] && countStack[len(countStack)-1] >= 3 {
			letterStack = letterStack[:len(letterStack)-1]
			countStack = countStack[:len(countStack)-1]
		}
		if len(letterStack) == 0 || c != letterStack[len(letterStack)-1] {
			letterStack = append(letterStack, c)
			countStack = append(countStack, 1)
		} else {
			countStack[len(countStack)-1] += 1
		}
	}
	if len(countStack) > 0 && countStack[len(countStack)-1] >= 3 {
		letterStack = letterStack[:len(letterStack)-1]
		countStack = countStack[:len(countStack)-1]
	}
	for i := 0; i < len(letterStack); i++ {
		for j := 0; j < countStack[i]; j++ {
			sb.WriteByte(letterStack[i])
		}
	}
	return sb.String()
}

type State struct {
	board string
	hand  string
	step  int
}
