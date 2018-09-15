// Copyright 2014 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package mathNodes

import (
	"testing"
)

func TestPlus(t *testing.T) {
	if g, w := Math["+"].Float64Function([]float64{2, 3, 4, 5}), 5.0; g != w {
		t.Errorf("Math[+]([2,3,4,5]) = %v, want %v", g, w)
	}
}

var result float64

func runBenchmark(b *testing.B, sym string) {
	x := []float64{2, 3, 4, 5}
	var v float64
	f := Math[sym]
	for i := 0; i < b.N; i++ {
		v = f.Float64Function(x)
	}
	result = v
}

func BenchmarkPlus(b *testing.B)     { runBenchmark(b, "+") }
func BenchmarkMinus(b *testing.B)    { runBenchmark(b, "-") }
func BenchmarkMultiply(b *testing.B) { runBenchmark(b, "*") }
func BenchmarkDivide(b *testing.B)   { runBenchmark(b, "/") }
func BenchmarkMod(b *testing.B)      { runBenchmark(b, "Mod") }
func BenchmarkPow(b *testing.B)      { runBenchmark(b, "Pow") }
func BenchmarkSqrt(b *testing.B)     { runBenchmark(b, "Sqrt") }
func BenchmarkExp(b *testing.B)      { runBenchmark(b, "Exp") }
func BenchmarkPow10(b *testing.B)    { runBenchmark(b, "Pow10") }
func BenchmarkLn(b *testing.B)       { runBenchmark(b, "Ln") }
func BenchmarkLog(b *testing.B)      { runBenchmark(b, "Log") }
func BenchmarkLog2(b *testing.B)     { runBenchmark(b, "Log2") }
func BenchmarkFloor(b *testing.B)    { runBenchmark(b, "Floor") }
func BenchmarkCeil(b *testing.B)     { runBenchmark(b, "Ceil") }
func BenchmarkAbs(b *testing.B)      { runBenchmark(b, "Abs") }
func BenchmarkInv(b *testing.B)      { runBenchmark(b, "Inv") }
func BenchmarkNeg(b *testing.B)      { runBenchmark(b, "Neg") }
func BenchmarkNop(b *testing.B)      { runBenchmark(b, "Nop") }
func BenchmarkX2(b *testing.B)       { runBenchmark(b, "X2") }
func BenchmarkX3(b *testing.B)       { runBenchmark(b, "X3") }
func BenchmarkX4(b *testing.B)       { runBenchmark(b, "X4") }
func BenchmarkX5(b *testing.B)       { runBenchmark(b, "X5") }
func Benchmark3Rt(b *testing.B)      { runBenchmark(b, "3Rt") }
func Benchmark4Rt(b *testing.B)      { runBenchmark(b, "4Rt") }
func Benchmark5Rt(b *testing.B)      { runBenchmark(b, "5Rt") }
func BenchmarkAdd3(b *testing.B)     { runBenchmark(b, "Add3") }
func BenchmarkSub3(b *testing.B)     { runBenchmark(b, "Sub3") }
func BenchmarkMul3(b *testing.B)     { runBenchmark(b, "Mul3") }
func BenchmarkDiv3(b *testing.B)     { runBenchmark(b, "Div3") }
func BenchmarkAdd4(b *testing.B)     { runBenchmark(b, "Add4") }
func BenchmarkSub4(b *testing.B)     { runBenchmark(b, "Sub4") }
func BenchmarkMul4(b *testing.B)     { runBenchmark(b, "Mul4") }
func BenchmarkDiv4(b *testing.B)     { runBenchmark(b, "Div4") }
func BenchmarkMin2(b *testing.B)     { runBenchmark(b, "Min2") }
func BenchmarkMin3(b *testing.B)     { runBenchmark(b, "Min3") }
func BenchmarkMin4(b *testing.B)     { runBenchmark(b, "Min4") }
func BenchmarkMax2(b *testing.B)     { runBenchmark(b, "Max2") }
func BenchmarkMax3(b *testing.B)     { runBenchmark(b, "Max3") }
func BenchmarkMax4(b *testing.B)     { runBenchmark(b, "Max4") }
func BenchmarkAvg2(b *testing.B)     { runBenchmark(b, "Avg2") }
func BenchmarkAvg3(b *testing.B)     { runBenchmark(b, "Avg3") }
func BenchmarkAvg4(b *testing.B)     { runBenchmark(b, "Avg4") }
func BenchmarkLogi(b *testing.B)     { runBenchmark(b, "Logi") }
func BenchmarkLogi2(b *testing.B)    { runBenchmark(b, "Logi2") }
func BenchmarkLogi3(b *testing.B)    { runBenchmark(b, "Logi3") }
func BenchmarkLogi4(b *testing.B)    { runBenchmark(b, "Logi4") }
func BenchmarkGau(b *testing.B)      { runBenchmark(b, "Gau") }
func BenchmarkGau2(b *testing.B)     { runBenchmark(b, "Gau2") }
func BenchmarkGau3(b *testing.B)     { runBenchmark(b, "Gau3") }
func BenchmarkGau4(b *testing.B)     { runBenchmark(b, "Gau4") }
func BenchmarkZero(b *testing.B)     { runBenchmark(b, "Zero") }
func BenchmarkOne(b *testing.B)      { runBenchmark(b, "One") }
func BenchmarkZero2(b *testing.B)    { runBenchmark(b, "Zero2") }
func BenchmarkOne2(b *testing.B)     { runBenchmark(b, "One2") }
func BenchmarkPi(b *testing.B)       { runBenchmark(b, "Pi") }
func BenchmarkE(b *testing.B)        { runBenchmark(b, "E") }
func BenchmarkSin(b *testing.B)      { runBenchmark(b, "Sin") }
func BenchmarkCos(b *testing.B)      { runBenchmark(b, "Cos") }
func BenchmarkTan(b *testing.B)      { runBenchmark(b, "Tan") }
func BenchmarkCsc(b *testing.B)      { runBenchmark(b, "Csc") }
func BenchmarkSec(b *testing.B)      { runBenchmark(b, "Sec") }
func BenchmarkCot(b *testing.B)      { runBenchmark(b, "Cot") }
func BenchmarkAsin(b *testing.B)     { runBenchmark(b, "Asin") }
func BenchmarkAcos(b *testing.B)     { runBenchmark(b, "Acos") }
func BenchmarkAtan(b *testing.B)     { runBenchmark(b, "Atan") }
func BenchmarkAcsc(b *testing.B)     { runBenchmark(b, "Acsc") }
func BenchmarkAsec(b *testing.B)     { runBenchmark(b, "Asec") }
func BenchmarkAcot(b *testing.B)     { runBenchmark(b, "Acot") }
func BenchmarkSinh(b *testing.B)     { runBenchmark(b, "Sinh") }
func BenchmarkCosh(b *testing.B)     { runBenchmark(b, "Cosh") }
func BenchmarkTanh(b *testing.B)     { runBenchmark(b, "Tanh") }
func BenchmarkCsch(b *testing.B)     { runBenchmark(b, "Csch") }
func BenchmarkSech(b *testing.B)     { runBenchmark(b, "Sech") }
func BenchmarkCoth(b *testing.B)     { runBenchmark(b, "Coth") }
func BenchmarkAsinh(b *testing.B)    { runBenchmark(b, "Asinh") }
func BenchmarkAcosh(b *testing.B)    { runBenchmark(b, "Acosh") }
func BenchmarkAtanh(b *testing.B)    { runBenchmark(b, "Atanh") }
func BenchmarkAcsch(b *testing.B)    { runBenchmark(b, "Acsch") }
func BenchmarkAsech(b *testing.B)    { runBenchmark(b, "Asech") }
func BenchmarkAcoth(b *testing.B)    { runBenchmark(b, "Acoth") }
func BenchmarkNOT(b *testing.B)      { runBenchmark(b, "NOT") }
func BenchmarkOR1(b *testing.B)      { runBenchmark(b, "OR1") }
func BenchmarkOR2(b *testing.B)      { runBenchmark(b, "OR2") }
func BenchmarkOR3(b *testing.B)      { runBenchmark(b, "OR3") }
func BenchmarkOR4(b *testing.B)      { runBenchmark(b, "OR4") }
func BenchmarkOR5(b *testing.B)      { runBenchmark(b, "OR5") }
func BenchmarkOR6(b *testing.B)      { runBenchmark(b, "OR6") }
func BenchmarkAND1(b *testing.B)     { runBenchmark(b, "AND1") }
func BenchmarkAND2(b *testing.B)     { runBenchmark(b, "AND2") }
func BenchmarkAND3(b *testing.B)     { runBenchmark(b, "AND3") }
func BenchmarkAND4(b *testing.B)     { runBenchmark(b, "AND4") }
func BenchmarkAND5(b *testing.B)     { runBenchmark(b, "AND5") }
func BenchmarkAND6(b *testing.B)     { runBenchmark(b, "AND6") }
func BenchmarkLT2A(b *testing.B)     { runBenchmark(b, "LT2A") }
func BenchmarkGT2A(b *testing.B)     { runBenchmark(b, "GT2A") }
func BenchmarkLOE2A(b *testing.B)    { runBenchmark(b, "LOE2A") }
func BenchmarkGOE2A(b *testing.B)    { runBenchmark(b, "GOE2A") }
func BenchmarkET2A(b *testing.B)     { runBenchmark(b, "ET2A") }
func BenchmarkNET2A(b *testing.B)    { runBenchmark(b, "NET2A") }
func BenchmarkLT2B(b *testing.B)     { runBenchmark(b, "LT2B") }
func BenchmarkGT2B(b *testing.B)     { runBenchmark(b, "GT2B") }
func BenchmarkLOE2B(b *testing.B)    { runBenchmark(b, "LOE2B") }
func BenchmarkGOE2B(b *testing.B)    { runBenchmark(b, "GOE2B") }
func BenchmarkET2B(b *testing.B)     { runBenchmark(b, "ET2B") }
func BenchmarkNET2B(b *testing.B)    { runBenchmark(b, "NET2B") }
func BenchmarkLT2C(b *testing.B)     { runBenchmark(b, "LT2C") }
func BenchmarkGT2C(b *testing.B)     { runBenchmark(b, "GT2C") }
func BenchmarkLOE2C(b *testing.B)    { runBenchmark(b, "LOE2C") }
func BenchmarkGOE2C(b *testing.B)    { runBenchmark(b, "GOE2C") }
func BenchmarkET2C(b *testing.B)     { runBenchmark(b, "ET2C") }
func BenchmarkNET2C(b *testing.B)    { runBenchmark(b, "NET2C") }
func BenchmarkLT2D(b *testing.B)     { runBenchmark(b, "LT2D") }
func BenchmarkGT2D(b *testing.B)     { runBenchmark(b, "GT2D") }
func BenchmarkLOE2D(b *testing.B)    { runBenchmark(b, "LOE2D") }
func BenchmarkGOE2D(b *testing.B)    { runBenchmark(b, "GOE2D") }
func BenchmarkET2D(b *testing.B)     { runBenchmark(b, "ET2D") }
func BenchmarkNET2D(b *testing.B)    { runBenchmark(b, "NET2D") }
func BenchmarkLT2E(b *testing.B)     { runBenchmark(b, "LT2E") }
func BenchmarkGT2E(b *testing.B)     { runBenchmark(b, "GT2E") }
func BenchmarkLOE2E(b *testing.B)    { runBenchmark(b, "LOE2E") }
func BenchmarkGOE2E(b *testing.B)    { runBenchmark(b, "GOE2E") }
func BenchmarkET2E(b *testing.B)     { runBenchmark(b, "ET2E") }
func BenchmarkNET2E(b *testing.B)    { runBenchmark(b, "NET2E") }
func BenchmarkLT2F(b *testing.B)     { runBenchmark(b, "LT2F") }
func BenchmarkGT2F(b *testing.B)     { runBenchmark(b, "GT2F") }
func BenchmarkLOE2F(b *testing.B)    { runBenchmark(b, "LOE2F") }
func BenchmarkGOE2F(b *testing.B)    { runBenchmark(b, "GOE2F") }
func BenchmarkET2F(b *testing.B)     { runBenchmark(b, "ET2F") }
func BenchmarkNET2F(b *testing.B)    { runBenchmark(b, "NET2F") }
func BenchmarkLT2G(b *testing.B)     { runBenchmark(b, "LT2G") }
func BenchmarkGT2G(b *testing.B)     { runBenchmark(b, "GT2G") }
func BenchmarkLOE2G(b *testing.B)    { runBenchmark(b, "LOE2G") }
func BenchmarkGOE2G(b *testing.B)    { runBenchmark(b, "GOE2G") }
func BenchmarkET2G(b *testing.B)     { runBenchmark(b, "ET2G") }
func BenchmarkNET2G(b *testing.B)    { runBenchmark(b, "NET2G") }
func BenchmarkLT3A(b *testing.B)     { runBenchmark(b, "LT3A") }
func BenchmarkGT3A(b *testing.B)     { runBenchmark(b, "GT3A") }
func BenchmarkLOE3A(b *testing.B)    { runBenchmark(b, "LOE3A") }
func BenchmarkGOE3A(b *testing.B)    { runBenchmark(b, "GOE3A") }
func BenchmarkET3A(b *testing.B)     { runBenchmark(b, "ET3A") }
func BenchmarkNET3A(b *testing.B)    { runBenchmark(b, "NET3A") }
func BenchmarkLT3B(b *testing.B)     { runBenchmark(b, "LT3B") }
func BenchmarkGT3B(b *testing.B)     { runBenchmark(b, "GT3B") }
func BenchmarkLOE3B(b *testing.B)    { runBenchmark(b, "LOE3B") }
func BenchmarkGOE3B(b *testing.B)    { runBenchmark(b, "GOE3B") }
func BenchmarkET3B(b *testing.B)     { runBenchmark(b, "ET3B") }
func BenchmarkNET3B(b *testing.B)    { runBenchmark(b, "NET3B") }
func BenchmarkLT3C(b *testing.B)     { runBenchmark(b, "LT3C") }
func BenchmarkGT3C(b *testing.B)     { runBenchmark(b, "GT3C") }
func BenchmarkLOE3C(b *testing.B)    { runBenchmark(b, "LOE3C") }
func BenchmarkGOE3C(b *testing.B)    { runBenchmark(b, "GOE3C") }
func BenchmarkET3C(b *testing.B)     { runBenchmark(b, "ET3C") }
func BenchmarkNET3C(b *testing.B)    { runBenchmark(b, "NET3C") }
func BenchmarkLT3D(b *testing.B)     { runBenchmark(b, "LT3D") }
func BenchmarkGT3D(b *testing.B)     { runBenchmark(b, "GT3D") }
func BenchmarkLOE3D(b *testing.B)    { runBenchmark(b, "LOE3D") }
func BenchmarkGOE3D(b *testing.B)    { runBenchmark(b, "GOE3D") }
func BenchmarkET3D(b *testing.B)     { runBenchmark(b, "ET3D") }
func BenchmarkNET3D(b *testing.B)    { runBenchmark(b, "NET3D") }
func BenchmarkLT3E(b *testing.B)     { runBenchmark(b, "LT3E") }
func BenchmarkGT3E(b *testing.B)     { runBenchmark(b, "GT3E") }
func BenchmarkLOE3E(b *testing.B)    { runBenchmark(b, "LOE3E") }
func BenchmarkGOE3E(b *testing.B)    { runBenchmark(b, "GOE3E") }
func BenchmarkET3E(b *testing.B)     { runBenchmark(b, "ET3E") }
func BenchmarkNET3E(b *testing.B)    { runBenchmark(b, "NET3E") }
func BenchmarkLT3F(b *testing.B)     { runBenchmark(b, "LT3F") }
func BenchmarkGT3F(b *testing.B)     { runBenchmark(b, "GT3F") }
func BenchmarkLOE3F(b *testing.B)    { runBenchmark(b, "LOE3F") }
func BenchmarkGOE3F(b *testing.B)    { runBenchmark(b, "GOE3F") }
func BenchmarkET3F(b *testing.B)     { runBenchmark(b, "ET3F") }
func BenchmarkNET3F(b *testing.B)    { runBenchmark(b, "NET3F") }
func BenchmarkLT3G(b *testing.B)     { runBenchmark(b, "LT3G") }
func BenchmarkGT3G(b *testing.B)     { runBenchmark(b, "GT3G") }
func BenchmarkLOE3G(b *testing.B)    { runBenchmark(b, "LOE3G") }
func BenchmarkGOE3G(b *testing.B)    { runBenchmark(b, "GOE3G") }
func BenchmarkET3G(b *testing.B)     { runBenchmark(b, "ET3G") }
func BenchmarkNET3G(b *testing.B)    { runBenchmark(b, "NET3G") }
func BenchmarkLT3H(b *testing.B)     { runBenchmark(b, "LT3H") }
func BenchmarkGT3H(b *testing.B)     { runBenchmark(b, "GT3H") }
func BenchmarkLOE3H(b *testing.B)    { runBenchmark(b, "LOE3H") }
func BenchmarkGOE3H(b *testing.B)    { runBenchmark(b, "GOE3H") }
func BenchmarkET3H(b *testing.B)     { runBenchmark(b, "ET3H") }
func BenchmarkNET3H(b *testing.B)    { runBenchmark(b, "NET3H") }
func BenchmarkLT3I(b *testing.B)     { runBenchmark(b, "LT3I") }
func BenchmarkGT3I(b *testing.B)     { runBenchmark(b, "GT3I") }
func BenchmarkLOE3I(b *testing.B)    { runBenchmark(b, "LOE3I") }
func BenchmarkGOE3I(b *testing.B)    { runBenchmark(b, "GOE3I") }
func BenchmarkET3I(b *testing.B)     { runBenchmark(b, "ET3I") }
func BenchmarkNET3I(b *testing.B)    { runBenchmark(b, "NET3I") }
func BenchmarkLT3J(b *testing.B)     { runBenchmark(b, "LT3J") }
func BenchmarkGT3J(b *testing.B)     { runBenchmark(b, "GT3J") }
func BenchmarkLOE3J(b *testing.B)    { runBenchmark(b, "LOE3J") }
func BenchmarkGOE3J(b *testing.B)    { runBenchmark(b, "GOE3J") }
func BenchmarkET3J(b *testing.B)     { runBenchmark(b, "ET3J") }
func BenchmarkNET3J(b *testing.B)    { runBenchmark(b, "NET3J") }
func BenchmarkLT3K(b *testing.B)     { runBenchmark(b, "LT3K") }
func BenchmarkGT3K(b *testing.B)     { runBenchmark(b, "GT3K") }
func BenchmarkLOE3K(b *testing.B)    { runBenchmark(b, "LOE3K") }
func BenchmarkGOE3K(b *testing.B)    { runBenchmark(b, "GOE3K") }
func BenchmarkET3K(b *testing.B)     { runBenchmark(b, "ET3K") }
func BenchmarkNET3K(b *testing.B)    { runBenchmark(b, "NET3K") }
func BenchmarkLT3L(b *testing.B)     { runBenchmark(b, "LT3L") }
func BenchmarkGT3L(b *testing.B)     { runBenchmark(b, "GT3L") }
func BenchmarkLOE3L(b *testing.B)    { runBenchmark(b, "LOE3L") }
func BenchmarkGOE3L(b *testing.B)    { runBenchmark(b, "GOE3L") }
func BenchmarkET3L(b *testing.B)     { runBenchmark(b, "ET3L") }
func BenchmarkNET3L(b *testing.B)    { runBenchmark(b, "NET3L") }
func BenchmarkLT4A(b *testing.B)     { runBenchmark(b, "LT4A") }
func BenchmarkGT4A(b *testing.B)     { runBenchmark(b, "GT4A") }
func BenchmarkLOE4A(b *testing.B)    { runBenchmark(b, "LOE4A") }
func BenchmarkGOE4A(b *testing.B)    { runBenchmark(b, "GOE4A") }
func BenchmarkET4A(b *testing.B)     { runBenchmark(b, "ET4A") }
func BenchmarkNET4A(b *testing.B)    { runBenchmark(b, "NET4A") }
func BenchmarkLT4B(b *testing.B)     { runBenchmark(b, "LT4B") }
func BenchmarkGT4B(b *testing.B)     { runBenchmark(b, "GT4B") }
func BenchmarkLOE4B(b *testing.B)    { runBenchmark(b, "LOE4B") }
func BenchmarkGOE4B(b *testing.B)    { runBenchmark(b, "GOE4B") }
func BenchmarkET4B(b *testing.B)     { runBenchmark(b, "ET4B") }
func BenchmarkNET4B(b *testing.B)    { runBenchmark(b, "NET4B") }
func BenchmarkLT4C(b *testing.B)     { runBenchmark(b, "LT4C") }
func BenchmarkGT4C(b *testing.B)     { runBenchmark(b, "GT4C") }
func BenchmarkLOE4C(b *testing.B)    { runBenchmark(b, "LOE4C") }
func BenchmarkGOE4C(b *testing.B)    { runBenchmark(b, "GOE4C") }
func BenchmarkET4C(b *testing.B)     { runBenchmark(b, "ET4C") }
func BenchmarkNET4C(b *testing.B)    { runBenchmark(b, "NET4C") }
func BenchmarkLT4D(b *testing.B)     { runBenchmark(b, "LT4D") }
func BenchmarkGT4D(b *testing.B)     { runBenchmark(b, "GT4D") }
func BenchmarkLOE4D(b *testing.B)    { runBenchmark(b, "LOE4D") }
func BenchmarkGOE4D(b *testing.B)    { runBenchmark(b, "GOE4D") }
func BenchmarkET4D(b *testing.B)     { runBenchmark(b, "ET4D") }
func BenchmarkNET4D(b *testing.B)    { runBenchmark(b, "NET4D") }
func BenchmarkLT4E(b *testing.B)     { runBenchmark(b, "LT4E") }
func BenchmarkGT4E(b *testing.B)     { runBenchmark(b, "GT4E") }
func BenchmarkLOE4E(b *testing.B)    { runBenchmark(b, "LOE4E") }
func BenchmarkGOE4E(b *testing.B)    { runBenchmark(b, "GOE4E") }
func BenchmarkET4E(b *testing.B)     { runBenchmark(b, "ET4E") }
func BenchmarkNET4E(b *testing.B)    { runBenchmark(b, "NET4E") }
func BenchmarkLT4F(b *testing.B)     { runBenchmark(b, "LT4F") }
func BenchmarkGT4F(b *testing.B)     { runBenchmark(b, "GT4F") }
func BenchmarkLOE4F(b *testing.B)    { runBenchmark(b, "LOE4F") }
func BenchmarkGOE4F(b *testing.B)    { runBenchmark(b, "GOE4F") }
func BenchmarkET4F(b *testing.B)     { runBenchmark(b, "ET4F") }
func BenchmarkNET4F(b *testing.B)    { runBenchmark(b, "NET4F") }
func BenchmarkLT4G(b *testing.B)     { runBenchmark(b, "LT4G") }
func BenchmarkGT4G(b *testing.B)     { runBenchmark(b, "GT4G") }
func BenchmarkLOE4G(b *testing.B)    { runBenchmark(b, "LOE4G") }
func BenchmarkGOE4G(b *testing.B)    { runBenchmark(b, "GOE4G") }
func BenchmarkET4G(b *testing.B)     { runBenchmark(b, "ET4G") }
func BenchmarkNET4G(b *testing.B)    { runBenchmark(b, "NET4G") }
func BenchmarkLT4H(b *testing.B)     { runBenchmark(b, "LT4H") }
func BenchmarkGT4H(b *testing.B)     { runBenchmark(b, "GT4H") }
func BenchmarkLOE4H(b *testing.B)    { runBenchmark(b, "LOE4H") }
func BenchmarkGOE4H(b *testing.B)    { runBenchmark(b, "GOE4H") }
func BenchmarkET4H(b *testing.B)     { runBenchmark(b, "ET4H") }
func BenchmarkNET4H(b *testing.B)    { runBenchmark(b, "NET4H") }
func BenchmarkLT4I(b *testing.B)     { runBenchmark(b, "LT4I") }
func BenchmarkGT4I(b *testing.B)     { runBenchmark(b, "GT4I") }
func BenchmarkLOE4I(b *testing.B)    { runBenchmark(b, "LOE4I") }
func BenchmarkGOE4I(b *testing.B)    { runBenchmark(b, "GOE4I") }
func BenchmarkET4I(b *testing.B)     { runBenchmark(b, "ET4I") }
func BenchmarkNET4I(b *testing.B)    { runBenchmark(b, "NET4I") }
func BenchmarkLT4J(b *testing.B)     { runBenchmark(b, "LT4J") }
func BenchmarkGT4J(b *testing.B)     { runBenchmark(b, "GT4J") }
func BenchmarkLOE4J(b *testing.B)    { runBenchmark(b, "LOE4J") }
func BenchmarkGOE4J(b *testing.B)    { runBenchmark(b, "GOE4J") }
func BenchmarkET4J(b *testing.B)     { runBenchmark(b, "ET4J") }
func BenchmarkNET4J(b *testing.B)    { runBenchmark(b, "NET4J") }
func BenchmarkLT4K(b *testing.B)     { runBenchmark(b, "LT4K") }
func BenchmarkGT4K(b *testing.B)     { runBenchmark(b, "GT4K") }
func BenchmarkLOE4K(b *testing.B)    { runBenchmark(b, "LOE4K") }
func BenchmarkGOE4K(b *testing.B)    { runBenchmark(b, "GOE4K") }
func BenchmarkET4K(b *testing.B)     { runBenchmark(b, "ET4K") }
func BenchmarkNET4K(b *testing.B)    { runBenchmark(b, "NET4K") }
func BenchmarkLT4L(b *testing.B)     { runBenchmark(b, "LT4L") }
func BenchmarkGT4L(b *testing.B)     { runBenchmark(b, "GT4L") }
func BenchmarkLOE4L(b *testing.B)    { runBenchmark(b, "LOE4L") }
func BenchmarkGOE4L(b *testing.B)    { runBenchmark(b, "GOE4L") }
func BenchmarkET4L(b *testing.B)     { runBenchmark(b, "ET4L") }
func BenchmarkNET4L(b *testing.B)    { runBenchmark(b, "NET4L") }
