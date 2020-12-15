package calculatesuare

import (
	"testing"
)

var testCases = []struct {
	name     string
	actual   []string
	expected float64
	er		 string
}{
	{actual: []string{"wqe","1"}, expected: 3.14,name:"circle"},
	{actual: []string{"wqe","1", "2"}, expected: 2,name:"rectangle"},
	{actual: []string{"wqe","3", "4", "5"}, expected: 6,name:"triangle"},
	{actual: []string{"wqe"}, expected: 0, er: "no arguments",name:"no arguments"},
	{actual: []string{"wqe","0"}, expected: 0, er:"no valid",name:"no valid"},
	{actual: []string{"wqe","0", "1"}, expected: 0, er:"no valid",name:"no valid"},
	{actual: []string{"wqe","0", "1", "2"}, expected: 0, er:"no valid",name:"no valid"},
	{actual: []string{"wqe","0", "0", "0"}, expected: 0, er:"no valid",name:"no valid"},
	{actual: []string{"wqe",""}, expected: 0, er:"one side no valid",name:"no valid"},
}

func TestCalculateSuare(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			a, err := CalculateSuare(tt.actual)
			if err!=nil{
				if a!=tt.expected || err.Error()!=tt.er{
					t.Errorf("got %f  (err: %s) , want %f (err: %s)", a,  err.Error(), tt.expected, tt.er)
				}
			}else{
				if a!=tt.expected{
					t.Errorf("got %f  (err: %s) , want %f (err: %s)", a,  err.Error(), tt.expected, tt.er)
				}
			}
		})
	}
}
