// This problem was asked by Microsoft.

// A number is considered perfect if its digits sum up to exactly 10.

// Given a positive integer n, return the n-th perfect number.

// For example, given 1, you should return 19. Given 2, you should return 28.
package num

import "strconv"

// Returns perfect number.
func Perfect(n int) int {
	if n < 1 {
		return -1
	}
	ret, err := strconv.Atoi(strconv.Itoa(n) + strconv.Itoa(10-n))
	if err != nil {
		return -1
	}
	return ret
}
