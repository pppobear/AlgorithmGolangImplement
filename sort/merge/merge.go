package merge

import (
	"pppobear.cn/AlgorithmGolangImplement/sort"
	"pppobear.cn/AlgorithmGolangImplement/sort/insertion"
)

const insertionSortThreshold = 47

func SortBU(sorter sort.Sorter) {
	var (
		n   = sorter.Len()
		ist = insertionSortThreshold
	)
	for i := 0; i < n; i += ist {
		if i+ist < n {
			insertion.SortPart(sorter, i, i+ist)
		} else {
			insertion.SortPart(sorter, i, n)
		}
	}
	for size := ist; size < n; size *= 2 {
		for i := 0; i < n-size; i += 2 * size {
			if sorter.Less(size+i, size+i-1) {
				if i+2*size < n {
					merge(sorter, i, i+size, i+2*size)
				} else {
					merge(sorter, i, i+size, n)
				}
			}
		}
	}
}

func SortRec(sorter sort.Sorter) {
	sortRecursive(sorter, 0, sorter.Len())
}

func sortRecursive(sorter sort.Sorter, l, r int) {
	if r-l < insertionSortThreshold {
		insertion.SortPart(sorter, l, r)
		return
	}
	mid := int(uint(l+r) >> 1)
	sortRecursive(sorter, l, mid)
	sortRecursive(sorter, mid, r)
	merge(sorter, l, mid, r)
}

func merge(sorter sort.Sorter, l, mid, r int) {
	sl := sorter.(sort.IntSlice)
	aux := make(sort.IntSlice, r-l)
	copy(aux, sl[l:r])
	i, j := l, mid
	for k := i; k < r; k++ {
		if j == r {
			sl[k] = aux[i-l]
			i++
		} else if i == mid {
			sl[k] = aux[j-l]
			j++
		} else if aux.Less(i-l, j-l) {
			sl[k] = aux[i-l]
			i++
		} else {
			sl[k] = aux[j-l]
			j++
		}
	}
}
