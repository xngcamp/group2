package main

import (
	"fmt"
	"math"
)

type Graphic interface {
	area() float64
}
type Rect struct{
	weight,height float64
}
type Circle struct{
	radius float64
}
type Trian struct {
	di,gao float64
}
func(r Rect) area()float64{
	return r.weight * r.height
}
func (c Circle) area()float64{
	return c.radius * c.radius * math.Pi
}
func (t Trian) area()float64{
	return t.di * t.gao * 0.5
}
func measure(g Graphic){
	fmt.Println(g)
	fmt.Println(g.area())
}
func main(){
	var r Graphic =Rect{3,4}
	var c Graphic =Circle{5}
	var t Graphic =Trian{5,12}

	measure(r)
	measure(c)
	measure(t)
}


