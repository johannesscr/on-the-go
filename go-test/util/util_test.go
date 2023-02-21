package util

import (
	"fmt"
	"testing"
)

func TestStrUtil(t *testing.T) {
	// test one
	s := StrUtil("test string")
	if s != "new string test string" {
		t.Errorf("Expected '%s' got '%s'",
			"new string test string", s)
	}
}

func ExampleSum() {
	fmt.Println(Sum(2, 3))
	fmt.Println(Sum(6, 3))
	// Output:
	// 5
	// 9
}

func ExampleSum_second() {
	fmt.Println(Sum(6, 3))
	// Output:
	// 9
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkSum_second(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
			18, 19, 20)
	}
}

func TestSum(t *testing.T) {
	// test table
	type test struct {
		values []int
		answer int
	}

	tests := []test{
		test{[]int{1, 3}, 4},
		test{[]int{12, 5, 1, 2}, 20},
		test{[]int{-12, 5}, -7},
	}
	for _, v := range tests {
		x := Sum(v.values...)
		if x != v.answer {
			t.Errorf("Expected %d got %d", v.answer, x)
		}
	}
}