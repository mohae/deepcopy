// deepcopy deep copies maps, slices, etc. A standard copy will copy the
// pointers: deep copy copies the values pointed to.
// 
// Only what is needed has been implemented. Could make more dynamic, at the 
// cost of reflection. Either adjust as needed or create a new function.
// 
// Copyright (c)2014, Joel Scoble (github.com/mohae), all rights reserved.
// License: MIT, for more details check the included LICENSE.txt.
package deepcopy

import (
	"reflect"
)

var typeOfStrings = reflect.TypeOf([]string(nil))
var typeOfInts = reflect.TypeOf([]int(nil))

// InterfaceToSliceStrings takes an interface that is a slice of strings
// and returns a deep copy of it as a slice of strings.
func InterfaceToSliceStrings(v interface{}) []string {
	if v == nil {
		return nil
	}
	var sl []string

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		sLen := s.Len()

		for i := 0; i < sLen; i++ {
			sl = append(sl, s.Index(i).Interface().(string))
		}

	default:
		return nil
	}

	return sl
}

// SliceStrings deep copies a slice of strings
func SliceStrings(s []string) []string{
	if s == nil {
		return nil
	}
	
	var sl []string

	sLen := len(s)

	for i := 0; i < sLen; i++ {
		sl = append(sl, s[i])
	}

	return sl
}

// InterfaceToSliceInts takes an interface that is a slice of ints and returns 
// a deep copy of it as a slice of strings. An error is returned if the 
// interface is not a slice of strings.
func InterfaceToSliceInts(v interface{}) []int {
	if v == nil {
		return nil
	}
	var sl []int

	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		sLen := s.Len()

		for i := 0; i < sLen; i++ {
			sl = append(sl, s.Index(i).Interface().(int))
		}

	default:
		return nil
	}
	return sl
}

// SliceInts deep copies a slice of ints.
func SliceInts(s []int) []int{
	if s == nil {
		return nil
	}
	
	var sl []int

	sLen := len(s)

	for i := 0; i < sLen; i++ {
		sl = append(sl, s[i])
	}

	return sl
}

// MapStringInterface makes a deep copy of a map[string]interface{} and
// returns the copy of the map[string]interface{}
//
// Supported:
//	[]string
//	[]int
//	map[string]interface{}
// TODO convert embedded error string with an actual error, con, callee would
//	have to check for error. 
//	Otherwise add error return parm and abort conversion whenever an 
//	unexpected type is encountered. Probably the best option.
func MapStringInterface(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		return nil
	}

	c := map[string]interface{}{}
	var tmp interface{}

	for k, v := range m {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:

			switch reflect.ValueOf(v).Type() {
			case typeOfStrings:
				tmp = InterfaceToSliceStrings(v)
			case typeOfInts:
				tmp = InterfaceToSliceInts(v)
			default:
				tmp = "Error unsupported Type: " + reflect.ValueOf(v).Type().String() + " is not supported."
			}

		case reflect.Map:
			tmp = MapStringInterface(v.(map[string]interface{}))
		}
		c[k] = tmp
	}

	return c
}
