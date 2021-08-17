package matrix

import (
	"fmt"
	"log"
)

type Matrix struct {
	data           []int
	length, height int
}

func NewMatrix(length int, height int, values ...int) *Matrix {
	if len(values) > 0 && len(values) < height*length {
		log.Panicf("incorrect number of values %d", height*length)
	}
	m := Matrix{
		data:   values,
		length: length,
		height: height,
	}
	if values == nil {
		m.data = make([]int, length*height)
	}
	return &m
}

func (m *Matrix) String() {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.length; j++ {
			fmt.Printf("%d ", m.data[m.length*i+j])
		}
		println()
	}
}

func Multiple(ma, mb *Matrix) *Matrix {
	if ma.length != mb.height {
		log.Panicf("uncompability matrix dimensions first [%d,%d], second [%d,%d]",
			ma.length, ma.height, mb.length, mb.height)
	}

	mc := NewMatrix(mb.length, ma.height)
	for i := 0; i < ma.height; i++ {
		for j := 0; j < mb.length; j++ {
			for k := 0; k < ma.length; k++ {
				mc.data[i*mc.length+j] += ma.data[ma.length*i+k] * mb.data[k*mb.length+j]
			}

		}
	}
	return mc
}
