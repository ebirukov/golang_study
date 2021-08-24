// +build ignore
//go:generate go run asm_avo_generator.go -out operation.s -stubs operation_stub.go

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

var maskStepBytes = creatData("mask")
var maskStep64Bytes = creatData64("mask64")

const YMM_BLOCK_SIZE = 32

func main() {
	baseMicroCoreTemplate()
	microCore16x8Template()
	microCore24x8Template()
	microCore4x5by4x4Template()
	Generate()
}

func creatData(name string) Mem {
	maskStepBytes := GLOBL(name, RODATA|NOPTR)
	DATA(0, U64(0x0000000100000001))
	DATA(8, U64(0x0000000100000001))
	DATA(16, U64(0x0000000100000001))
	DATA(24, U64(0x0000000100000001))
	return maskStepBytes
}

func creatData64(name string) Mem {
	maskStepBytes := GLOBL(name, RODATA|NOPTR)
	DATA(0, U64(0x0000000000000001))
	DATA(8, U64(0x0000000000000001))
	DATA(16, U64(0x0000000000000001))
	DATA(24, U64(0x0000000000000001))
	return maskStepBytes
}


func loadSliceToVRegisters(slicePtr Mem, numOfRegisters int) []VecVirtual {
	vRegs := make([]VecVirtual, numOfRegisters)
	for i := 0; i < numOfRegisters; i++ {
		vRegs[i] = YMM()
	}
	for i := 0; i < numOfRegisters; i++ {
		VMOVUPS(slicePtr.Offset(YMM_BLOCK_SIZE*i), vRegs[i])
	}
	return vRegs
}

func loadSlice64ToVRegisters(slicePtr Mem, numOfRegisters int) []VecVirtual {
	vRegs := make([]VecVirtual, numOfRegisters)
	for i := 0; i < numOfRegisters; i++ {
		vRegs[i] = YMM()
	}
	for i := 0; i < numOfRegisters; i++ {
		VMOVUPD(slicePtr.Offset(YMM_BLOCK_SIZE*i), vRegs[i])
	}
	return vRegs
}

func microCore16x8Template() {
	Comment("multiplying 8x3 matrix by 8x8 matrix with single precision numbers")
	TEXT("MultipleMicroCore16x8", NOSPLIT, "func(a []float32, b []float32, c []float32)")
	aPtr := Mem{Base: Load(Param("a").Base(), GP64())}
	bPtr := Mem{Base: Load(Param("b").Base(), GP64())}
	cPtr := Mem{Base: Load(Param("c").Base(), GP64())}
	VZEROALL()

	Comment("load data of slice b to YMM registers")
	blockA := loadSliceToVRegisters(aPtr, 2)

	Comment("load data of slice b to YMM registers")
	blockB := loadSliceToVRegisters(bPtr, 8)

	Comment("load c values to YMM registers")
	blockC := loadSliceToVRegisters(cPtr, 2)

	step := YMM()
	VMOVUPS(maskStepBytes, step)
	indexMask := YMM()
	tempA := YMM()
	for j := 0; j < 2; j++ {
		for i := 0; i < 8; i++ {
			Commentf("broadcast 32-bit packed %d element from register of vector a[%d] to all elements of temporary vector a[%d:%d] register ", i, j,  j, i)
			VPERMPS(blockA[j], indexMask, tempA)
			Commentf("operation with vector registers c[%d] = a[%d:%d] * b[%d] + c[%d]", j, j, i, i, j)
			VFMADD231PS(tempA, blockB[i], blockC[j])
			Commentf("increment index mask for select %d packed element", i)
			VADDPS(step, indexMask, indexMask)
		}
		Comment("clear permutation index mask")
		VXORPS(indexMask, indexMask, indexMask)
	}

	Comment("store results to slice res")
	for i := 0; i < 2; i++ {
		VMOVUPS(blockC[i], cPtr.Offset(YMM_BLOCK_SIZE*i))
	}
	RET()
}

func microCore24x8Template() {
	Commentf("multiplying 8x3 matrix by 8x8 matrix with single precision numbers")
	TEXT("MultipleMicroCore24x8", NOSPLIT, "func(a []float32, b []float32, c []float32)")
	aPtr := Mem{Base: Load(Param("a").Base(), GP64())}
	bPtr := Mem{Base: Load(Param("b").Base(), GP64())}
	cPtr := Mem{Base: Load(Param("c").Base(), GP64())}

	VZEROALL()

	Comment("load data of slice a to YMM registers")
	blockA := loadSliceToVRegisters(aPtr, 3)

	Comment("load data of slice b to YMM registers")
	blockB := loadSliceToVRegisters(bPtr, 8)

	Comment("load data of slice c to YMM registers")
	blockC := loadSliceToVRegisters(cPtr, 3)

	indexMask := YMM()
	tempA := YMM()
	for j := 0; j < 3; j++ {
		for i := 0; i < 8; i++ {
			Commentf("broadcast 32-bit packed %d element from register of vector a[%d] to all elements of temporary vector a[%d:%d] register ", i, j,  j, i)
			VPERMPS(blockA[j], indexMask, tempA)
			Commentf("operation with vector registers c[%d] = a[%d:%d] * b[%d] + c[%d]", j, j, i, i, j)
			VFMADD231PS(tempA, blockB[i], blockC[j])
			Commentf("increment index mask for select %d packed element", i)
			VADDPS(maskStepBytes, indexMask, indexMask)
		}
		Comment("clear permutation index mask")
		VXORPS(indexMask, indexMask, indexMask)
	}

	Comment("store results to slice c")
	for i := 0; i < 3; i++ {
		VMOVUPS(blockC[i], cPtr.Offset(YMM_BLOCK_SIZE*i))
	}
	RET()
}

func microCore4x5by4x4Template() {
	Commentf("multiplying 4x5 matrix by 4x4 matrix with double precision numbers")
	TEXT("MultipleMicroCore4x5by4x4", NOSPLIT, "func(a []float64, b []float64, c []float64)")
	aPtr := Mem{Base: Load(Param("a").Base(), GP64())}
	bPtr := Mem{Base: Load(Param("b").Base(), GP64())}
	cPtr := Mem{Base: Load(Param("c").Base(), GP64())}

	VZEROALL()

	Comment("load data of slice a to YMM registers")
	blockA := loadSlice64ToVRegisters(aPtr, 5)

	Comment("load data of slice b to YMM registers")
	blockB := loadSlice64ToVRegisters(bPtr, 4)

	Comment("load data of slice c to YMM registers")
	blockC := loadSlice64ToVRegisters(cPtr, 5)

	indexMask := YMM()
	tempA := YMM()
	for j := 0; j < 5; j++ {
		for i := 0; i < 4; i++ {
			Commentf("broadcast 32-bit packed %d element from register of vector a[%d] to all elements of temporary vector a[%d:%d] register ", i, j,  j, i)
			VPERMPS(blockA[j], indexMask, tempA)
			Commentf("operation with vector registers c[%d] = a[%d:%d] * b[%d] + c[%d]", j, j, i, i, j)
			VFMADD231PD(tempA, blockB[i], blockC[j])
			Commentf("increment index mask for select %d packed element", i)
			VADDPS(maskStepBytes, indexMask, indexMask)
		}
		Comment("clear permutation index mask")
		VXORPS(indexMask, indexMask, indexMask)
	}

	Comment("store results to slice c")
	for i := 0; i < 5; i++ {
		VMOVUPD(blockC[i], cPtr.Offset(YMM_BLOCK_SIZE*i))
	}
	RET()
}

func baseMicroCoreTemplate() {
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
}
