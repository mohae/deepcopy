package deepcopy

import (
	"fmt"
	"testing"
)

var stringSliceOrig = []string{
	"element1",
	"element2",
	"element3",
}

var intSliceOrig = []int{
	1,
	2,
	3,
}

func TestStringSlice(t *testing.T) {
	cpy := InterfaceToStringSlice(nil)
	if cpy != nil {
		t.Errorf("got %v; want nil", cpy)
	}
	cpy = InterfaceToStringSlice(stringSliceOrig)
	if len(cpy) != len(stringSliceOrig) {
		t.Errorf("len of copy was %d; want %d", len(cpy), len(stringSliceOrig))
	}
	if &cpy == &stringSliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range stringSliceOrig {
		if v != cpy[i] {
			t.Errorf("got %s at index %d of the copy; want %s", cpy[i], i, v)
		}
	}
}

func TestIntSlice(t *testing.T) {
	cpy := InterfaceToIntSlice(nil)
	if cpy != nil {
		t.Errorf("got %v; want nil", cpy)
	}
	cpy = InterfaceToIntSlice(intSliceOrig)
	if len(cpy) != len(intSliceOrig) {
		t.Errorf("len of copy was %d; want %d", len(cpy), len(intSliceOrig))
	}
	if &cpy == &intSliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range intSliceOrig {
		if v != cpy[i] {
			t.Errorf("got %d at index %d of the copy; want %d", cpy[i], i, v)
		}
	}
}

// These tests test that all supported basic types are copied correctly.  This
// is done by copying a []T.
func TestRecursiveCopySlices(t *testing.T) {
	var tst interface{}
	cpy := Copy(tst)
	if cpy != nil {
		t.Errorf("got %v; want nil", cpy)
	}

	cpy = Copy(stringSliceOrig)
	cpyS, ok := cpy.([]string)
	if !ok {
		t.Errorf("copy of []string: expected the interface to contain a []string; it didn't")
		goto intSlice
	}
	if len(cpyS) != len(stringSliceOrig) {
		t.Errorf("[]string: len of copy was %d; want %d", len(cpyS), len(stringSliceOrig))
		goto intSlice
	}
	if &cpyS == &stringSliceOrig {
		t.Error("[]string: address of copy was the same as original; they should be different")
		goto intSlice
	}
	for i, v := range stringSliceOrig {
		if v != cpyS[i] {
			t.Errorf("[]string: got %s at index %d of the copy; want %s", cpyS[i], i, v)
		}
	}

intSlice:
	cpy = Copy(intSliceOrig)
	cpyI, ok := cpy.([]int)
	if !ok {
		t.Errorf("copy of []int: expected the interface to contain a []int; it didn't")
		goto int8Slice
	}
	if len(cpyI) != len(intSliceOrig) {
		t.Errorf("[]int: len of copy was %d; want %d", len(cpyI), len(intSliceOrig))
		goto int8Slice
	}
	if &cpyI == &intSliceOrig {
		t.Error("[]int: address of copy was the same as original; they should be different")
		goto int8Slice
	}
	for i, v := range intSliceOrig {
		if v != cpyI[i] {
			t.Errorf("[]int: got %d at index %d of the copy; want %d", cpyI[i], i, v)
		}
	}

int8Slice:
	int8SliceOrig := []int8{0, 1, 2, 3, 42}
	cpy = Copy(int8SliceOrig)
	cpyI8, ok := cpy.([]int8)
	if !ok {
		t.Errorf("copy of []int8: expected the interface to contain a []int8; it didn't")
		goto int16Slice
	}
	if len(cpyI8) != len(int8SliceOrig) {
		t.Errorf("[]int8: len of copy was %d; want %d", len(cpyI8), len(int8SliceOrig))
		goto int16Slice
	}
	if &cpyI8 == &int8SliceOrig {
		t.Error("[]int8: address of copy was the same as original; they should be different")
		goto int16Slice
	}
	for i, v := range int8SliceOrig {
		if v != cpyI8[i] {
			t.Errorf("[]int8: got %d at index %d of the copy; want %d", cpyI8[i], i, v)
		}
	}

int16Slice:
	int16SliceOrig := []int16{0, 1, 2, 3, 42}
	cpy = Copy(int16SliceOrig)
	cpyI16, ok := cpy.([]int16)
	if !ok {
		t.Errorf("copy of []int16: expected the interface to contain a []int16; it didn't")
		goto int32Slice
	}
	if len(cpyI16) != len(int16SliceOrig) {
		t.Errorf("[]int16: len of copy was %d; want %d", len(cpyI16), len(int16SliceOrig))
		goto int32Slice
	}
	if &cpyI16 == &int16SliceOrig {
		t.Error("[]int16: address of copy was the same as original; they should be different")
		goto int32Slice
	}
	for i, v := range int16SliceOrig {
		if v != cpyI16[i] {
			t.Errorf("[]int16: got %d at index %d of the copy; want %d", cpyI16[i], i, v)
		}
	}

int32Slice:
	int32SliceOrig := []int32{0, 1, 2, 3, 42}
	cpy = Copy(int32SliceOrig)
	cpyI32, ok := cpy.([]int32)
	if !ok {
		t.Errorf("copy of []int32: expected the interface to contain a []int32; it didn't")
		goto int64Slice
	}
	if len(cpyI32) != len(int32SliceOrig) {
		t.Errorf("[]int32: len of copy was %d; want %d", len(cpyI32), len(int32SliceOrig))
		goto int64Slice
	}
	if &cpyI32 == &int32SliceOrig {
		t.Error("[]int32: address of copy was the same as original; they should be different")
		goto int64Slice
	}
	for i, v := range int32SliceOrig {
		if v != cpyI32[i] {
			t.Errorf("[]int32: got %d at index %d of the copy; want %d", cpyI32[i], i, v)
		}
	}

int64Slice:
	int64SliceOrig := []int64{0, 1, 2, 3, 42}
	cpy = Copy(int64SliceOrig)
	cpyI64, ok := cpy.([]int64)
	if !ok {
		t.Errorf("copy of []int64: expected the interface to contain a []int64; it didn't")
		goto float32Slice
	}
	if len(cpyI64) != len(int64SliceOrig) {
		t.Errorf("[]int64: len of copy was %d; want %d", len(cpyI64), len(int64SliceOrig))
		goto float32Slice
	}
	if &cpyI64 == &int64SliceOrig {
		t.Error("[]int64: address of copy was the same as original; they should be different")
		goto float32Slice
	}
	for i, v := range int64SliceOrig {
		if v != cpyI64[i] {
			t.Errorf("[]int64: got %d at index %d of the copy; want %d", cpyI64[i], i, v)
		}
	}

float32Slice:
	float32SliceOrig := []float32{0, 1, 2, 3, 42}
	cpy = Copy(float32SliceOrig)
	cpyF32, ok := cpy.([]float32)
	if !ok {
		t.Errorf("copy of []float32: expected the interface to contain a []float32; it didn't")
		goto float64Slice
	}
	if len(cpyF32) != len(float32SliceOrig) {
		t.Errorf("[]float32: len of copy was %d; want %d", len(cpyF32), len(float32SliceOrig))
		goto float64Slice
	}
	if &cpyF32 == &float32SliceOrig {
		t.Error("[]float32: address of copy was the same as original; they should be different")
		goto float64Slice
	}
	for i, v := range float32SliceOrig {
		if v != cpyF32[i] {
			t.Errorf("[]float32: got %f at index %d of the copy; want %f", cpyF32[i], i, v)
		}
	}

float64Slice:
	float64SliceOrig := []float64{0, 1, 2, 3, 42}
	cpy = Copy(float64SliceOrig)
	cpyF64, ok := cpy.([]float64)
	if !ok {
		t.Errorf("copy of []float64: expected the interface to contain a []float64; it didn't")
		goto boolSlice
	}
	if len(cpyF64) != len(float64SliceOrig) {
		t.Errorf("[]float64: len of copy was %d; want %d", len(cpyF64), len(float64SliceOrig))
		goto boolSlice
	}
	if &cpyF64 == &float64SliceOrig {
		t.Error("[]float64: address of copy was the same as original; they should be different")
		goto boolSlice
	}
	for i, v := range float64SliceOrig {
		if v != cpyF64[i] {
			t.Errorf("[]float64: got %f at index %d of the copy; want %f", cpyF64[i], i, v)
		}
	}

boolSlice:
	boolSliceOrig := []bool{true, false, false, true}
	cpy = Copy(boolSliceOrig)
	cpyB, ok := cpy.([]bool)
	if !ok {
		t.Errorf("copy of []bool: expected the interface to contain a []bool; it didn't")
		goto infSlice
	}
	if len(cpyB) != len(boolSliceOrig) {
		t.Errorf("[]bool: len of copy was %d; want %d", len(cpyB), len(boolSliceOrig))
		goto infSlice
	}
	if &cpyB == &boolSliceOrig {
		t.Error("[]bool: address of copy was the same as original; they should be different")
		goto infSlice
	}
	for i, v := range boolSliceOrig {
		if v != cpyB[i] {
			t.Errorf("[]bool: got %t at index %d of the copy; want %t", cpyB[i], i, v)
		}
	}

infSlice:
	infSliceOrig := []interface{}{true, 1, "hello"}
	cpy = Copy(infSliceOrig)
	cpyInf, ok := cpy.([]interface{})
	if !ok {
		t.Errorf("copy of []interface{}: expected the interface to contain a []interface{}; it didn't")
		goto byteSlice
	}
	if len(cpyInf) != len(infSliceOrig) {
		t.Errorf("[]interface{}: len of copy was %d; want %d", len(cpyInf), len(infSliceOrig))
		goto byteSlice
	}
	if &cpyInf == &infSliceOrig {
		t.Error("[]interface{}: address of copy was the same as original; they should be different")
		goto byteSlice
	}
	for i, v := range infSliceOrig {
		if v != cpyInf[i] {
			t.Errorf("[]interface{}: got %v at index %d of the copy; want %v", cpyInf[i], i, v)
		}
	}

byteSlice:
	byteSlice := []byte("hello")
	cpy = Copy(byteSlice)
	cpyByte, ok := cpy.([]byte)
	if !ok {
		t.Errorf("copy of []byte: expected the interface to contain a []byte; it didn't")
		goto uintSlice
	}
	if len(cpyByte) != len(byteSlice) {
		t.Errorf("[]byte: len of copy was %d; want %d", len(cpyByte), len(byteSlice))
		goto uintSlice
	}
	if &cpyByte == &byteSlice {
		t.Error("[]byte: address of copy was the same as original; they should be different")
		goto uintSlice
	}
	if string(cpyByte) != string(byteSlice) {
		t.Errorf("[]byte: got %s; want %s", string(cpyByte), string(byteSlice))
	}

uintSlice:
	uintSliceOrig := []uint{0, 1, 2, 3, 42}
	cpy = Copy(uintSliceOrig)
	cpyUi, ok := cpy.([]uint)
	if !ok {
		t.Errorf("copy of []uint: expected the interface to contain a []uint; it didn't")
		goto uint8Slice
	}
	if len(cpyUi) != len(uintSliceOrig) {
		t.Errorf("[]uint: len of copy was %d; want %d", len(cpyUi), len(uintSliceOrig))
		goto uint8Slice
	}
	if &cpyUi == &uintSliceOrig {
		t.Error("[]uint: address of copy was the same as original; they should be different")
		goto uint8Slice
	}
	for i, v := range uintSliceOrig {
		if v != cpyUi[i] {
			t.Errorf("[]uint: got %d at index %d of the copy; want %d", cpyUi[i], i, v)
		}
	}

uint8Slice:
	uint8SliceOrig := []uint8{0, 1, 2, 3, 42}
	cpy = Copy(uint8SliceOrig)
	cpyUi8, ok := cpy.([]uint8)
	if !ok {
		t.Errorf("copy of []uint8: expected the interface to contain a []uint8; it didn't")
		goto uint16Slice
	}
	if len(cpyUi8) != len(uint8SliceOrig) {
		t.Errorf("[]uint8: len of copy was %d; want %d", len(cpyUi8), len(uint8SliceOrig))
		goto uint16Slice
	}
	if &cpyUi8 == &uint8SliceOrig {
		t.Error("[]uint8: address of copy was the same as original; they should be different")
		goto uint16Slice
	}
	for i, v := range uint8SliceOrig {
		if v != cpyUi8[i] {
			t.Errorf("[]uint8: got %d at index %d of the copy; want %d", cpyUi8[i], i, v)
		}
	}

uint16Slice:
	uint16SliceOrig := []uint16{0, 1, 2, 3, 42}
	cpy = Copy(uint16SliceOrig)
	cpyUi16, ok := cpy.([]uint16)
	if !ok {
		t.Errorf("copy of []uint16: expected the interface to contain a []uint16; it didn't")
		goto uint32Slice
	}
	if len(cpyUi16) != len(uint16SliceOrig) {
		t.Errorf("[]uint16: len of copy was %d; want %d", len(cpyUi16), len(uint16SliceOrig))
		goto uint32Slice
	}
	if &cpyUi16 == &uint16SliceOrig {
		t.Error("[]uint16: address of copy was the same as original; they should be different")
		goto uint32Slice
	}
	for i, v := range uint16SliceOrig {
		if v != cpyUi16[i] {
			t.Errorf("[]uint16: got %d at index %d of the copy; want %d", cpyUi16[i], i, v)
		}
	}

uint32Slice:
	uint32SliceOrig := []uint32{0, 1, 2, 3, 42}
	cpy = Copy(uint32SliceOrig)
	cpyUi32, ok := cpy.([]uint32)
	if !ok {
		t.Errorf("copy of []uint32: expected the interface to contain a []uint32; it didn't")
		goto uint64Slice
	}
	if len(cpyUi32) != len(uint32SliceOrig) {
		t.Errorf("[]uint32: len of copy was %d; want %d", len(cpyUi32), len(uint32SliceOrig))
		goto uint64Slice
	}
	if &cpyUi32 == &uint32SliceOrig {
		t.Error("[]uint32: address of copy was the same as original; they should be different")
		goto uint64Slice
	}
	for i, v := range uint32SliceOrig {
		if v != cpyUi32[i] {
			t.Errorf("[]uint32: got %d at index %d of the copy; want %d", cpyUi32[i], i, v)
		}
	}

uint64Slice:
	uint64SliceOrig := []uint64{0, 1, 2, 3, 42}
	cpy = Copy(uint64SliceOrig)
	cpyUi64, ok := cpy.([]uint64)
	if !ok {
		t.Errorf("copy of []uint64: expected the interface to contain a []uint64; it didn't")
		goto complex64Slice
	}
	if len(cpyUi64) != len(uint64SliceOrig) {
		t.Errorf("[]uint64: len of copy was %d; want %d", len(cpyUi64), len(uint64SliceOrig))
		goto complex64Slice
	}
	if &cpyUi64 == &uint64SliceOrig {
		t.Error("[]uint64: address of copy was the same as original; they should be different")
		goto complex64Slice
	}
	for i, v := range uint64SliceOrig {
		if v != cpyUi64[i] {
			t.Errorf("[]uint64: got %d at index %d of the copy; want %d", cpyUi64[i], i, v)
		}
	}

complex64Slice:
	complex64SliceOrig := []complex64{complex64(-65 + 11i), complex64(66 + 10i)}
	cpy = Copy(complex64SliceOrig)
	cpyC64, ok := cpy.([]complex64)
	if !ok {
		t.Errorf("copy of []complex64: expected the interface to contain a []complex64; it didn't")
		goto complex128Slice
	}
	if len(cpyC64) != len(complex64SliceOrig) {
		t.Errorf("[]complex64: len of copy was %d; want %d", len(cpyC64), len(complex64SliceOrig))
		goto complex128Slice
	}
	if &cpyC64 == &complex64SliceOrig {
		t.Error("[]complex64: address of copy was the same as original; they should be different")
		goto complex128Slice
	}
	for i, v := range complex64SliceOrig {
		if v != cpyC64[i] {
			t.Errorf("[]complex64: got %v at index %d of the copy; want %v", cpyC64[i], i, v)
		}
	}

complex128Slice:
	complex128SliceOrig := []complex128{complex128(-65 + 11i), complex128(66 + 10i)}
	cpy = Copy(complex128SliceOrig)
	cpyC128, ok := cpy.([]complex128)
	if !ok {
		t.Errorf("copy of []complex128: expected the interface to contain a []complex128; it didn't")
		goto arr4
	}
	if len(cpyC128) != len(complex128SliceOrig) {
		t.Errorf("[]complex128: len of copy was %d; want %d", len(cpyC128), len(complex128SliceOrig))
		goto arr4
	}
	if &cpyC128 == &complex128SliceOrig {
		t.Error("[]complex128: address of copy was the same as original; they should be different")
		goto arr4
	}
	for i, v := range complex128SliceOrig {
		if v != cpyC128[i] {
			t.Errorf("[]complex128: got %v at index %d of the copy; want %v", cpyC128[i], i, v)
		}
	}

arr4:
	arr4Array := [...]string{"a", "b", "c", "d"}
	cpy = Copy(arr4Array)
	cpyArr4, ok := cpy.([4]string)
	if !ok {
		t.Errorf("copy of [4]string: expected the interface to contain a [4]string; it didn't")
		goto done
	}
	if &cpyArr4 == &arr4Array {
		t.Error("[4]string: address of copy was the same as original; they should be different")
		goto done
	}
	for i, v := range arr4Array {
		if v != cpyArr4[i] {
			t.Errorf("[4]string: got %s at index %d of the copy; want %s", cpyArr4[i], i, v)
		}
	}

done:
}

// not meant to be exhaustive
func TestComplexSlices(t *testing.T) {
	orig3Int := [][][]int{[][]int{[]int{1, 2, 3}, []int{11, 22, 33}}, [][]int{[]int{7, 8, 9}, []int{66, 77, 88, 99}}}
	cpy := Copy(orig3Int)
	cpyI, ok := cpy.([][][]int)
	if !ok {
		t.Errorf("copy of [][][]int: expected the interface to contain a [][][]int; it didn't")
		goto sliceMap
	}
	if &cpyI == &orig3Int {
		t.Error("[][][]int: address of copy was the same as original; they should be different")
	}
	if len(orig3Int) != len(cpyI) {
		t.Errorf("[][][]int: len of copy was %d; want %d", len(cpyI), len(orig3Int))
		goto sliceMap
	}
	for i, v := range orig3Int {
		if len(v) != len(cpyI[i]) {
			t.Errorf("[][][]int: len of element %d was %d; want %d", i, len(cpyI[i]), len(v))
			continue
		}
		for j, vv := range v {
			if len(vv) != len(cpyI[i][j]) {
				t.Errorf("[][][]int: len of element %d:%d was %d, want %d", i, j, len(cpyI[i][j]), len(vv))
				continue
			}
			for k, vvv := range vv {
				if vvv != cpyI[i][j][k] {
					t.Errorf("[][][]int: element %d:%d:%d was %d, want %d", i, j, k, cpyI[i][j][k], vvv)
				}
			}
		}

	}

sliceMap:
	slMap := []map[int]string{map[int]string{0: "a", 1: "b"}, map[int]string{10: "k", 11: "l", 12: "m"}}
	cpy = Copy(slMap)
	cpyM, ok := cpy.([]map[int]string)
	if !ok {
		t.Errorf("copy of []map[int]string: expected the interface to contain a []map[int]string; it didn't")
		goto done
	}
	if &cpyM == &slMap {
		t.Error("[]map[int]string: address of copy was the same as original; they should be different")
	}
	if len(slMap) != len(cpyM) {
		t.Errorf("[]map[int]string: len of copy was %d; want %d", len(cpyM), len(slMap))
		goto done
	}
	for i, v := range slMap {
		if len(v) != len(cpyM[i]) {
			t.Errorf("[]map[int]string: len of element %d was %d; want %d", i, len(cpyM[i]), len(v))
			continue
		}
		for k, vv := range v {
			val, ok := cpyM[i][k]
			if !ok {
				t.Errorf("[]map[int]string: element %d was expected to have a value at key %d, it didn't", i, k)
				continue
			}
			if val != vv {
				t.Errorf("[]map[int]string: element %d, key %d: got %s, want %s", i, k, val, vv)
			}
		}
	}
done:
}

type A struct {
	Int    int
	String string
	UintSl []uint
	NilSl  []string
	Map    map[string]int
	MapB   map[string]*B
	SliceB []*B
	B
}

type B struct {
	Vals []string
}

var AStruct = A{
	Int:    42,
	String: "Konichiwa",
	UintSl: []uint{0, 1, 2, 3},
	Map:    map[string]int{"a": 1, "b": 2},
	MapB: map[string]*B{
		"hi":  &B{Vals: []string{"hello", "bonjour"}},
		"bye": &B{Vals: []string{"good-bye", "au revoir"}},
	},
	SliceB: []*B{
		&B{Vals: []string{"Ciao", "Aloha"}},
	},
	B: B{Vals: []string{"42"}},
}

func TestStructA(t *testing.T) {
	cpy := Copy(AStruct)
	a, ok := cpy.(A)
	if !ok {
		t.Error("expected copy to be of type AStruct, it wasn't")
		return
	}
	if &a == &AStruct {
		t.Error("expected copy to have a different address than the original; it was the same")
		return
	}
	if a.Int != AStruct.Int {
		t.Errorf("A.Int: got %v, want %v", a.Int, AStruct.Int)
	}
	if a.String != AStruct.String {
		t.Errorf("A.String: got %v; want %v", a.String, AStruct.String)
	}
	if fmt.Sprintf("%p", a.UintSl) == fmt.Sprintf("%p", AStruct.UintSl) {
		t.Error("A.Uintsl: expected the copies address to be different; it wasn't")
		goto AMap
	}
	if len(a.UintSl) != len(AStruct.UintSl) {
		t.Errorf("A.UintSl: got len of %d, want %d", len(a.UintSl), len(AStruct.UintSl))
		goto AMap
	}
	for i, v := range AStruct.UintSl {
		if a.UintSl[i] != v {
			t.Errorf("A.UintSl %d: got %d, want %d", i, a.UintSl[i], v)
		}
	}
AMap:
	if fmt.Sprintf("%p", a.Map) == fmt.Sprintf("%p", AStruct.Map) {
		t.Error("A.Map: expected the copy's address to be different; it wasn't")
		goto AMapB
	}
	if len(a.Map) != len(AStruct.Map) {
		t.Errorf("A.Map: got len of %d, want %d", len(a.Map), len(AStruct.Map))
		goto AMapB
	}
	for k, v := range AStruct.Map {
		val, ok := a.Map[k]
		if !ok {
			t.Errorf("A.Map: expected the key %s to exist in the copy, it didn't", k)
			continue
		}
		if val != v {
			t.Errorf("A.Map[%s]: got %d, want %d", k, val, v)
		}
	}

AMapB:
	if fmt.Sprintf("%p", a.MapB) == fmt.Sprintf("%p", AStruct.MapB) {
		t.Error("A.MapB: expected the copy's address to be different; it wasn't")
		goto ASliceB
	}
	if len(a.MapB) != len(AStruct.MapB) {
		t.Errorf("A.MapB: got len of %d, want %d", len(a.MapB), len(AStruct.MapB))
		goto ASliceB
	}
	for k, v := range AStruct.MapB {
		val, ok := a.MapB[k]
		if !ok {
			t.Errorf("A.MapB: expected the key %s to exist in the copy, it didn't", k)
			continue
		}
		if fmt.Sprintf("%p", val) == fmt.Sprintf("%p", v) {
			t.Errorf("A.MapB[%s]: expected the addresses of the values to be different; they weren't", k)
			continue
		}
		for i, vv := range v.Vals {
			if vv != val.Vals[i] {
				t.Errorf("A.MapB[%s].Vals[%d]: got %s want %s", k, i, vv, val.Vals[i])
			}
		}
	}
ASliceB:
	if fmt.Sprintf("%p", AStruct.SliceB) == fmt.Sprintf("%p", a.SliceB) {
		t.Error("A.SliceB: expected the copy's address to be different; it wasn't")
		goto B
	}

	if len(AStruct.SliceB) != len(a.SliceB) {
		t.Errorf("A.SliceB: got length of %d; want %d", len(a.SliceB), len(AStruct.SliceB))
		goto B
	}

	for i, v := range AStruct.SliceB {
		if fmt.Sprintf("%p", v) == fmt.Sprintf("%p", a.SliceB[i]) {
			t.Errorf("A.SliceB[%d]: expected them to have different addresses, they didn't", i)
			continue
		}
		if len(v.Vals) != len(a.SliceB[i].Vals) {
			t.Errorf("A.SliceB[%d]: expected B's vals to have the same length, they didn't", i)
			continue
		}
		for j, val := range v.Vals {
			if val != a.SliceB[i].Vals[j] {
				t.Errorf("A.SliceB[%d].Vals[%d]: got %v; want %v", i, j, a.SliceB[i].Vals[j], val)
			}
		}
	}
B:
	if fmt.Sprintf("%p", &AStruct.B) == fmt.Sprintf("%p", &a.B) {
		t.Error("A.B: expected them to have different addresses, they didn't")
		return
	}
	if fmt.Sprintf("%p", AStruct.B.Vals) == fmt.Sprintf("%p", a.B.Vals) {
		t.Error("A.B.Vals: expected them to have different addresses, they didn't")
		return
	}
	if len(AStruct.B.Vals) != len(a.B.Vals) {
		t.Error("A.B.Vals: expected their lengths to be the same, they weren't")
		return
	}
	for i, v := range AStruct.B.Vals {
		if v != a.B.Vals[i] {
			t.Errorf("A.B.Vals[%d]: got %s want %s", i, a.B.Vals[i], v)
		}
	}
}
