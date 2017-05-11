package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		for data.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArr []int

func (p IntArr) Len() int {
	return len(p)
}

func (p IntArr) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

/*func intSort(a []){
	Sort(IntArr(a))
}*/
