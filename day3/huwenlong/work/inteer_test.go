package work

import "testing"

func TestGraph(t *testing.T)  {
	 tests := []struct{
		graph Graph
		area float64
	}{
		{Triangle{2.5,4},5},
		{Circle{1},3.14},
	}
	 for _,test := range tests{
	 	if area := test.graph.Area(); area != test.area{
	 		t.Errorf("Error: Area Input graph: %v area: %f want %f",test.graph,area,test.area)
		}
	 }
}
