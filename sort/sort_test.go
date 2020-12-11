package sort

import (
	"reflect"
	"sort"
	"testing"
)

var testCases = []struct {
	name     string
	actual   []int
	expected []int
}{
	{actual: []int{5, 7, 3, 8, 1, 9, 0, 2, 5, 10}, expected: []int{0, 1, 2, 3, 5, 5, 7, 8, 9, 10}, name: "simple sort slice"},
	{actual: []int{1, 1, 1, 1, 1, 1, 1, 1}, expected: []int{1, 1, 1, 1, 1, 1, 1, 1}, name: "sort equal"},
	{actual: []int{1}, expected: []int{1}, name: "sort one"},
	{actual: []int{}, expected: []int{}, name: "sort empty"},
	{name: "sort nil"},
	{actual: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}, name: "sort sorted slice"},
}

func TestSortMy(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			SortT(tt.actual)
			if !reflect.DeepEqual(tt.actual, tt.expected) {
				t.Errorf("got %q, want %q", tt.actual, tt.expected)
			}

		})
	}
}
func TestSortGo(t *testing.T) {
	for _, tt := range testCases {
		t.Run("TestSortGo", func(t *testing.T) {
			sort.Ints(tt.actual)

			if !reflect.DeepEqual(tt.actual, tt.expected) {
				t.Errorf("got %q, want %q", tt.actual, tt.expected)
			}
		})
	}
}
