package easy_test

func destCity(paths [][]string) string {
	ans := paths[0][1]
	cache := map[string]string{}
	for i := range paths {
		if ans == paths[i][0] {
			ans = paths[i][1]
			continue
		}
		cache[paths[i][0]] = paths[i][1]
	}
	for {
		if _, ok := cache[ans]; !ok {
			break
		}
		ans = cache[ans]
	}
	return ans
}
