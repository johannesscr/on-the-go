// Package mystr does string formatting
package mystr

import "strings"

// Cat concatenates the elements of a slice
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// Join joins the string elements
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
