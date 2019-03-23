package sort

import (
	"math/rand"
	"time"
)

func GenerateRandomSlice(n, rangeL, rangeR uint) Sorter {
	ret := IntSlice(make([]int, n))
	if rangeL > rangeR {
		return ret
	}
	for i := range ret {
		ret[i] = rand.Intn(int(rangeR-rangeL)) + int(rangeL)
	}
	return ret
}

func GenerateNearlyOrderedSlice(n, swapTimes uint) Sorter {
	ret := IntSlice(make([]int, n))
	for i := range ret {
		ret[i] = i
	}
	rand.Shuffle(int(swapTimes), func(i, j int) { ret[i], ret[j] = ret[j], ret[i] })
	return ret
}

func TestSort(sort Algorithm, sorter Sorter) time.Duration {
	startTime := time.Now()
	sort(sorter)
	endTime := time.Now()
	return endTime.Sub(startTime)
}
