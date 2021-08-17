package easy_test

import (
	"testing"
)

func Test_checkRecord(t *testing.T) {
	println(checkRecord("PPALLP"))
	println(checkRecord("PPALLL"))
}

func checkRecord(s string) bool {
	absent, late := 0, 0
	for _, v := range s {
		if v == 'A' {
			absent++
			late = 0
		} else if v == 'L' {
			late++
		} else if v == 'P' {
			late = 0
		}
		if absent > 1 || late > 2 {
			return false
		}
	}
	return true
}
