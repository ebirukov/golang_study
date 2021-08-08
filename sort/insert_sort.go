package sort

import "sort"

func LibInsertionSort(arr []int) []int {
	data := sort.IntSlice(arr)
	for i := 1; i < data.Len(); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
	return data
}

func InsertionSort(values []int) []int {
	for current := 1; current < len(values); current++ {
		leftIdx := current
		checkedValue := values[current]
		for k := current - 1; k >= 0; k-- {
			if values[k] <= checkedValue {
				leftIdx = k + 1
				break
			}
			leftIdx--
		}
		if leftIdx < current {
			copy(values[leftIdx+1:current+1], values[leftIdx:current])
			values[leftIdx] = checkedValue
		}
	}
	return values
}
