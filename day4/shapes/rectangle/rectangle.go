package rectangle

type Rectangle struct {
	A float64
	B float64
}

func (r Rectangle) Area() float64 {
	return r.A * r.B
}

func (r Rectangle) IsValid() bool {
	if r.A != 0 && r.B != 0 {
		return true
	}
	return false
}

func NewRectangle(a, b float64) Rectangle {
	var rec Rectangle
	rec.A = a
	rec.B = b
	return rec
}
