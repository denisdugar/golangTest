package unique

func UniqueArr(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	mass := []int{0}
	var uniq []int
	var equal bool = false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(mass); j++ {
			if arr[i] == mass[j] {
				equal = true
			}
		}
		if equal == false {
			mass = append(mass, arr[i])
		}
		if equal == true {
			equal = false
		}
	}
	uniq = append(mass[:0], mass[1:]...)
	return uniq

}
