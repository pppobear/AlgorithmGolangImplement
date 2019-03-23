package selection

import "pppobear.cn/AlgorithmGolangImplement/sort"

func Sort(sorter sort.Sorter) {
	n := sorter.Len()
	var minIdx int
	for i := 0; i < n-1; i++ {
		minIdx = i
		for j := i + 1; j < n; j++ {
			if sorter.Less(j, minIdx) {
				minIdx = j
			}
		}
		sorter.Swap(minIdx, i)
	}
}
