// Package util is for utilities
package util

import "fmt"

// StrUtil is to format a special string hello
func StrUtil(s string) string {
	return fmt.Sprintf("new string %s", s)
}

// Sum adds an unlimited number of integer values together and returns the total
func Sum(x ...int)  int {
	total := 0
	for _, val := range x {
		total += val
	}
	return total
}
