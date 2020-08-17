package main

func Product(list []int) []int {
	ret := make([]int, 0)
	for i := range list {
		ret = append(ret, product(list, i))
	}
	return ret
}

func product(list []int, k int) int {
	ret := 1
	for i, elem := range list {
		if i == k {
			continue
		}
		ret *= elem
	}
	return ret
}
