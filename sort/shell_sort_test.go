package sort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	type args struct {
		a []int
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
			args{[]int{}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
