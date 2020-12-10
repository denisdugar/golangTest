package sort



func SortT(ar[] int) []int{
	if ar==nil {return nil}
	var mass []int
	mass = ar
	for gap:=len(mass)/2; gap>0; gap/=2 {
		for i:=gap; i<len(mass); i++ {
			x:=mass[i]
			j:=i
			for ; j>=gap && mass[j-gap]>x; j-=gap {
				mass[j]=mass[j-gap]
			}
			mass[j]=x
		}
	}
	return mass
}