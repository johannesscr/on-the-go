package main

import "testing"

func TestMySum(t *testing.T) {
	// test one
	x := mySum(2, 3)
	if x != 5 {
		t.Error("expected 5, got", x)
	}
	// test two
	x = mySum(3, 3)
	if x == 5 {
		t.Error("expected 6, got", x)
	}
}

func BenchmarkMySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mySum(1, 2, 3, 4, 5, 6)
	}
}