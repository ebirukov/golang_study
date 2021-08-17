package matrix

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMultiple(t *testing.T) {
	type args struct {
		ma *Matrix
		mb *Matrix
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{
			"identityMatrix1",
			args{
				ma: NewMatrix(3, 2, 1, 2, 3, 4, 5, 6),
				mb: NewMatrix(2, 3, 1, 1, 1, 1, 1, 1),
			},
			NewMatrix(2, 2, 6, 6, 15, 15),
		},
		{
			"identityMatrix2",
			args{
				ma: NewMatrix(3, 2, 1, 1, 1, 1, 1, 1),
				mb: NewMatrix(2, 3, 1, 1, 1, 1, 1, 1),
			},
			NewMatrix(2, 2, 3, 3, 3, 3),
		},
		{
			"5x3*4x5Matrix",
			args{
				ma: NewMatrix(5, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15),
				mb: NewMatrix(4, 5, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1),
			},
			NewMatrix(4, 3, 140, 125, 110, 95, 440, 400, 360, 320, 740, 675, 610, 545),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiple(tt.args.ma, tt.args.mb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMultiple1000x1000(b *testing.B) {
	benchmarkMultiple(b, 1000, 1000)
}

var res *Matrix

func createMatrixValues(n int) []int {
	rand.Seed(time.Now().Unix())
	values := make([]int, n)
	for i := range values {
		values[i] = rand.Int()
	}
	return values
}

func benchmarkMultiple(b *testing.B, length, height int) {
	ma := NewMatrix(length, height, createMatrixValues(length*height)...)
	mb := NewMatrix(length, height, createMatrixValues(height*length)...)
	b.ResetTimer()
	var mc *Matrix
	for i := 0; i < b.N; i++ {
		mc = Multiple(ma, mb)
	}
	res = mc
}
