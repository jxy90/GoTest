package easy_test

import "testing"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func Test_getImportance(t *testing.T) {
	employees := make([]*Employee, 0)
	employees = append(employees, &Employee{
		Id:           1,
		Importance:   2,
		Subordinates: []int{2},
	})
	employees = append(employees, &Employee{
		Id:         2,
		Importance: 3,
	})
	println(getImportance(employees, 1))
}

func getImportance(employees []*Employee, id int) int {
	hash := map[int]*Employee{}
	for _, employee := range employees {
		hash[employee.Id] = employee
	}
	ans := 0
	var getImportanceHelper func(id int)
	getImportanceHelper = func(id int) {
		employee := hash[id]
		ans += employee.Importance
		for _, subordinate := range employee.Subordinates {
			getImportanceHelper(subordinate)
		}
	}
	getImportanceHelper(id)
	return ans
}

func getImportance1(employees []*Employee, id int) int {
	ans := 0
	for _, employee := range employees {
		if employee.Id == id {
			ans += employee.Importance
			for _, subordinate := range employee.Subordinates {
				ans += getImportance(employees, subordinate)
			}
		}
	}
	return ans
}
