package hard_test

import (
	"fmt"
	"sort"
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

}
