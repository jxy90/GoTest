package hard_test

import (
	"fmt"
	"testing"
)

func Test_780(t *testing.T) {
	fmt.Println("1")
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > sx && ty > sy {
		if tx > ty {
			tx = tx % ty
		} else {
			ty = ty % tx
		}
	}
	if sx == tx && sy == ty {
		return true
	} else if sx == tx {
		return ty > sy && (ty-sy)%tx == 0
	} else if sy == ty {
		return tx > sx && (tx-sx)%ty == 0
	}
	return false
}
