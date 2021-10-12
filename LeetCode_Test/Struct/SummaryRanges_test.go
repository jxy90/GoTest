package Struct

import (
	"fmt"
	"testing"
)

func Test_SummaryRanges(t *testing.T) {
	obj := ConstructorSummaryRanges()
	obj.AddNum(1)
	obj.AddNum(9)
	obj.AddNum(2)
	fmt.Println(obj.GetIntervals())
}

type SummaryRanges struct {
	ranges [][]int
}

func ConstructorSummaryRanges() SummaryRanges {
	return SummaryRanges{ranges: [][]int{}}
}

func (this *SummaryRanges) AddNum(val int) {
	for i, ints := range this.ranges {
		a, b := ints[0], ints[1]
		if a <= val && val <= b {
			return
		} else if b+1 == val {
			ints[1] = val
			//判断需不需要和后一个合并
			if i < len(this.ranges)-1 && this.ranges[i+1][0]-1 == this.ranges[i][1] {
				this.ranges[i+1][0] = a
				//移除当前项
				temp := append(this.ranges[:i], this.ranges[i+1:]...)
				this.ranges = temp
			}
			return
		} else if a-1 == val {
			ints[0] = val
			//判断需不需要和前一个合并
			if i > 0 && this.ranges[i-1][1]+1 == this.ranges[i][0] {
				this.ranges[i-1][1] = b
				//移除当前项
				temp := append(this.ranges[:i], this.ranges[i+1:]...)
				this.ranges = temp
			}
			return
		} else if a > val {
			//val 插入 i 到的前面
			this.ranges = append(this.ranges[:i+1], this.ranges[i:]...)
			this.ranges[i] = []int{val, val}
			//this.ranges = temp
			return
		}
	}
	this.ranges = append(this.ranges, []int{val, val})
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.ranges
}
