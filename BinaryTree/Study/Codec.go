package main

import "strconv"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	levelLoop2(root, 0)
	str := ""
	for _, v := range serializeString {
		if v == "" {
			str += "null,"
		} else {
			str += v + ","
		}
	}
	str = str[:len(str)-1]
	return "[" + str + "]"
}

var serializeString []string

func levelLoop2(node *TreeNode, index int) {
	if node == nil {
		return
	}
	depth := TopDown(node, 0)
	if serializeString == nil {
		serializeString = make([]string, 2*depth+1)
	}
	serializeString[index] = strconv.Itoa(node.Val)
	levelLoop2(node.Left, 2*index+1)
	levelLoop2(node.Right, 2*index+2)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
