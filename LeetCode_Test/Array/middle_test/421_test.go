package middle_test

import "github.com/jxy90/GoTest/Utils/CommonUtil"

func findMaximumXOR(nums []int) int {
	ans := 0
	root := ConstructorTire()
	for i := range nums {
		root.Add(i)
		j := root.GetVal(i)
		ans = CommonUtil.Max(ans, i^j)
	}

	return ans
}

type Tire struct {
	Tires [2]*Tire
}

func ConstructorTire() *Tire {
	return &Tire{}
}

func (root *Tire) Add(num int) {
	p := root
	for i := 31; i >= 0; i-- {
		temp := (num >> i) & 1
		if p.Tires[temp] == nil {
			p.Tires[temp] = &Tire{}
		}
		p = p.Tires[temp]
	}
}

func (root *Tire) GetVal(x int) int {
	ans := 0
	p := root

	for i := 31; i >= 0; i-- {
		a := (x >> i) & 1
		b := 1 - a
		if p.Tires[b] != nil {
			ans |= b << i
			p = p.Tires[b]
		} else {
			ans |= a << i
			p = p.Tires[a]
		}
	}
	return ans
}
