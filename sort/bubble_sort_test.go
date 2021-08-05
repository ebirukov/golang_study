package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{70, 40, 50, 20, 60, 100, 80, 30, 10, 90}},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		},
		{
			"case2",
			args{[]int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10}},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		},
		{
			"case3",
			args{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}},
			[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		},
		{
			"case3",
			args{[]int{100, 100, 80, 80, 70, 70, 10, 10, 30, 30}},
			[]int{10, 10, 30, 30, 70, 70, 80, 80, 100, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
