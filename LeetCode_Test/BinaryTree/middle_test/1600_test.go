package middle_test

import (
	"fmt"
	"testing"
)

func Test_1600(t *testing.T) {

	//Your ThroneInheritance object will be instantiated and called as such:
	c := Constructor1600("king")    // 继承顺序：king
	c.Birth("king", "andy")         // 继承顺序：king > andy
	c.Birth("king", "bob")          // 继承顺序：king > andy > bob
	c.Birth("king", "catherine")    // 继承顺序：king > andy > bob > catherine
	c.Birth("andy", "matthew")      // 继承顺序：king > andy > matthew > bob > catherine
	c.Birth("bob", "alex")          // 继承顺序：king > andy > matthew > bob > alex > catherine
	c.Birth("bob", "asha")          // 继承顺序：king > andy > matthew > bob > alex > asha > catherine
	data := c.GetInheritanceOrder() // 返回 ["king", "andy", "matthew", "bob", "alex", "asha", "catherine"]
	fmt.Println(data)
	c.Death("bob")                   // 继承顺序：king > andy > matthew > bob（已经去世）> alex > asha > catherine
	data2 := c.GetInheritanceOrder() // 返回 ["king", "andy", "matthew", "alex", "asha", "catherine"]
	fmt.Println(data2)

}

type PersonTreeNode struct {
	Name  string
	Alive bool
	Child []*PersonTreeNode
}

type ThroneInheritance struct {
	person map[string]*PersonTreeNode
	tree   *PersonTreeNode
}

func Constructor1600(kingName string) ThroneInheritance {
	king := &PersonTreeNode{
		Name:  kingName,
		Alive: true,
	}
	person := make(map[string]*PersonTreeNode)
	person[kingName] = king
	return ThroneInheritance{
		person: person,
		tree:   king,
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	if this.person[parentName] != nil {
		if this.person[parentName].Child == nil {
			this.person[parentName].Child = make([]*PersonTreeNode, 0)
		}
		child := &PersonTreeNode{
			Name:  childName,
			Alive: true,
		}
		this.person[parentName].Child = append(this.person[parentName].Child, child)
		this.person[childName] = child
	}
}

func (this *ThroneInheritance) Death(name string) {
	if this.person[name] != nil {
		this.person[name].Alive = false
	}
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	ans := make([]string, 0)
	var Successor func(string)
	Successor = func(x string) {
		//根左右
		//if this.person[x] != nil {
		if this.person[x].Alive {
			ans = append(ans, x)
		}
		for i := range this.person[x].Child {
			Successor(this.person[x].Child[i].Name)
		}
		//}
	}
	Successor(this.tree.Name)
	return ans
}
