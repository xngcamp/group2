package main

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		graph Graph
		area  float64
	}{
		{Rectangle{1, 2}, 2},
		{Circle{1}, PI},
		{Triangle{2, 3}, 3},
	}
	for _,test := range tests {
		if area := test.graph.Area(); area != test.area {
			t.Errorf("Error: Area input: %v expect: %v actual: %v", test.graph, test.area, area)
		}
	}
}
