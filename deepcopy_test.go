package deepcopy

import (
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
		t.Errorf("len of copy was %d; want %d", len(cpyS), len(stringSliceOrig))
	}
	if &cpyS == &stringSliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range stringSliceOrig {
		if v != cpyS[i] {
			t.Errorf("got %s at index %d of the copy; want %s", cpyS[i], i, v)
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
		t.Errorf("len of copy was %d; want %d", len(cpyI), len(intSliceOrig))
	}
	if &cpyI == &intSliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range intSliceOrig {
		if v != cpyI[i] {
			t.Errorf("got %d at index %d of the copy; want %d", cpyI[i], i, v)
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
		t.Errorf("len of copy was %d; want %d", len(cpyI8), len(int8SliceOrig))
	}
	if &cpyI8 == &int8SliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range int8SliceOrig {
		if v != cpyI8[i] {
			t.Errorf("got %d at index %d of the copy; want %d", cpyI8[i], i, v)
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
		t.Errorf("len of copy was %d; want %d", len(cpyI16), len(int16SliceOrig))
	}
	if &cpyI16 == &int16SliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range int16SliceOrig {
		if v != cpyI16[i] {
			t.Errorf("got %d at index %d of the copy; want %d", cpyI16[i], i, v)
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
		t.Errorf("len of copy was %d; want %d", len(cpyI32), len(int32SliceOrig))
	}
	if &cpyI32 == &int32SliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range int32SliceOrig {
		if v != cpyI32[i] {
			t.Errorf("got %d at index %d of the copy; want %d", cpyI32[i], i, v)
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
		t.Errorf("len of copy was %d; want %d", len(cpyI64), len(int64SliceOrig))
	}
	if &cpyI64 == &int64SliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range int64SliceOrig {
		if v != cpyI64[i] {
			t.Errorf("got %d at index %d of the copy; want %d", cpyI64[i], i, v)
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
		t.Errorf("len of copy was %d; want %d", len(cpyF32), len(float32SliceOrig))
	}
	if &cpyF32 == &float32SliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range float32SliceOrig {
		if v != cpyF32[i] {
			t.Errorf("got %f at index %d of the copy; want %f", cpyF32[i], i, v)
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
		t.Errorf("len of copy was %d; want %d", len(cpyF64), len(float64SliceOrig))
	}
	if &cpyF64 == &float64SliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range float64SliceOrig {
		if v != cpyF64[i] {
			t.Errorf("got %f at index %d of the copy; want %f", cpyF64[i], i, v)
		}
	}

boolSlice:
	boolSliceOrig := []bool{true, false, false, true}
	cpy = Copy(boolSliceOrig)
	cpyB, ok := cpy.([]bool)
	if !ok {
		t.Errorf("copy of []int8: expected the interface to contain a []int8; it didn't")
		goto infSlice
	}
	if len(cpyB) != len(boolSliceOrig) {
		t.Errorf("len of copy was %d; want %d", len(cpyB), len(boolSliceOrig))
	}
	if &cpyB == &boolSliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range boolSliceOrig {
		if v != cpyB[i] {
			t.Errorf("got %t at index %d of the copy; want %t", cpyB[i], i, v)
		}
	}

infSlice:
	infSliceOrig := []interface{}{true, 1, "hello"}
	cpy = Copy(infSliceOrig)
	cpyInf, ok := cpy.([]interface{})
	if !ok {
		t.Errorf("copy of []int8: expected the interface to contain a []int8; it didn't")
		goto done
	}
	if len(cpyInf) != len(infSliceOrig) {
		t.Errorf("len of copy was %d; want %d", len(cpyInf), len(infSliceOrig))
	}
	if &cpyInf == &infSliceOrig {
		t.Error("address of copy was the same as original; they should be different")
	}
	for i, v := range infSliceOrig {
		if v != cpyInf[i] {
			t.Errorf("got %v at index %d of the copy; want %v", cpyInf[i], i, v)
		}
	}

done:
}
