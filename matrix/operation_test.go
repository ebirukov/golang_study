package matrix

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMultiple(t *testing.T) {
	testMultiple(t, Multiple)
}

func TestMultipleV(t *testing.T) {
	testMultiple(t, MultipleV)
}

func testMultiple(t *testing.T, m func(ma, mb *Matrix) *Matrix) {
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
		{
			"4x4*4x4Matrix",
			args{
				ma: NewMatrix(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16),
				mb: NewMatrix(4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1),
			},
			NewMatrix(4, 4, 10, 10, 10, 10, 26, 26, 26, 26, 42, 42, 42, 42, 58, 58, 58, 58),
		},
		{
			"8x8*8x8Matrix",
			args{
				ma: NewMatrix(8, 8, constMatrixValues(8*8, 2.0)...),
				mb: NewMatrix(8, 8, constMatrixValues(8*8, 1.0)...),
			},
			NewMatrix(8, 8, constMatrixValues(8*8, 16.0)...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m(tt.args.ma, tt.args.mb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMultipleVector1000x1000(b *testing.B) {
	benchmarkMultiple(b, 1000, 1000, MultipleV)
}

func benchmarkMultipleAVX2_8x8(b *testing.B) {
	benchmarkMultiple(b, 8, 8, MultipleAVX)
}

func BenchmarkMultiple_8x8(b *testing.B) {
	benchmarkMultiple(b, 8, 8, Multiple)
}

func BenchmarkMultipleAVX2_1000x1000(b *testing.B) {
	benchmarkMultiple(b, 1000, 1000, MultipleAVX)
}

func BenchmarkMultiple1000x1000(b *testing.B) {
	benchmarkMultiple(b, 1000, 1000, Multiple)
}


var res []float32 = make([]float32, 24*8)

func BenchmarkMultipleMicroCore24x8(b *testing.B) {
	benchmarkMultipleMicroCore(b, 24, 8, MultipleMicroCore24x8)
}

func BenchmarkMultipleMicroCore16x8(b *testing.B) {
	benchmarkMultipleMicroCore(b, 16, 8, MultipleMicroCore16x8)
}

func benchmarkMultipleMicroCore(b *testing.B, length, height int, coreMethod func([]float32, []float32, []float32)) {
	ma := createMatrixValues(length*height)
	mb := createMatrixValues(height*length)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coreMethod(ma, mb, res)
	}
}

func constMatrixValues(n int, value float32) []float32 {
	rand.Seed(time.Now().Unix())
	values := make([]float32, n)
	for i := range values {
		values[i] = value
	}
	return values
}

func createMatrixValues(n int) []float32 {
	rand.Seed(time.Now().Unix())
	values := make([]float32, n)
	for i := range values {
		values[i] = rand.Float32()
	}
	return values
}

func benchmarkMultiple(b *testing.B, length, height int, method func(ma, mb *Matrix) *Matrix) {
	ma := NewMatrix(length, height, createMatrixValues(length*height)...)
	mb := NewMatrix(length, height, createMatrixValues(height*length)...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		method(ma, mb)
	}
}
