package heap

import "pppobear.cn/AlgorithmGolangImplement/sort"

func Sort(sorter sort.Sorter) {
	n := sorter.Len()
	for i := (n - 2) / 2; i >= 0; i-- {
		shiftDown(sorter, n, i)
	}
	for i := n - 1; i > 0; i-- {
		sorter.Swap(0, i)
		shiftDown(sorter, i, 0)
	}
}

func shiftDown(sorter sort.Sorter, size int, i int) {
	var maxChild int
	for 2*i+1 < size {
		maxChild = 2*i + 1
		if maxChild+1 < size && sorter.Less(maxChild, maxChild+1) {
			maxChild++
		}
		if sorter.Less(maxChild, i) {
			break
		}
		sorter.Swap(maxChild, i)
		i = maxChild
	}
}
