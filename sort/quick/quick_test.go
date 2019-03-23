package quick

import (
	"testing"

	"pppobear.cn/AlgorithmGolangImplement/sort"
)

func TestSort(t *testing.T) {
	type args struct {
		sorter sort.Sorter
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"对1000万个元素的无序切片进行 快速排序",
			args{sort.GenerateRandomSlice(10000000, 0, 10000000*5)},
		},
		{
			"对1000万个元素的近乎有序切片进行 快速排序",
			args{sort.GenerateNearlyOrderedSlice(10000000, 10000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := sort.TestSort(Sort, tt.args.sorter)
			if !sort.IsAscending(tt.args.sorter) {
				t.Error("排序失败")
			}
			t.Logf("耗时：%f", dur.Seconds())
		})
	}
}
