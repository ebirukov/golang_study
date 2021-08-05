package main

import (
	"golang_study/heap"
	"golang_study/sort"
)

func main() {
	sort.Sort([]int{2, 3, 4, 1})
	values := []int{5, 1, 3, 8, 2, 4}
	sort.BuildTree(values)

	heapValues := []int{70, 40, 50, 20, 60, 100, 80, 30, 10, 90}
	h := &heap.Heap{}

	for _, node := range h.Sort(heapValues) {
		println(node.Key)
	}
}
