package sort

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want *tree
	}{
		{
			name: "case1",
			args: args{ []int{5,1,3,8,2,4} },
			want: &tree{
				value: 5,
				left: &tree{
					value: 1,
					right: &tree{
						value: 3,
						left: &tree{value: 2},
						right: &tree{value: 4},
					},
				},
				right: &tree{value: 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTree(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		value int
		t     *tree
	}
	tests := []struct {
		name string
		args args
		want *tree
	}{
		{
			name: "case0",
			args: args{ value: 5, t: nil},
			want: &tree{value: 5, left: nil, right: nil},
		},
		{
			name: "case2",
			args: args{
				value: 1,
				t: &tree{
					value: 5,
					left: nil,
					right: nil,
				},
			},
			want: &tree{
				value: 5,
				left: &tree{value: 1, left: nil, right: nil},
				right: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.value, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}