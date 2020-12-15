package triangle

import (
	"math"
)

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	p := (t.A + t.B + t.C) / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func NewTriangle(a, b, c float64) Triangle {
	var t Triangle
	t.A = a
	t.B = b
	t.C = c
	return t
}

func (t Triangle) IsValid() bool {
	return t.A+t.B > t.C && t.A+t.C > t.B && t.B+t.C > t.A
}
