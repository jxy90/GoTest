package hard_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	//fmt.Println(getString([]string{"This", "is", "an"}, 16))
	//fmt.Println(getString([]string{"example", "of", "text"}, 16))
	//fmt.Println(getString([]string{"justification."}, 16))
	//fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
	//fmt.Println(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	fmt.Println(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
}

func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	//记录当前数组范围
	left, right := 0, 0
	//记录字符长度
	count := 0
	ans := make([]string, 0)
	for right < n {
		word := words[right]
		//添加下一个字符的长度
		newCount := count + len(word)
		//如果不是以一个字符,加一个空格
		if left != right {
			newCount++
		}
		if newCount <= maxWidth {
			//范围右移,跟新count
			count = newCount
			right++
		} else {
			ans = append(ans, getString(words[left:right], maxWidth, false))
			left = right
			count = 0
		}
	}
	ans = append(ans, getString(words[left:right], maxWidth, true))
	return ans
}

func getString(words []string, maxWidth int, last bool) string {
	wordSpaceCount := len(words) - 1
	wordLength := 0
	for _, word := range words {
		wordLength += len(word)
	}
	spaceCount := maxWidth - wordLength
	x, y := 0, 0
	sb := strings.Builder{}
	if wordSpaceCount == 0 {
		sb.WriteString(words[0])
		for i := 0; i < spaceCount; i++ {
			sb.WriteString(" ")
		}
	} else {
		if last {
			for i := 0; i <= wordSpaceCount; i++ {
				sb.WriteString(words[i])
				if spaceCount > 0 {
					sb.WriteString(" ")
					spaceCount--
				}
			}
			for i := 0; i < spaceCount; i++ {
				sb.WriteString(" ")
			}
		} else {
			x, y = spaceCount/wordSpaceCount, spaceCount%wordSpaceCount
			for i := 0; i <= wordSpaceCount; i++ {
				sb.WriteString(words[i])
				for j := 0; j < x && wordSpaceCount != i; j++ {
					sb.WriteString(" ")
				}
				if y > 0 {
					sb.WriteString(" ")
					y--
				}
			}
		}
	}
	return sb.String()
}
