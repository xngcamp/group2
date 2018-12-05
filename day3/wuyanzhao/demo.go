package main

import (
	"fmt"
	"math"
)

type Graph interface {
	Area() float64
}
type Triangle struct {
	bottom float64
	high   float64
	area   float64
}
type Circle struct {
	rid  float64
	area float64
}
type Rectangle struct {
	length float64
	width  float64
	area   float64
}

func (t Triangle) Area() float64 {
	t.area = float64(1) / float64(2) * t.bottom * t.high
	return t.area
}
func (c Circle) Area() float64 {
	c.area = math.Pi * math.Pow(c.rid, 2) / 2
	return c.area
}
func (r Rectangle) Area() float64 {
	r.area = r.length * r.width
	return r.area
}
func main() {
	var graph Graph
	//三角形
	graph = Triangle{bottom: 6, high: 4}
	fmt.Println(graph.Area())
	//圆
	graph = Circle{rid: 4}
	fmt.Println(graph.Area())
	//矩形
	graph = Rectangle{length: 4, width: 4}
	fmt.Println(graph.Area())
}
