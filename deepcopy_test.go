package deepcopy

import (
	"fmt"
	"testing"
	"time"
)

// just basic is this working stuff
func TestSimple(t *testing.T) {
	Strings := []string{"a", "b", "c"}
	cpy := Copy(Strings)
	cpyS, ok := cpy.([]string)
	if !ok {
		t.Errorf("copy []string: expected the interface to contain []string; it didn't")
		goto TestBools
	}
	if fmt.Sprintf("%p", cpyS) == fmt.Sprintf("%p", Strings) {
		t.Error("[]string: address of copy was the same as original; they should be different")
		goto TestBools
	}
	if len(cpyS) != len(Strings) {
		t.Errorf("[]string: len was %d; want %d", len(cpyS), len(Strings))
		goto TestBools
	}
	for i, v := range Strings {
		if v != cpyS[i] {
			t.Errorf("[]string: got %v at index %d of the copy; want %v", cpyS[i], i, v)
		}
	}

TestBools:
	Bools := []bool{true, true, false, false}
	cpy = Copy(Bools)
	cpyB, ok := cpy.([]bool)
	if !ok {
		t.Errorf("copy []bool: expected the interface to contain []bool; it didn't")
		goto TestBytes
	}
	if fmt.Sprintf("%p", cpyB) == fmt.Sprintf("%p", Bools) {
		t.Error("[]bool: address of copy was the same as original; they should be different")
		goto TestBytes
	}
	if len(cpyB) != len(Bools) {
		t.Errorf("[]bool: len was %d; want %d", len(cpyB), len(Bools))
		goto TestBytes
	}
	for i, v := range Bools {
		if v != cpyB[i] {
			t.Errorf("[]bool: got %v at index %d of the copy; want %v", cpyB[i], i, v)
		}
	}

TestBytes:
	Bytes := []byte("hello")
	cpy = Copy(Bytes)
	cpyBt, ok := cpy.([]byte)
	if !ok {
		t.Errorf("copy []byte: expected the interface to contain []byte; it didn't")
		goto CopyInts
	}
	if fmt.Sprintf("%p", cpyBt) == fmt.Sprintf("%p", Bytes) {
		t.Error("[]byte: address of copy was the same as original; they should be different")
		goto CopyInts
	}
	if len(cpyBt) != len(Bytes) {
		t.Errorf("[]byte: len was %d; want %d", len(cpyBt), len(Bytes))
		goto CopyInts
	}
	for i, v := range Bytes {
		if v != cpyBt[i] {
			t.Errorf("[]byte: got %v at index %d of the copy; want %v", cpyBt[i], i, v)
		}
	}

CopyInts:
	Ints := []int{42}
	cpy = Copy(Ints)
	cpyI, ok := cpy.([]int)
	if !ok {
		t.Errorf("copy []int: expected the interface to contain []int; it didn't")
		goto CopyUints
	}
	if fmt.Sprintf("%p", cpyI) == fmt.Sprintf("%p", Ints) {
		t.Error("[]int: address of copy was the same as original; they should be different")
		goto CopyUints
	}
	if len(cpyI) != len(Ints) {
		t.Errorf("[]int: len was %d; want %d", len(cpyI), len(Ints))
		goto CopyUints
	}
	for i, v := range Ints {
		if v != cpyI[i] {
			t.Errorf("[]int: got %v at index %d of the copy; want %v", cpyI[i], i, v)
		}
	}

CopyUints:
	Uints := []uint{1, 2, 3, 4, 5}
	cpy = Copy(Uints)
	cpyU, ok := cpy.([]uint)
	if !ok {
		t.Errorf("copy []uint: expected the interface to contain []uint; it didn't")
		goto CopyFloat32s
	}
	if fmt.Sprintf("%p", cpyU) == fmt.Sprintf("%p", Uints) {
		t.Error("[]uint: address of copy was the same as original; they should be different")
		goto CopyFloat32s
	}
	if len(cpyU) != len(Uints) {
		t.Errorf("[]uint: len was %d; want %d", len(cpyU), len(Uints))
		goto CopyFloat32s
	}
	for i, v := range Uints {
		if v != cpyU[i] {
			t.Errorf("[]uint: got %v at index %d of the copy; want %v", cpyU[i], i, v)
		}
	}

CopyFloat32s:
	Float32s := []float32{3.14}
	cpy = Copy(Float32s)
	cpyF, ok := cpy.([]float32)
	if !ok {
		t.Errorf("copy []float32: expected the interface to contain []float32; it didn't")
		goto CopyInterfaces
	}
	if fmt.Sprintf("%p", cpyF) == fmt.Sprintf("%p", Float32s) {
		t.Error("[]float32: address of copy was the same as original; they should be different")
		goto CopyInterfaces
	}
	if len(cpyF) != len(Float32s) {
		t.Errorf("[]float32: len was %d; want %d", len(cpyF), len(Float32s))
		goto CopyInterfaces
	}
	for i, v := range Float32s {
		if v != cpyF[i] {
			t.Errorf("[]float32: got %v at index %d of the copy; want %v", cpyF[i], i, v)
		}
	}

CopyInterfaces:
	Interfaces := []interface{}{"a", 42, true, 4.32}
	cpy = Copy(Interfaces)
	cpyIf, ok := cpy.([]interface{})
	if !ok {
		t.Errorf("copy []interface{}: expected the interface to contain []interface{}; it didn't")
		return
	}
	if fmt.Sprintf("%p", cpyIf) == fmt.Sprintf("%p", Interfaces) {
		t.Error("[]interface{}: address of copy was the same as original; they should be different")
		return
	}
	if len(cpyIf) != len(Interfaces) {
		t.Errorf("[]interface{}: len was %d; want %d", len(cpyIf), len(Interfaces))
		return
	}
	for i, v := range Interfaces {
		if v != cpyIf[i] {
			t.Errorf("[]interface{}: got %v at index %d of the copy; want %v", cpyIf[i], i, v)
		}
	}
}

type Basics struct {
	String      string
	Strings     []string
	StringArr   [4]string
	Bool        bool
	Bools       []bool
	Byte        byte
	Bytes       []byte
	Int         int
	Ints        []int
	Int8        int8
	Int8s       []int8
	Int16       int16
	Int16s      []int16
	Int32       int32
	Int32s      []int32
	Int64       int64
	Int64s      []int64
	Uint        uint
	Uints       []uint
	Uint8       uint8
	Uint8s      []uint8
	Uint16      uint16
	Uint16s     []uint16
	Uint32      uint32
	Uint32s     []uint32
	Uint64      uint64
	Uint64s     []uint64
	Float32     float32
	Float32s    []float32
	Float64     float64
	Float64s    []float64
	Complex64   complex64
	Complex64s  []complex64
	Complex128  complex128
	Complex128s []complex128
	Interface   interface{}
	Interfaces  []interface{}
}

// These tests test that all supported basic types are copied correctly.  This
// is done by copying a struct with fields of most of the basic types as []T.
func TestMostTypes(t *testing.T) {
	test := Basics{
		String:      "kimchi",
		Strings:     []string{"uni", "ika"},
		StringArr:   [4]string{"malort", "barenjager", "fernet", "salmiakki"},
		Bool:        true,
		Bools:       []bool{true, false, true},
		Byte:        'z',
		Bytes:       []byte("abc"),
		Int:         42,
		Ints:        []int{0, 1, 3, 4},
		Int8:        8,
		Int8s:       []int8{8, 9, 10},
		Int16:       16,
		Int16s:      []int16{16, 17, 18, 19},
		Int32:       32,
		Int32s:      []int32{32, 33},
		Int64:       64,
		Int64s:      []int64{64},
		Uint:        420,
		Uints:       []uint{11, 12, 13},
		Uint8:       81,
		Uint8s:      []uint8{81, 82},
		Uint16:      160,
		Uint16s:     []uint16{160, 161, 162, 163, 164},
		Uint32:      320,
		Uint32s:     []uint32{320, 321},
		Uint64:      640,
		Uint64s:     []uint64{6400, 6401, 6402, 6403},
		Float32:     32.32,
		Float32s:    []float32{32.32, 33},
		Float64:     64.1,
		Float64s:    []float64{64, 65, 66},
		Complex64:   complex64(-64 + 12i),
		Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
		Complex128:  complex128(-128 + 12i),
		Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
		Interfaces:  []interface{}{42, true, "pan-galactic"},
	}

	cpy := Copy(test)
	basic, ok := cpy.(Basics)
	if !ok {
		t.Errorf("copy of Basics: expected the interface to contain a Basics struct; it didn't")
		return
	}

	// see if they point to the same location
	if fmt.Sprintf("%p", &basic) == fmt.Sprintf("%p", &test) {
		t.Error("address of copy was the same as original; they should be different")
		return
	}

	// Go through each field and check to see it got copied properly
	if basic.String != test.String {
		t.Errorf("String: got %v; want %v", basic.String, test.String)
	}
	if fmt.Sprintf("%p", basic.Strings) == fmt.Sprintf("%p", test.Strings) {
		t.Error("Strings: address of copy was the same as original; they should be different")
		goto StringArr
	}
	if len(basic.Strings) != len(test.Strings) {
		t.Errorf("Strings: len was %d; want %d", len(basic.Strings), len(test.Strings))
		goto StringArr
	}
	for i, v := range test.Strings {
		if v != basic.Strings[i] {
			t.Errorf("Strings: got %v at index %d of the copy; want %v", basic.Strings[i], i, v)
		}
	}

StringArr:
	if fmt.Sprintf("%p", &basic.StringArr) == fmt.Sprintf("%p", &test.StringArr) {
		t.Error("StringArr: address of copy was the same as original; they should be different")
		goto Bools
	}
	for i, v := range test.StringArr {
		if v != basic.StringArr[i] {
			t.Errorf("StringArr: got %v at index %d of the copy; want %v", basic.StringArr[i], i, v)
		}
	}

Bools:
	if basic.Bool != test.Bool {
		t.Errorf("Bool: got %v; want %v", basic.Bool, test.Bool)
	}
	if fmt.Sprintf("%p", basic.Bools) == fmt.Sprintf("%p", test.Bools) {
		t.Error("Bools: address of copy was the same as original; they should be different")
		goto Bytes
	}
	if len(basic.Bools) != len(test.Bools) {
		t.Errorf("Bools: len was %d; want %d", len(basic.Bools), len(test.Bools))
		goto Bytes
	}
	for i, v := range test.Bools {
		if v != basic.Bools[i] {
			t.Errorf("Bools: got %v at index %d of the copy; want %v", basic.Bools[i], i, v)
		}
	}

Bytes:
	if basic.Byte != test.Byte {
		t.Errorf("Byte: got %v; want %v", basic.Byte, test.Byte)
	}
	if fmt.Sprintf("%p", basic.Bytes) == fmt.Sprintf("%p", test.Bytes) {
		t.Error("Bytes: address of copy was the same as original; they should be different")
		goto Ints
	}
	if len(basic.Bytes) != len(test.Bytes) {
		t.Errorf("Bytes: len was %d; want %d", len(basic.Bytes), len(test.Bytes))
		goto Ints
	}
	for i, v := range test.Bytes {
		if v != basic.Bytes[i] {
			t.Errorf("Bytes: got %v at index %d of the copy; want %v", basic.Bytes[i], i, v)
		}
	}

Ints:
	if basic.Int != test.Int {
		t.Errorf("Int: got %v; want %v", basic.Int, test.Int)
	}
	if fmt.Sprintf("%p", basic.Ints) == fmt.Sprintf("%p", test.Ints) {
		t.Error("Ints: address of copy was the same as original; they should be different")
		goto Int8s
	}
	if len(basic.Ints) != len(test.Ints) {
		t.Errorf("Ints: len was %d; want %d", len(basic.Ints), len(test.Ints))
		goto Int8s
	}
	for i, v := range test.Ints {
		if v != basic.Ints[i] {
			t.Errorf("Ints: got %v at index %d of the copy; want %v", basic.Ints[i], i, v)
		}
	}

Int8s:
	if basic.Int8 != test.Int8 {
		t.Errorf("Int8: got %v; want %v", basic.Int8, test.Int8)
	}
	if fmt.Sprintf("%p", basic.Int8s) == fmt.Sprintf("%p", test.Int8s) {
		t.Error("Int8s: address of copy was the same as original; they should be different")
		goto Int16s
	}
	if len(basic.Int8s) != len(test.Int8s) {
		t.Errorf("Int8s: len was %d; want %d", len(basic.Int8s), len(test.Int8s))
		goto Int16s
	}
	for i, v := range test.Int8s {
		if v != basic.Int8s[i] {
			t.Errorf("Int8s: got %v at index %d of the copy; want %v", basic.Int8s[i], i, v)
		}
	}

Int16s:
	if basic.Int16 != test.Int16 {
		t.Errorf("Int16: got %v; want %v", basic.Int16, test.Int16)
	}
	if fmt.Sprintf("%p", basic.Int16s) == fmt.Sprintf("%p", test.Int16s) {
		t.Error("Int16s: address of copy was the same as original; they should be different")
		goto Int32s
	}
	if len(basic.Int16s) != len(test.Int16s) {
		t.Errorf("Int16s: len was %d; want %d", len(basic.Int16s), len(test.Int16s))
		goto Int32s
	}
	for i, v := range test.Int16s {
		if v != basic.Int16s[i] {
			t.Errorf("Int16s: got %v at index %d of the copy; want %v", basic.Int16s[i], i, v)
		}
	}

Int32s:
	if basic.Int32 != test.Int32 {
		t.Errorf("Int32: got %v; want %v", basic.Int32, test.Int32)
	}
	if fmt.Sprintf("%p", basic.Int32s) == fmt.Sprintf("%p", test.Int32s) {
		t.Error("Int32s: address of copy was the same as original; they should be different")
		goto Int64s
	}
	if len(basic.Int32s) != len(test.Int32s) {
		t.Errorf("Int32s: len was %d; want %d", len(basic.Int32s), len(test.Int32s))
		goto Int64s
	}
	for i, v := range test.Int32s {
		if v != basic.Int32s[i] {
			t.Errorf("Int32s: got %v at index %d of the copy; want %v", basic.Int32s[i], i, v)
		}
	}

Int64s:
	if basic.Int64 != test.Int64 {
		t.Errorf("Int64: got %v; want %v", basic.Int64, test.Int64)
	}
	if fmt.Sprintf("%p", basic.Int64s) == fmt.Sprintf("%p", test.Int64s) {
		t.Error("Int64s: address of copy was the same as original; they should be different")
		goto Uints
	}
	if len(basic.Int64s) != len(test.Int64s) {
		t.Errorf("Int64s: len was %d; want %d", len(basic.Int64s), len(test.Int64s))
		goto Uints
	}
	for i, v := range test.Int64s {
		if v != basic.Int64s[i] {
			t.Errorf("Int64s: got %v at index %d of the copy; want %v", basic.Int64s[i], i, v)
		}
	}

Uints:
	if basic.Uint != test.Uint {
		t.Errorf("Uint: got %v; want %v", basic.Uint, test.Uint)
	}
	if fmt.Sprintf("%p", basic.Uints) == fmt.Sprintf("%p", test.Uints) {
		t.Error("Uints: address of copy was the same as original; they should be different")
		goto Uint8s
	}
	if len(basic.Uints) != len(test.Uints) {
		t.Errorf("Uints: len was %d; want %d", len(basic.Uints), len(test.Uints))
		goto Uint8s
	}
	for i, v := range test.Uints {
		if v != basic.Uints[i] {
			t.Errorf("Uints: got %v at index %d of the copy; want %v", basic.Uints[i], i, v)
		}
	}

Uint8s:
	if basic.Uint8 != test.Uint8 {
		t.Errorf("Uint8: got %v; want %v", basic.Uint8, test.Uint8)
	}
	if fmt.Sprintf("%p", basic.Uint8s) == fmt.Sprintf("%p", test.Uint8s) {
		t.Error("Uint8s: address of copy was the same as original; they should be different")
		goto Uint16s
	}
	if len(basic.Uint8s) != len(test.Uint8s) {
		t.Errorf("Uint8s: len was %d; want %d", len(basic.Uint8s), len(test.Uint8s))
		goto Uint16s
	}
	for i, v := range test.Uint8s {
		if v != basic.Uint8s[i] {
			t.Errorf("Uint8s: got %v at index %d of the copy; want %v", basic.Uint8s[i], i, v)
		}
	}

Uint16s:
	if basic.Uint16 != test.Uint16 {
		t.Errorf("Uint16: got %v; want %v", basic.Uint16, test.Uint16)
	}
	if fmt.Sprintf("%p", basic.Uint16s) == fmt.Sprintf("%p", test.Uint16s) {
		t.Error("Uint16s: address of copy was the same as original; they should be different")
		goto Uint32s
	}
	if len(basic.Uint16s) != len(test.Uint16s) {
		t.Errorf("Uint16s: len was %d; want %d", len(basic.Uint16s), len(test.Uint16s))
		goto Uint32s
	}
	for i, v := range test.Uint16s {
		if v != basic.Uint16s[i] {
			t.Errorf("Uint16s: got %v at index %d of the copy; want %v", basic.Uint16s[i], i, v)
		}
	}

Uint32s:
	if basic.Uint32 != test.Uint32 {
		t.Errorf("Uint32: got %v; want %v", basic.Uint32, test.Uint32)
	}
	if fmt.Sprintf("%p", basic.Uint32s) == fmt.Sprintf("%p", test.Uint32s) {
		t.Error("Uint32s: address of copy was the same as original; they should be different")
		goto Uint64s
	}
	if len(basic.Uint32s) != len(test.Uint32s) {
		t.Errorf("Uint32s: len was %d; want %d", len(basic.Uint32s), len(test.Uint32s))
		goto Uint64s
	}
	for i, v := range test.Uint32s {
		if v != basic.Uint32s[i] {
			t.Errorf("Uint32s: got %v at index %d of the copy; want %v", basic.Uint32s[i], i, v)
		}
	}

Uint64s:
	if basic.Uint64 != test.Uint64 {
		t.Errorf("Uint64: got %v; want %v", basic.Uint64, test.Uint64)
	}
	if fmt.Sprintf("%p", basic.Uint64s) == fmt.Sprintf("%p", test.Uint64s) {
		t.Error("Uint64s: address of copy was the same as original; they should be different")
		goto Float32s
	}
	if len(basic.Uint64s) != len(test.Uint64s) {
		t.Errorf("Uint64s: len was %d; want %d", len(basic.Uint64s), len(test.Uint64s))
		goto Float32s
	}
	for i, v := range test.Uint64s {
		if v != basic.Uint64s[i] {
			t.Errorf("Uint64s: got %v at index %d of the copy; want %v", basic.Uint64s[i], i, v)
		}
	}

Float32s:
	if basic.Float32 != test.Float32 {
		t.Errorf("Float32: got %v; want %v", basic.Float32, test.Float32)
	}
	if fmt.Sprintf("%p", basic.Float32s) == fmt.Sprintf("%p", test.Float32s) {
		t.Error("Float32s: address of copy was the same as original; they should be different")
		goto Float64s
	}
	if len(basic.Float32s) != len(test.Float32s) {
		t.Errorf("Float32s: len was %d; want %d", len(basic.Float32s), len(test.Float32s))
		goto Float64s
	}
	for i, v := range test.Float32s {
		if v != basic.Float32s[i] {
			t.Errorf("Float32s: got %v at index %d of the copy; want %v", basic.Float32s[i], i, v)
		}
	}

Float64s:
	if basic.Float64 != test.Float64 {
		t.Errorf("Float64: got %v; want %v", basic.Float64, test.Float64)
	}
	if fmt.Sprintf("%p", basic.Float64s) == fmt.Sprintf("%p", test.Float64s) {
		t.Error("Float64s: address of copy was the same as original; they should be different")
		goto Complex64s
	}
	if len(basic.Float64s) != len(test.Float64s) {
		t.Errorf("Float64s: len was %d; want %d", len(basic.Float64s), len(test.Float64s))
		goto Complex64s
	}
	for i, v := range test.Float64s {
		if v != basic.Float64s[i] {
			t.Errorf("Float64s: got %v at index %d of the copy; want %v", basic.Float64s[i], i, v)
		}
	}

Complex64s:
	if basic.Complex64 != test.Complex64 {
		t.Errorf("Complex64: got %v; want %v", basic.Complex64, test.Complex64)
	}
	if fmt.Sprintf("%p", basic.Complex64s) == fmt.Sprintf("%p", test.Complex64s) {
		t.Error("Complex64s: address of copy was the same as original; they should be different")
		goto Complex128s
	}
	if len(basic.Complex64s) != len(test.Complex64s) {
		t.Errorf("Complex64s: len was %d; want %d", len(basic.Complex64s), len(test.Complex64s))
		goto Complex128s
	}
	for i, v := range test.Complex64s {
		if v != basic.Complex64s[i] {
			t.Errorf("Complex64s: got %v at index %d of the copy; want %v", basic.Complex64s[i], i, v)
		}
	}

Complex128s:
	if basic.Complex128 != test.Complex128 {
		t.Errorf("Complex128s: got %v; want %v", basic.Complex128s, test.Complex128s)
	}
	if &(basic.Complex128s) == &(test.Complex128s) {
		t.Error("Complex128s: address of copy was the same as original; they should be different")
		goto Interfaces
	}
	if len(basic.Complex128s) != len(test.Complex128s) {
		t.Errorf("Complex128s: len was %d; want %d", len(basic.Complex128s), len(test.Complex128s))
		goto Interfaces
	}
	for i, v := range test.Complex128s {
		if v != basic.Complex128s[i] {
			t.Errorf("Complex128s: got %v at index %d of the copy; want %v", basic.Complex128s[i], i, v)
		}
	}

Interfaces:
	if basic.Interface != test.Interface {
		t.Errorf("Interface: got %v; want %v", basic.Interface, test.Interface)
	}
	if fmt.Sprintf("%p", basic.Interfaces) == fmt.Sprintf("%p", test.Interfaces) {
		t.Error("Interfaces: address of copy was the same as original; they should be different")
		return
	}
	if len(basic.Interfaces) != len(test.Interfaces) {
		t.Errorf("Interfaces: len was %d; want %d", len(basic.Interfaces), len(test.Interfaces))
		return
	}
	for i, v := range test.Interfaces {
		if v != basic.Interfaces[i] {
			t.Errorf("Interfaces: got %v at index %d of the copy; want %v", basic.Interfaces[i], i, v)
		}
	}
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
	if fmt.Sprintf("%p", cpyI) == fmt.Sprintf("%p", orig3Int) {
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
	if fmt.Sprintf("%p", cpyM) == fmt.Sprintf("%p", slMap) {
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
	NilMap map[string]string
	Map    map[string]int
	MapB   map[string]*B
	SliceB []*B
	B
	T time.Time
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
	T: time.Now(),
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
	if a.NilSl != nil {
		t.Errorf("A.NilSl: nil slice expected")
		goto AMap
	}
	if a.NilMap != nil {
		t.Errorf("A.NilSl: nil map expected")
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
		goto T
	}
	if fmt.Sprintf("%p", AStruct.B.Vals) == fmt.Sprintf("%p", a.B.Vals) {
		t.Error("A.B.Vals: expected them to have different addresses, they didn't")
		goto T
	}
	if len(AStruct.B.Vals) != len(a.B.Vals) {
		t.Error("A.B.Vals: expected their lengths to be the same, they weren't")
		goto T
	}
	for i, v := range AStruct.B.Vals {
		if v != a.B.Vals[i] {
			t.Errorf("A.B.Vals[%d]: got %s want %s", i, a.B.Vals[i], v)
		}
	}
T:
	if fmt.Sprintf("%p", &AStruct.T) == fmt.Sprintf("%p", &a.T) {
		t.Error("A.T: expected them to have different addresses, they didn't")
		return
	}
	if AStruct.T != a.T {
		t.Errorf("A.T: got %v, want %v", a.T, AStruct.T)
	}
}

type Unexported struct {
	A  string
	B  int
	aa string
	bb int
	cc []int
	dd map[string]string
}

func TestUnexportedFields(t *testing.T) {
	u := &Unexported{
		A:  "A",
		B:  42,
		aa: "aa",
		bb: 42,
		cc: []int{1, 2, 3},
		dd: map[string]string{"hello": "bonjour"},
	}
	cpy := Copy(u)
	v := cpy.(*Unexported)
	if v == u {
		t.Error("expected addresses to be different, they weren't")
		return
	}
	if u.A != v.A {
		t.Errorf("Unexported.A: got %s want %s", v.A, u.A)
	}
	if u.B != v.B {
		t.Errorf("Unexported.A: got %d want %d", v.B, u.B)
	}
	if v.aa != "" {
		t.Errorf("Unexported.aa: unexported field should not be set, it was set to %s", v.aa)
	}
	if v.bb != 0 {
		t.Errorf("Unexported.bb: unexported field should not be set, it was set to %d", v.bb)
	}
	if v.cc != nil {
		t.Errorf("Unexported.cc: unexported field should not be set, it was set to %#v", v.cc)
	}
	if v.dd != nil {
		t.Errorf("Unexported.dd: unexported field should not be set, it was set to %#v", v.dd)
	}
}
