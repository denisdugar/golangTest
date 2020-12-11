package sort

func SortT(ar []int) {
	for gap := len(ar) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(ar); i++ {
			x := ar[i]
			j := i
			for ; j >= gap && ar[j-gap] > x; j -= gap {
				ar[j] = ar[j-gap]
			}
			ar[j] = x
		}
	}
}
