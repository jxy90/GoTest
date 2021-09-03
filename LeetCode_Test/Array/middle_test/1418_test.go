package middle_test

import (
	"sort"
	"strconv"
	"testing"
)

func Test_displayTable(t *testing.T) {
	fmt.Println(displayTable([][]string{{"David", "3", "Ceviche"}, {"Corina", "10", "Beef Burrito"}, {"David", "3", "Fried Chicken"}, {"Carla", "5", "Water"}, {"Carla", "5", "Ceviche"}, {"Rous", "3", "Ceviche"}}))
}

func displayTable(orders [][]string) [][]string {
	//桌子食物,数量
	tableFoods := map[string]int{}
	//桌子数量
	tables := map[string]bool{}
	//食物种类
	foods := map[string]bool{}
	for i := range orders {
		order := orders[i]
		tableFoodName := order[1] + order[2]
		tableFoods[tableFoodName] += 1
		tables[order[1]] = true
		foods[order[2]] = true
	}
	//桌子排序
	tableSorts := make([]int, 0, len(tables))
	for s := range tables {
		table, _ := strconv.Atoi(s)
		tableSorts = append(tableSorts, table)
	}
	sort.Ints(tableSorts)
	//食物排序
	foodSorts := make([]string, 0, len(foods))
	for s := range foods {
		foodSorts = append(foodSorts, s)
	}
	sort.Strings(foodSorts)
	//构建表格
	ans := make([][]string, len(tables)+1)
	//for i := range ans {
	//	ans[i] = make([]string, len(foods)+1)
	//}
	ans[0] = make([]string, len(foods)+1)
	ans[0][0] = "Table"
	for i, foodName := range foodSorts {
		ans[0][i+1] = foodName
	}
	for i, tableName := range tableSorts {
		ans[i+1] = make([]string, len(foods)+1)
		ans[i+1][0] = strconv.Itoa(tableName)
		for j, foodName := range foodSorts {
			ans[i+1][j+1] = strconv.Itoa(tableFoods[strconv.Itoa(tableName)+foodName])
		}
	}
	return ans
}
