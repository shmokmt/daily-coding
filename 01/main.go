package main

func HasCombination(list []int, k int) bool {
	for i, elem := range list {
		for j := i + 1; j < len(list); j++ {
			if elem+list[j] == k {
				return true
			}
		}
	}
	return false
}
