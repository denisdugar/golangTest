package unique

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	actual   []int
	expected []int
}{
	{actual: []int{1, 1, 2, 2, 3, 1, 1, 3, 4}, expected: []int{1, 2, 3, 4}, name: "simle task"},
	{actual: []int{1, 1, 1, 1, 1}, expected: []int{1}, name: "all equal"},
	{actual: []int{}, expected: []int{}, name: "empty slice"},
	{actual: []int{1, 2, 3}, expected: []int{1, 2, 3}, name: "all unique"},
	{actual: []int{1}, expected: []int{1}, name: "one element"},
	{name: "nill"},
}

func TestFuncUnique(t *testing.T) {
	var arr []int
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			arr = UniqueArr(tt.actual)
			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("got %q, want %q", arr, tt.expected)
			}

		})
	}
}
