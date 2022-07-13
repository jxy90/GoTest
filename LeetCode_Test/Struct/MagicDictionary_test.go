package Struct

type MagicDictionary struct {
	lengthMap map[int][]string
}

func MagicDictionaryConstructor() MagicDictionary {
	return MagicDictionary{
		lengthMap: map[int][]string{},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, s := range dictionary {
		this.lengthMap[len(s)] = append(this.lengthMap[len(s)], s)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	dics := this.lengthMap[len(searchWord)]
	for _, dic := range dics {
		count := 0
		for i := range dic {
			if dic[i] != searchWord[i] {
				count++
			}
			if count > 1 {
				break
			}
		}
		if count == 1 {
			return true
		}
	}
	return false
}
