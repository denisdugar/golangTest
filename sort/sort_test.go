package sort

import (
	"reflect"
	"sort"
	"testing"
)

var testCases = []struct {
	actual   []int
	expected []int
}{
	{actual: []int{5, 7, 3, 8, 1, 9, 0, 2, 5, 10}, expected: []int{0, 1, 2, 3, 5, 5, 7, 8, 9, 10}},
	{actual: []int{1, 1, 1, 1, 1, 1, 1, 1}, expected: []int{1, 1, 1, 1, 1, 1, 1, 1}},
	{actual: []int{1}, expected: []int{1}},
	{actual: []int{}, expected: []int{}},
	{actual: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
}

func TestFlagParser(t *testing.T) {
	var a []int
	var b []int
	for _, tt := range testCases {
		t.Run("test", func(t *testing.T) {
			a = tt.actual
			SortT(a)
			b = tt.actual
			sort.Ints(b)
			if !reflect.DeepEqual(a, tt.expected) {
				t.Errorf("got %q, want %q", a, tt.expected)
			}
			if !reflect.DeepEqual(b, tt.expected) {
				t.Errorf("got %q, want %q", b, tt.expected)
			}
		})
	}
}
