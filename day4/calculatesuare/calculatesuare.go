package calculatesuare

import (
	"errors"
	"projects/day1/day4/shapes/circle"
	"projects/day1/day4/shapes/rectangle"
	"projects/day1/day4/shapes/triangle"
	"strconv"
)

type MyShapes interface {
	Area() float64
	IsValid() bool
}

func CalculateSuare(args []string) (float64, error) {
	var sh MyShapes
	noArg  := errors.New("no arguments")
	noValid := errors.New("no valid")
	noValidSide:= errors.New("one side no valid")
	if len(args)<2{
		return 0, noArg
	}
	if len(args) == 2 {
		r, err := strconv.ParseFloat(args[1], 64)
		if err!=nil{
			return 0, noValidSide
		}
		sh = circle.NewCircle(r)
	}
	if len(args) == 3 {
		a, err := strconv.ParseFloat(args[1], 64)
		b, err1 := strconv.ParseFloat(args[2], 64)
		if err!=nil || err1!=nil{
			return 0, noValidSide
		}
		sh = rectangle.NewRectangle(a, b)
	}
	if len(args) == 4 {
		a, err := strconv.ParseFloat(args[1], 64)
		b, err1 := strconv.ParseFloat(args[2], 64)
		c, err2 := strconv.ParseFloat(args[3], 64)
		if err!=nil || err1!=nil || err2!=nil{
			return 0, noValidSide
		}
		sh = triangle.NewTriangle(a, b, c)
	}

	if sh.IsValid() {
		return sh.Area(), nil
	}
	return 0, noValid
}
