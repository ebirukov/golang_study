package main

import (
	"reflect"
	"testing"
)

func initList(n int) *Node {
	var head *Node
	for i := n; i > 0; i-- {
		head = &Node{i: i, next: head}
	}

	return head
}

func initListArr(data ...int) *Node {
	var head *Node
	for _, i := range data {
		head = &Node{i: i, next: head}
	}

	return head
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *Node
		n    int
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"emptyList",
			args{
				head: nil,
				n:    5,
			},
			nil,
		},
		{
			"tenElementList",
			args{
				head: initList(10),
				n:    5,
			},
			initListArr(10, 9, 8, 7, 5, 4, 3, 2, 1),
		},
		{
			"removeZeroElementList",
			args{
				head: initList(10),
				n:    0,
			},
			initListArr(10, 9, 8, 7, 6, 5, 4, 3, 2, 1),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_revertRecursive(t *testing.T) {
	type args struct {
		node *Node
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"oneElementList",
			args{
				node: initList(1),
			},

			initListArr(1),
		},
		{
			"tenElementList",
			args{
				node: initList(10),
			},

			initListArr(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
		},
		{
			"emptyList",
			args{
				nil,
			},
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := revertRecursive(nil, tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("revertRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_drop(t *testing.T) {
	type args struct {
		node *Node
		n    int
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"tenElementList",
			args{
				node: initList(10),
				n:    5,
			},

			initListArr(10, 9, 8, 7, 6),
		},
		{
			"tenElementList2",
			args{
				node: initList(10),
				n:    9,
			},

			initListArr(10),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := drop(tt.args.node, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("drop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dropNFromEnd(t *testing.T) {
	type args struct {
		head *Node
		n    int
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"tenElementList",
			args{
				head: initList(10),
				n:    5,
			},

			initList(5),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := dropNFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dropNFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
