package middle_test

import (
	"testing"
)

func Test_433(t *testing.T) {
	t.Log(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	t.Log(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}

func minMutation(start, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankSet := map[string]struct{}{}
	for _, s := range bank {
		bankSet[s] = struct{}{}
	}
	if _, ok := bankSet[end]; !ok {
		return -1
	}

	q1 := []string{start}
	q2 := []string{end}
	visited := map[string]bool{}
	visited[end] = true
	for step := 0; q1 != nil; step++ {
		tmp := q1
		q1 = nil
		for _, cur := range tmp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x {
						nxt := cur[:i] + string(y) + cur[i+1:]
						if _, ok := bankSet[nxt]; ok {
							//fmt.Println(nxt)
							if visited[nxt] {
								return step + 1
							}
							//delete(bankSet, nxt)
							visited[nxt] = true
							q1 = append(q1, nxt)
						}
					}
				}
			}
		}
		//fmt.Println("-------")
		q1, q2 = q2, q1
	}
	return -1
}

func minMutation0(start, end string, bank []string) int {
	if start == end {
		return 0
	}
	bankSet := map[string]struct{}{}
	for _, s := range bank {
		bankSet[s] = struct{}{}
	}
	if _, ok := bankSet[end]; !ok {
		return -1
	}

	q := []string{start}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x {
						nxt := cur[:i] + string(y) + cur[i+1:]
						if _, ok := bankSet[nxt]; ok {
							if nxt == end {
								return step + 1
							}
							delete(bankSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}
