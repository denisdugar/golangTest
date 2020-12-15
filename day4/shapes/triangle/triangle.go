package triangle

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	return (t.A * t.B) / 2
}

func NewTriangle(a, b, c float64) Triangle {
	var t Triangle
	t.A = a
	t.B = b
	t.C = c
	return t
}

func (t Triangle) IsValid() bool {
	if t.A != 0 && t.B != 0 && t.C != 0 {
		return true
	}
	return false
}
