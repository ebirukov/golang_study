// Code generated by command: go run asm_avo_generator.go -out operation.s -stubs operation_stub.go. DO NOT EDIT.

#include "textflag.h"

// func multipleMicroCore(a []float32, b []float32, c []float32)
// Requires: AVX, FMA3
TEXT ·multipleMicroCore(SB), NOSPLIT, $0-72
	MOVQ a_base+0(FP), AX
	MOVQ b_base+24(FP), CX
	MOVQ c_base+48(FP), DX

	// clear all YMM registers
	VZEROALL

	// load data of slice b to YMM registers
	VMOVUPS (CX), Y0
	VMOVUPS 32(CX), Y1
	VMOVUPS 64(CX), Y2
	VMOVUPS 96(CX), Y3
	VMOVUPS 128(CX), Y4
	VMOVUPS 160(CX), Y5
	VMOVUPS 192(CX), Y6

	// load res values to YMM registers
	VMOVUPS (DX), Y7
	VMOVUPS 32(DX), Y8
	VMOVUPS 64(DX), Y9
	VMOVUPS 96(DX), Y10
	VMOVUPS 128(DX), Y11
	VMOVUPS 160(DX), Y12
	VMOVUPS 192(DX), Y13

	// operation with a,b,res slice blocks like res = a * b + res
	VXORPS       Y14, Y14, Y14
	VBROADCASTSS (AX), Y14
	VFMADD231PS  Y14, Y0, Y7
	VXORPS       Y0, Y0, Y0
	VBROADCASTSS 32(AX), Y0
	VFMADD231PS  Y0, Y1, Y8
	VXORPS       Y0, Y0, Y0
	VBROADCASTSS 64(AX), Y0
	VFMADD231PS  Y0, Y2, Y9
	VXORPS       Y0, Y0, Y0
	VBROADCASTSS 96(AX), Y0
	VFMADD231PS  Y0, Y3, Y10
	VXORPS       Y0, Y0, Y0
	VBROADCASTSS 128(AX), Y0
	VFMADD231PS  Y0, Y4, Y11
	VXORPS       Y0, Y0, Y0
	VBROADCASTSS 160(AX), Y0
	VFMADD231PS  Y0, Y5, Y12
	VXORPS       Y0, Y0, Y0
	VBROADCASTSS 192(AX), Y0
	VFMADD231PS  Y0, Y6, Y13

	// store results to slice res
	VMOVUPS Y7, (DX)
	VMOVUPS Y8, 32(DX)
	VMOVUPS Y9, 64(DX)
	VMOVUPS Y10, 96(DX)
	VMOVUPS Y11, 128(DX)
	VMOVUPS Y12, 160(DX)
	VMOVUPS Y13, 192(DX)
	RET