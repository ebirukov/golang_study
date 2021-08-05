package sort

import (
	"reflect"
	"testing"
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
			if got := Sort(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
