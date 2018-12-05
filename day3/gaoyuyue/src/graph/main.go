package main

const PI  = 3.14

type Graph interface {
	Area() float64
}

type Triangle struct {
	bottom float64
	height float64
}

func (t Triangle) Area() float64 {
	return t.bottom * t.height / 2
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * PI
}

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}