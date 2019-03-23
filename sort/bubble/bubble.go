package bubble

import "pppobear.cn/AlgorithmGolangImplement/sort"

func Sort(sorter sort.Sorter) {
	n := sorter.Len()
	var newN int
	for n > 0 {
		newN = 0
		for i := 1; i < n; i++ {
			if sorter.Less(i, i-1) {
				sorter.Swap(i, i-1)
				newN = i
			}
		}
		n = newN
	}
}
