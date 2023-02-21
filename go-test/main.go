// Package main the entry to the package
package main

import (
	"fmt"
	"github.com/JohannesScr/test-in-go/util"
)

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println(util.StrUtil("hi"))
}

// mySum computes the total of
func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}