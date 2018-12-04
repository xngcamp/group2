package main

import "testing"

func TestCrossover(t *testing.T) {
	tests := []struct {
		ns    []int
		xs    []int
		ys    []int
		newXs []int
		newYs []int
	}{
		//{[]int{1}, []int{1,2,3}, []int{10,11,12}, []int{1,11,12}, []int{10,2,3}},
		{[]int{1,3}, []int{1,1,1,1,1,1}, []int{2,2,2,2,2,2}, []int{1,2,2,1,1,1}, []int{2,1,1,2,2,2}},
	}

	for _,test := range tests {
		if newXs, newYs := Crossover(test.ns, test.xs, test.ys); !Equals(newXs, test.newXs) || !Equals(newYs, test.newYs) {
			t.Errorf(
				"Error: Crossover input ns: %v xs: %v ys: %v expect: newXs: %v newYs: %v actual: newXs: %v newYs: %v.",
				test.ns, test.xs, test.ys, test.newXs, test.newYs, newXs, newYs,
			)
		}
	}
}

func Equals(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}