package structure

import "testing"

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
