package structure

import (
	"reflect"
	"testing"
)

func newSet(values ...int) *IntSet {
	s := NewIntSet(10000)
	for i := 0; i < len(values); i++ {
		s = s.Add(values[i])
	}
	return s
}

func TestIntSet_Add(t *testing.T) {
	type fields struct {
		words []uint64
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			"case1",
			fields{[]uint64{}},
			args{65},
			fields{[]uint64{0, 2}},
		},
		{
			"case2",
			fields{[]uint64{}},
			args{0},
			fields{[]uint64{1}},
		},
		{
			"case3",
			fields{[]uint64{0, 2}},
			args{0},
			fields{[]uint64{1, 2}},
		},
		{
			"case4",
			fields{[]uint64{}},
			args{4},
			fields{[]uint64{16}},
		},
		{
			"case5",
			fields{[]uint64{}},
			args{640},
			fields{[]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			want := &IntSet{
				words: tt.want.words,
			}
			s.Add(tt.args.v)
			if s.String() != want.String() {
				t.Errorf("Add() = %v, want %v", s.String(), want.String())
			}
		})
	}
}

func TestIntSet_String(t *testing.T) {
	type fields struct {
		words []uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"case1",
			fields{words: []uint64{1, 2}},
			"{ 0 65 }",
		},
		{
			"case2",
			fields{words: []uint64{}},
			"{ }",
		},
		{
			"case3",
			fields{[]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
			"{ 640 }",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &IntSet{
				words: tt.fields.words,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Has(t *testing.T) {
	type fields struct {
		words []uint64
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"case1",
			fields{[]uint64{}},
			args{0},
			false,
		},
		{
			"case2",
			fields{[]uint64{0, 0, 1}},
			args{128},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := IntSet{
				words: tt.fields.words,
			}
			if got := s.Has(tt.args.v); got != tt.want {
				t.Errorf("Has() = %v, want %v, bitset %s", got, tt.want, s.String())
			}
		})
	}
}

func TestIntSet_unionWith(t *testing.T) {
	type args struct {
		t *IntSet
	}

	tests := []struct {
		name    string
		initial *IntSet
		args    args
		want    *IntSet
	}{
		{
			"case1",
			newSet(456, 750),
			args{newSet(749, 750, 385, 6493)},
			newSet(749, 385, 456, 750, 6493),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.initial
			if got := s.unionWith(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unionWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_intersectWith(t *testing.T) {
	type args struct {
		t *IntSet
	}

	tests := []struct {
		name    string
		initial *IntSet
		args    args
		want    *IntSet
	}{
		{
			"case1",
			newSet(456, 750),
			args{newSet(749, 750, 385, 6493)},
			newSet(750),
		},
		{
			"case2",
			newSet(456, 457),
			args{newSet(440, 480)},
			newSet(),
		},
		{
			"case3",
			newSet(),
			args{newSet(440, 480)},
			newSet(),
		},
		{
			"case4",
			&IntSet{},
			args{&IntSet{}},
			&IntSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.initial
			if got := s.intersectWith(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("differenceWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_differenceWith(t *testing.T) {
	type args struct {
		t *IntSet
	}

	tests := []struct {
		name    string
		initial *IntSet
		args    args
		want    *IntSet
	}{
		{
			"case1",
			newSet(456, 750),
			args{newSet(749, 750, 385, 6493)},
			newSet(749, 385, 456, 6493),
		},
		{
			"case2",
			newSet(456, 457),
			args{newSet(440, 480)},
			newSet(440, 480, 456, 457),
		},
		{
			"case3",
			&IntSet{},
			args{newSet(440, 480)},
			newSet(440, 480),
		},
		{
			"case4",
			&IntSet{},
			args{&IntSet{}},
			&IntSet{},
		},
		{
			"case5",
			newSet(440, 480),
			args{newSet(440, 480)},
			newSet(),
		},
		{
			"case6",
			newSet(0, 0),
			args{newSet(0, 0)},
			newSet(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.initial
			if got := s.differenceWith(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("differenceWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Len(t *testing.T) {
	type fields struct {
		initial *IntSet
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"case1",
			fields{newSet()},
			0,
		},
		{
			"case2",
			fields{newSet(749, 750, 385, 6493)},
			4,
		},
		{
			"case3",
			fields{filledSet(500000)},
			500000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.fields.initial
			if got := s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

func filledSet(size int) *IntSet {
	values := make([]int, size)
	for i := 0; i < len(values); i++ {
		values[i] = i
	}
	return newSet(values...)
}

func BenchmarkIntSet_Len1000(b *testing.B) {
	benchmarkIntSetLen(b, 1000)
}

func BenchmarkIntSet_Len50000(b *testing.B) {
	benchmarkIntSetLen(b, 50000)
}

func BenchmarkIntSet_Len500000(b *testing.B) {
	benchmarkIntSetLen(b, 500000)
}

func benchmarkIntSetLen(b *testing.B, size int) {
	set := filledSet(size)
	var length int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		length = set.Len()
	}
	result = length
}

func BenchmarkIntSet_Clear500000(b *testing.B) {
	benchmarkIntSetClear(b, 500000)
}

func benchmarkIntSetClear(b *testing.B, size int) {
	set := filledSet(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Clear()
	}
}
