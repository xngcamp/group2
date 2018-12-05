package work

type Graph interface {
	Area() float64
}
type Triangle struct {
	line float64
	high float64
}
func (tri Triangle) Area() float64 {
	return (tri.line*tri.high)/2
}

type Circle struct {
	radius float64
}
func (cir Circle) Area() float64 {
	return 3.14*cir.radius*cir.radius
}

type Rectrangle struct {
	width float64
	height float64
}
func (rect Rectrangle)Area() float64 {
	return rect.width*rect.height
}
