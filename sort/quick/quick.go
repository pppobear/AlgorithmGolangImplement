package quick

import (
	"math/rand"
	"pppobear.cn/AlgorithmGolangImplement/sort"
	"pppobear.cn/AlgorithmGolangImplement/sort/insertion"
)

const insertionSortThreshold = 47

func Sort(sorter sort.Sorter) {
	qSort(sorter, 0, sorter.Len())
}

func qSort(sorter sort.Sorter, l, r int) {
	if r-l < insertionSortThreshold {
		insertion.SortPart(sorter, l, r)
		return
	}
	sorter.Swap(l, rand.Intn(r-l)+l)
	i, j := l+1, r-1
	for {
		for i < r && sorter.Less(i, l) {
			i++
		}
		for j > l && sorter.Less(l, j) {
			j--
		}
		if j < i {
			break
		}
		sorter.Swap(i, j)
		i++
		j--
	}
	sorter.Swap(l, j)
	qSort(sorter, l, j)
	qSort(sorter, i, r)
}
