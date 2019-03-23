package shell

import "pppobear.cn/AlgorithmGolangImplement/sort"

func Sort(sorter sort.Sorter) {
	n, h := sorter.Len(), 0
	for h < n {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < n; i++ {
			for j := i; j >= h && sorter.Less(j, j-h); j -= h {
				sorter.Swap(j, j-h)
			}
		}
		h /= 3
	}
}
