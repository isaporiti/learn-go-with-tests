package perimeter

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
