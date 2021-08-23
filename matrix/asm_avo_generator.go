// +build ignore
//go:generate go run asm_avo_generator.go -out operation.s -stubs operation_stub.go

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

func main() {
	TEXT("multipleMicroCore", NOSPLIT, "func(a []float32, b []float32, c []float32)")
	a := Mem{Base: Load(Param("a").Base(), GP64())}
	b := Mem{Base: Load(Param("b").Base(), GP64())}
	c := Mem{Base: Load(Param("c").Base(), GP64())}
	const YMM_BLOCK_SIZE = 32
	const UNROLL = 7
	Comment("clear all YMM registers")
	VZEROALL()

	Comment("load data of slice b to YMM registers")
	blockB := make([]VecVirtual, UNROLL)
	for i := 0; i < UNROLL; i++ {
		blockB[i] = YMM()
	}
	for i := 0; i < UNROLL; i++ {
		VMOVUPS(b.Offset(YMM_BLOCK_SIZE*i), blockB[i])
	}

	Comment("load res values to YMM registers")
	blockC := make([]VecVirtual, UNROLL)
	for i := 0; i < UNROLL; i++ {
		blockC[i] = YMM()
	}
	for i := 0; i < UNROLL; i++ {
		VMOVUPS(c.Offset(YMM_BLOCK_SIZE*i), blockC[i])
	}

	Comment("operation with a,b,res slice blocks like res = a * b + res")
	for i := 0; i < UNROLL; i++ {
		v := YMM()
		VXORPS(v, v, v)
		VBROADCASTSS(a.Offset(YMM_BLOCK_SIZE*i), v)
		VFMADD231PS(v, blockB[i], blockC[i])
	}

	Comment("store results to slice res")
	for i := 0; i < UNROLL; i++ {
		VMOVUPS(blockC[i], c.Offset(YMM_BLOCK_SIZE*i))
	}
	RET()
	Generate()
}
