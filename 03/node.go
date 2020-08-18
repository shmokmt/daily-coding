package main

import "encoding/json"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func Serialize(root *Node) string {
	if root == nil {
		return ""
	}
	hash := make(map[string]interface{})
	hash["val"] = root.Val
	if root.Left != nil {
		hash["left"] = root.Left
	}
	if root.Right != nil {
		hash["right"] = root.Right
	}
	b, err := json.Marshal(hash)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Deserialize(s string) *Node {
	var node Node
	err := json.Unmarshal([]byte(s), &node)
	if err != nil {
		panic(err)
	}
	return &node
}
