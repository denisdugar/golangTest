package circle

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return c.R * c.R * 3.14
}

func (c Circle) IsValid() bool {
	return c.R > 0
}

func NewCircle(r float64) Circle {
	var c Circle
	c.R = r
	return c
}