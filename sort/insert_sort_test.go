package sort

import (
	"golang_study/heap"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"casePair",
			args{[]int{2, 1}},
			[]int{1, 2},
		},
		{
			"caseSortedValues",
			args{[]int{1, 2}},
			[]int{1, 2},
		},
		{
			"caseRandomOrder",
			args{[]int{70, 40, 50, 20, 60, 100, 80, 30, 10, 90}},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		},
		{
			"caseSameValues",
			args{[]int{70, 70, 70, 70, 70, 70, 70, 70, 70, 70}},
			[]int{70, 70, 70, 70, 70, 70, 70, 70, 70, 70},
		},
		{
			"caseEmpty",
			args{[]int{}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLibInsertionSort(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"casePair",
			args{[]int{2, 1}},
			[]int{1, 2},
		},
		{
			"caseSortedValues",
			args{[]int{1, 2}},
			[]int{1, 2},
		},
		{
			"caseRandomOrder",
			args{[]int{70, 40, 50, 20, 60, 100, 80, 30, 10, 90}},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		},
		{
			"caseSameValues",
			args{[]int{70, 70, 70, 70, 70, 70, 70, 70, 70, 70}},
			[]int{70, 70, 70, 70, 70, 70, 70, 70, 70, 70},
		},
		{
			"caseEmpty",
			args{[]int{}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LibInsertionSort(tt.args.values)
			if !reflect.DeepEqual(tt.args.values, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", tt.args.values, tt.want)
			}
		})
	}
}

func generateRandomIntArray(size int) []int {
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}

var arr []int
var h = &heap.Heap{}

func init() {
	arr = generateRandomIntArray(1000000)
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h.Sort(arr)
	}
}

func BenchmarkInt(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := arr
		b.StartTimer()
		sort.Stable(sort.IntSlice(data))
		b.StopTimer()
	}
}

func BenchmarkStableInt64K(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := generateRandomIntArray(1000000)
		b.StartTimer()
		sort.Stable(sort.IntSlice(data))
		b.StopTimer()
	}
}

func BenchmarkSort(b *testing.B) {
	a := make([]int, len(arr))
	copy(a, arr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(a)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	a := make([]int, len(arr))
	copy(a, arr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(a)
	}
}

func BenchmarkShellSort(b *testing.B) {
	a := make([]int, len(arr))
	copy(a, arr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShellSort(a)
	}
}

func BenchmarkLibInsertionSort(b *testing.B) {
	a := make([]int, len(arr))
	copy(a, arr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LibInsertionSort(a)

	}
}
