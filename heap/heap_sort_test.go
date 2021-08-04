package heap

import (
	"reflect"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	type heap struct {
		heapArray []*Node
	}
	type args struct {
		key int
	}
	tests := []struct {
		name string
		heap heap
		args args
		want int
	}{
		{
			heap: heap{},
			args: args{70},
			want: 0,
		},
		{
			heap: heap{[]*Node{{70}}},
			args: args{40},
			want: 1,
		},
		{
			heap: heap{[]*Node{{70}, {40}}},
			args: args{50},
			want: 2,
		},
		{
			heap: heap{[]*Node{{70}, {40}, {50}}},
			args: args{20},
			want: 3,
		},
		{
			heap: heap{[]*Node{{70}, {40}, {50}, {20}}},
			args: args{60},
			want: 1,
		},
		{
			heap: heap{[]*Node{{70}, {60}, {50}, {20}, {40}}},
			args: args{100},
			want: 0,
		},
		{
			heap: heap{[]*Node{{100}, {60}, {70}, {20}, {40}, {50}}},
			args: args{80},
			want: 2,
		},
		{
			heap: heap{[]*Node{{100}, {60}, {80}, {20}, {40}, {50}, {70}}},
			args: args{30},
			want: 3,
		},
		{
			heap: heap{[]*Node{{100}, {60}, {80}, {30}, {40}, {50}, {70}, {20}}},
			args: args{10},
			want: 8,
		},
		{
			heap: heap{[]*Node{{100}, {60}, {80}, {30}, {40}, {50}, {70}, {20}, {10}}},
			args: args{90},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := &Heap{
				heapArray: tt.heap.heapArray,
			}
			if got := heap.Insert(tt.args.key); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Remove(t *testing.T) {
	type heap struct {
		heapArray []*Node
	}
	tests := []struct {
		name string
		heap heap
		want *Node
	}{
		{
			heap: heap{
				[]*Node{{100}, {90}, {80}, {30}, {60}, {50}, {70}, {20}, {10}, {40}},
			},
			want: &Node{100},
		},
		{
			heap: heap{
				[]*Node{{90}, {60}, {80}, {30}, {40}, {50}, {70}, {20}, {10}},
			},
			want: &Node{90},
		},
		{
			heap: heap{
				[]*Node{{80}, {60}, {70}, {30}, {40}, {50}, {10}, {20}},
			},
			want: &Node{80},
		},
		{
			heap: heap{
				[]*Node{{70}, {60}, {50}, {30}, {40}, {20}, {10}},
			},
			want: &Node{70},
		},
		{
			heap: heap{
				[]*Node{{60}, {40}, {50}, {30}, {10}, {20}},
			},
			want: &Node{60},
		},
		{
			heap: heap{
				[]*Node{{50}, {40}, {20}, {30}, {10}},
			},
			want: &Node{50},
		},
		{
			heap: heap{
				[]*Node{{40}, {30}, {20}, {10}},
			},
			want: &Node{40},
		},
		{
			heap: heap{
				[]*Node{{30}, {20}, {10}},
			},
			want: &Node{30},
		},
		{
			heap: heap{
				[]*Node{{20}, {10}},
			},
			want: &Node{20},
		},
		{
			heap: heap{
				[]*Node{{10}},
			},
			want: &Node{10},
		},
		{
			heap: heap{
				[]*Node{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := &Heap{
				heapArray: tt.heap.heapArray,
			}
			if got := heap.Remove(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
