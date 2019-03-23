package sort

//Sorter 表示可排序的接口
type Sorter interface {
	Len() int
	Less(a, b int) bool
	Swap(a, b int)
}

type Algorithm func(sorter Sorter)

//IsAscending 判断Sorter是否升序
func IsAscending(arr Sorter) bool { return isSorted(arr, false) }

//IsSorted 判断Sorter是否有序
func isSorted(arr Sorter, desc bool) bool {
	n := arr.Len()
	for i := 1; i < n; i++ {
		if arr.Less(i, i-1) != desc {
			return false
		}
	}
	return true
}

//IntSlice 整型数组
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Interface []interface{}

func (Interface) Len() int           { panic("implement me") }
func (Interface) Less(a, b int) bool { panic("implement me") }
func (Interface) Swap(a, b int)      { panic("implement me") }
