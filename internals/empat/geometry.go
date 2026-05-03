package empat

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius int
}
type Rectangle struct {
	Height int
	Width  int
}

func (c *Circle) Area() float64 {
	radiusQuadratic := float64(c.Radius * c.Radius)
	return float64(math.Pi * radiusQuadratic)
}
func (r *Rectangle) Area() float64 {
	return float64(r.Height * r.Width)
}
