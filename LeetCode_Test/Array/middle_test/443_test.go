package middle_test

import (
	"fmt"
	"testing"
)

func Test_compress(t *testing.T) {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	write, left := 0, 0
	for read, ch := range chars {
		if read == len(chars)-1 || ch != chars[read+1] {
			chars[write] = ch
			write++
			num := read - left + 1
			if num > 1 {
				anchor := write
				for ; num > 0; num /= 10 {
					chars[write] = '0' + byte(num%10)
					write++
				}
				s := chars[anchor:write]
				for i, n := 0, len(s); i < n/2; i++ {
					s[i], s[n-1-i] = s[n-1-i], s[i]
				}
			}
			left = read + 1
		}
	}
	return write
}
