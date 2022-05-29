package middle_test

import (
	"strconv"
	"strings"
	"testing"
)

func Test_468(t *testing.T) {
	//t.Log(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	t.Log(validIPAddress("256.256.256.256"))
}

func validIPAddress(queryIP string) string {
	IPv4, IPv6, Neither := "IPv4", "IPv6", "Neither"
	queryIPs := strings.Split(queryIP, ".")
	if len(queryIPs) == 1 {
		queryIPs = strings.Split(queryIP, ":")
	}
	n := len(queryIPs)
	if n == 4 {
		for i := range queryIPs {
			if len(queryIPs[i]) > 1 && queryIPs[i][0] == '0' {
				return Neither
			}
			x := -1
			x, err := strconv.Atoi(queryIPs[i])
			if err != nil || x < 0 || x > 255 {
				return Neither
			}
		}
		return IPv4
	} else if n == 8 {
		for i := range queryIPs {
			if len(queryIPs[i]) > 4 {
				return Neither
			}
			_, err := strconv.ParseUint(queryIPs[i], 16, 64)
			if err != nil {
				return Neither
			}
		}
		return IPv6
	}
	return Neither
}
