package sum

import "testing"

func TestInts(t *testing.T) {
	tt := []struct{
		name string
		numbers []int
		sum int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"nil", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Ints(tc.numbers...)
			if s != tc.sum {
				t.Errorf("for %v expected %v got %v",
					tc.name, tc.sum, s)
			}
		})
	}
}
