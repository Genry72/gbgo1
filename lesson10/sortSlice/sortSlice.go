package sortSlice

import (
	"sort"
)

//SortSlice сортирует слайс интов по возрастанию
func SortSlice(array []int) {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}
}

func SortSort(array []int) {
	sort.Ints(array)
}
