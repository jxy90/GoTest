package Codec

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(root *TreeNode) string {
	//sb := strings.Builder{}
	if root == nil {
		//sb.WriteString("#")
		//sb.WriteString(" ")
		return "#"
	} else {
		//Value := strconv.Itoa(root.Val)
		//sb.WriteString(Value)
		//sb.WriteString(" ")
		//LeftValue := Serialize(root.Left)
		//RightValue := Serialize(root.Right)
		//sb.WriteString(LeftValue)
		//sb.WriteString(RightValue)

	}
	//return sb.String()
	return strconv.Itoa(root.Val) + " " + Serialize(root.Left) + " " + Serialize(root.Right)

}

var index = -1

func Deserialize(str string) *TreeNode {
	strSlice := strings.Split(str, " ")
	root := loop(strSlice)
	index = -1
	return root
}

func loop(strSlice []string) *TreeNode {
	index++
	if strSlice[index] == "#" || index > len(strSlice)-1 {
		return nil
	}
	inValue, _ := strconv.Atoi(strSlice[index])
	root := &TreeNode{
		Val:   inValue,
		Left:  loop(strSlice),
		Right: loop(strSlice),
	}
	return root
}
