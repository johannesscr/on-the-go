package sum

// Ints adds together a list of integers
func Ints(i ...int) int {
	return ints(i)
}

// ints recursively adds together a slice of integers
func ints(i []int) int {
	if len(i) == 0 {
		return 0
	}
	return ints(i[1:]) + i[0]
}
