package main

import (
	"fmt"
	"sort"
)

func main() {
	// []int排序
	nums := []int{2, 31, 5, 6, 3}
	// 顺序
	sort.Ints(nums)
	fmt.Printf("[]int排序：%v\n", nums)
	// 使用sort.Reverse进行逆序排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Printf("[]int逆序排序：%v\n", nums)

	// []float64排序
	floats := []float64{2.2, 6.6, -5.3, 6.66, 3.12}
	// 顺序
	sort.Float64s(floats)
	fmt.Printf("[]float64排序：%v\n", floats)
	// 使用sort.Reverse进行逆序排序
	sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
	fmt.Printf("[]float64逆序排序：%v\n", floats)

	// []string排序
	names := []string{"abc", "12", "kk", "Jordan", "Ko", "DD"}
	// 顺序
	sort.Strings(names)
	fmt.Printf("[]string排序：%v\n", names)
	// 使用sort.Reverse进行逆序排序
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Printf("[]string逆序排序：%v\n", names)
}
