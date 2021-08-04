package main

import (
	"golang_study/heap"
	"golang_study/sort"
)

func main() {
	values := []int{5, 1, 3, 8, 2, 4}
	sort.BuildTree(values)

	heapValues := []int{70, 40, 50, 20, 60, 100, 80, 30, 10, 90}
	h := &heap.Heap{}
	for _, key := range heapValues {
		h.Insert(key)
	}
	sorted := make([]*heap.Node, 0)
	for value := h.Remove(); value != nil; value = h.Remove() {
		sorted = append(sorted, value)
	}
	for _, node := range sorted {
		println(node.Key)
	}
}
