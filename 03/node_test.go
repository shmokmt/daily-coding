package main

import "testing"

/*
This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'
*/
func Test(t *testing.T) {
	want := `{"left":{"Val":"left","Left":{"Val":"left.left","Left":null,"Right":null},"Right":null},"right":{"Val":"right","Left":null,"Right":null},"val":"root"}`
	node := &Node{
		Val: "root",
		Left: &Node{
			Val: "left",
			Left: &Node{
				Val: "left.left",
			},
		},
		Right: &Node{
			Val: "right",
		},
	}
	if got := Serialize(node); got != want {
		t.Errorf("want:%s , got: %s", want, got)
	}
	if got := Deserialize(Serialize(node)); got.Left.Left.Val != "left.left" {
		t.Errorf("want: left.left, got: %s", got.Left.Left.Val)
	}

}
