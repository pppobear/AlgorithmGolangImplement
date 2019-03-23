package insertion

import "pppobear.cn/AlgorithmGolangImplement/sort"

func Sort(sorter sort.Sorter) {
	for i := 1; i < sorter.Len(); i++ {
		for j := i; j > 0 && sorter.Less(j, j-1); j-- {
			sorter.Swap(j, j-1)
		}
	}
}

func SortPart(sorter sort.Sorter, l, r int) {
	for i := l + 1; i < r; i++ {
		for j := i; j > l && sorter.Less(j, j-1); j-- {
			sorter.Swap(j, j-1)
		}
	}
}
