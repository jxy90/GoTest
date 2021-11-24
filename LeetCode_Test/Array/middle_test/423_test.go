package middle_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_originalDigits(t *testing.T) {
	fmt.Println(originalDigits("owoztneoer"))
}

func originalDigits(s string) string {
	c := map[rune]int{}
	for _, ch := range s {
		c[ch]++
	}

	cnt := [10]int{}
	cnt[0] = c['z']
	cnt[2] = c['w']
	cnt[4] = c['u']
	cnt[6] = c['x']
	cnt[8] = c['g']

	cnt[3] = c['h'] - cnt[8]
	cnt[5] = c['f'] - cnt[4]
	cnt[7] = c['s'] - cnt[6]

	cnt[1] = c['o'] - cnt[0] - cnt[2] - cnt[4]

	cnt[9] = c['i'] - cnt[5] - cnt[6] - cnt[8]

	sb := strings.Builder{}
	for i := range cnt {
		for cnt[i] > 0 {
			sb.WriteString(strconv.Itoa(i))
			cnt[i]--
		}
	}
	return sb.String()
}
