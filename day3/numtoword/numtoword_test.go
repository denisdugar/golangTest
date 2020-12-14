package numtoword

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	actual   string
	expected string
}{
	{actual: "0", expected: "ноль", name: "0"},
	{actual: "5", expected: "пять", name: "5"},
	{actual: "10", expected: "десять", name: "10"},
	{actual: "12", expected: "двенадцать", name: "12"},
	{actual: "20", expected: "двадцать", name: "20"},
	{actual: "34", expected: "тридцать четыре", name: "34"},
	{actual: "100", expected: "сто", name: "100"},
	{actual: "102", expected: "сто два", name: "102"},
	{actual: "110", expected: "сто десять", name: "110"},
	{actual: "112", expected: "сто двенадцать", name: "112"},
	{actual: "120", expected: "сто двадцать", name: "120"},
	{actual: "999", expected: "девятьсот девяносто девять", name: "999"},
	{actual: "1000", expected: "", name: "1000"},
	{actual: "-11", expected: "", name: "-11"},
	{actual: "-111", expected: "", name: "-111"},
	{actual: "-1", expected: "", name: "-1"},
	{actual: "qqq", expected: "", name: "qqq"},
	{name: "nill!"},
}

func TestFuncNumToWord(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NumToWord(tt.actual)
			if !reflect.DeepEqual(s, tt.expected) {
				t.Errorf("got %q, want %q", s, tt.expected)
			}

		})
	}
}
