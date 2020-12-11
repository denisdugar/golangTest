package unique

func UniqueArr(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	mass := []int{}
	test := make(map[int]bool)
	for _, num := range arr {
		if _, value := test[num]; !value {
			test[num] = true
			mass = append(mass, num)
		}
	}
	return mass
}
