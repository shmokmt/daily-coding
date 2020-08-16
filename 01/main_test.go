package main

import "testing"

/*
1[Easy]
This problem was recently asked by Google.

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
*/
func TestRun(t *testing.T) {
	t.Run("fail case", func(t *testing.T) {
		if HasCombination([]int{10, 15, 3, 7}, 30) {
			t.Error("HasCombination is invalid")
		}
	})
	t.Run("success case", func(t *testing.T) {
		if !HasCombination([]int{10, 15, 3, 7}, 17) {
			t.Error("HasCombination is invalid")
		}
	})
}
