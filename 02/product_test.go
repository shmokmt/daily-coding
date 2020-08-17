package main

import (
	"reflect"
	"testing"
)

/*
This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/
func TestProduct(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(Product(input1), []int{120, 60, 40, 30, 24}) {
		t.Error("Product is invalid")
	}
	input2 := []int{3, 2, 1}
	if !reflect.DeepEqual(Product(input2), []int{2, 3, 6}) {
		t.Error("Product is invalid")
	}

}
