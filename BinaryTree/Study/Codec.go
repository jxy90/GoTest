package main

import (
	"strconv"
	"strings"
)

type Codec struct {
	serializeString []string
	deserializeList []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.serializeString = []string{}
	depth := TopDown(root, 0)
	if this.serializeString == nil {
		this.serializeString = make([]string, 2*depth+1)
	}
	this.serializeLoop(root, 0)
	str := ""
	for _, v := range this.serializeString {
		if v == "" {
			str += "null,"
		} else {
			str += v + ","
		}
	}
	if str == "" {
		str += "null,"
	}
	str = str[:len(str)-1]
	return "[" + str + "]"
}

func (this *Codec) serializeLoop(node *TreeNode, index int) {
	if node == nil {
		return
	}
	this.serializeString[index] = strconv.Itoa(node.Val)
	this.serializeLoop(node.Left, 2*index+1)
	this.serializeLoop(node.Right, 2*index+2)
}

var root *TreeNode

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.deserializeList = []string{}
	data = data[1 : len(data)-1]
	this.deserializeList = strings.Split(data, ",")
	if this.deserializeList[0] == "null" {
		return root
	}
	root = this.deserializeLoop(0)
	return root
}

const (
	L = "Left"
	R = "Right"
)

func (this *Codec) deserializeLoop(index int) *TreeNode {
	//var node *TreeNode
	if index > len(this.deserializeList)-1 {
		return nil
	}
	if this.deserializeList[index] == "null" {
		return nil
	}
	value, err := strconv.Atoi(this.deserializeList[index])
	if err != nil {
		panic(err)
	}
	root = &TreeNode{
		Val:   value,
		Left:  this.deserializeLoop(2*index + 1),
		Right: this.deserializeLoop(2*index + 2),
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
