package sort

import("testing"
	"reflect"
	"sort"
)

func TestSortT(t *testing.T){

	arr1 := []int{5,7,3,8,1,9,0,2,5,10}
	arrSort1 := []int{0,1,2,3,5,5,7,8,9,10}
	
	arr2 := []int{1,1,1,1,1,1,1,1}
	arrSort2 := []int{1,1,1,1,1,1,1,1}
	
	arr3 := []int{1}
	arrSort3 := []int{1}
	
	arr4 := []int{}
	arrSort4 := []int{}
	var arrSortMy []int
	var arrSortGo []int
	
	//сортировка моей функцией
	arrSortMy = SortT(arr1)

	//сортировка встроенной
	sort.Ints(arr1)
	arrSortGo = arr1
	

	//сравнение и проверка
	if !(reflect.DeepEqual(arrSortMy, arrSort1)){
		t.Errorf("got %q, want %q", arrSortMy, arrSort1)
	}
	if !(reflect.DeepEqual(arrSortGo, arrSort1)){
		t.Errorf("got %q, want %q", arrSortGo, arrSort1)
	}

	//сортировка моей функцией
	arrSortMy = SortT(arr2)

	//сортировка встроенной
	sort.Ints(arr2)
	arrSortGo = arr2
	

	//сравнение и проверка
	if !(reflect.DeepEqual(arrSortMy, arrSort2)){
		t.Errorf("got %q, want %q", arrSortMy, arrSort2)
	}
	if !(reflect.DeepEqual(arrSortGo, arrSort2)){
		t.Errorf("got %q, want %q", arrSortGo, arrSort2)
	}

	//сортировка моей функцией
	arrSortMy = SortT(arr3)

	//сортировка встроенной
	sort.Ints(arr3)
	arrSortGo = arr3
	

	//сравнение и проверка
	if !(reflect.DeepEqual(arrSortMy, arrSort3)){
		t.Errorf("got %q, want %q", arrSortMy, arrSort3)
	}
	if !(reflect.DeepEqual(arrSortGo, arrSort3)){
		t.Errorf("got %q, want %q", arrSortGo, arrSort3)
	}

	//сортировка моей функцией
	arrSortMy = SortT(arr4)

	//сортировка встроенной
	sort.Ints(arr4)
	arrSortGo = arr4
	

	//сравнение и проверка
	if !(reflect.DeepEqual(arrSortMy, arrSort4)){
		t.Errorf("got %q, want %q", arrSortMy, arrSort4)
	}
	if !(reflect.DeepEqual(arrSortGo, arrSort4)){
		t.Errorf("got %q, want %q", arrSortGo, arrSort4)
	}
}