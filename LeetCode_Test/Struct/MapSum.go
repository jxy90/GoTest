package Struct

type MapSum struct {
	cache map[string]int
	array []string
}

func ConstructorMapSum() MapSum {
	return MapSum{
		cache: map[string]int{},
		array: make([]string, 0),
	}
}

func (this *MapSum) Insert(key string, val int) {
	if _, ok := this.cache[key]; ok {
		this.cache[key] = val
		return
	}
	this.cache[key] = val
	this.array = append(this.array, key)
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for _, s := range this.array {
		n := len(prefix)
		if len(s) < n {
			continue
		}
		isSame := true
		for i := 0; i < n; i++ {
			if s[i] != prefix[i] {
				isSame = false
			}
		}
		if !isSame {
			continue
		}
		sum += this.cache[s]
	}
	return sum
}
