package matrix

import (
	"fmt"
	"log"
)

type Matrix struct {
	data           []float32
	length, height int
}

func NewMatrix(length int, height int, values ...float32) *Matrix {
	if len(values) > 0 && len(values) < height*length {
		log.Panicf("incorrect number of values %d", height*length)
	}
	m := Matrix{
		data:   values,
		length: length,
		height: height,
	}
	if values == nil {
		m.data = make([]float32, length*height)
	}
	return &m
}

func (m *Matrix) String() {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.length; j++ {
			fmt.Printf("%f ", m.data[m.length*i+j])
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

func MultipleV(ma, mb *Matrix) *Matrix {
	if ma.length != mb.height {
		log.Panicf("uncompability matrix dimensions first [%d,%d], second [%d,%d]",
			ma.length, ma.height, mb.length, mb.height)
	}

	mc := NewMatrix(mb.length, ma.height)
	for i := 0; i < ma.height; i++ {
			for k := 0; k < ma.length; k++ {
				a := ma.data[ma.length*i+k]
				for j := 0; j < mb.length; j++ {
					mc.data[i*mc.length+j] += a * mb.data[k*mb.length+j]
				}
			}

	}
	return mc
}

func MultipleAVX(ma, mb *Matrix) *Matrix {
	if ma.length != mb.height {
		log.Panicf("uncompability matrix dimensions first [%d,%d], second [%d,%d]",
			ma.length, ma.height, mb.length, mb.height)
	}

	mc := NewMatrix(mb.length, ma.height)
	for i := 0; i < mc.length; i++ {
		mc.data[i] = 0
	}
	for i := 0; i < ma.height; i++ {
		a := ma.data[ma.length*i:ma.length*(i+1)]
		c := mc.data[i*mc.length:(i+1)*mc.length]
		//fmt.Printf("mc -%f\n", mc.data)
		for k := 0; k < ma.length; k++ {
			b := mb.data[k*mb.length:(k+1)*mb.length]
			//fmt.Printf("c -%f\n", mc.data)
			/*			fmt.Printf("a -%f\n", a)
			fmt.Printf("b -%f\n", b)*/
			multipleMicroCore(a[k:],b,c)

		}

	}
	return mc
}