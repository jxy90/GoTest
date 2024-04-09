package classic

import (
	"fmt"
	"sync"
	"testing"
)

//循环打印数字和字母
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func TestNumLetter(t *testing.T) {
	printNumLetter()
}
func printNumLetter() {
	fmt.Println()
	wg := sync.WaitGroup{}
	wg.Add(1)
	letter, number := make(chan struct{}), make(chan struct{})

	go func() {
		i := 1
		for {
			select {
			case <-number:
				print(i)
				i++
				print(i)
				i++
				letter <- struct{}{}
			}
		}
	}()
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wg.Done()
					break
				}
				print(string(i))
				i++
				print(string(i))
				i++
				number <- struct{}{}
			}
		}
	}()

	number <- struct{}{}
	wg.Wait()
	println()
}
