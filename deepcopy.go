// deepcopy deep copies maps, slices, etc. A standard copy will copy the
// pointers: deep copy copies the values pointed to.
//
// For most use-cases, Copy should be used.  The others functions exist for
// backwards compatibility reasons.
//
// Copyright (c)2014, Joel Scoble (github.com/mohae), all rights reserved.
// License: MIT, for more details check the included LICENSE.txt.
package deepcopy

import (
	"reflect"
)

// InterfaceToStringSlice takes an interface that is a slice of strings
// and returns a deep copy of it as a slice of strings.
func InterfaceToStringSlice(v interface{}) []string {
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

// InterfaceToIntSlice takes an interface that is a slice of ints and returns
// a deep copy of it as a slice of strings. An error is returned if the
// interface is not a slice of strings.
func InterfaceToIntSlice(v interface{}) []int {
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

// Iface is an alias to Copy; this exists for backwards compatibility reasons.
func Iface(iface interface{}) interface{} {
	return Copy(iface)
}

// Copy creates a deep copy of whatever is passed to it and returns the copy
// in an interface{}.  The returned value will need to be asserted to the
// correct type.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Make the interface a reflect.Value
	original := reflect.ValueOf(src)

	// Make a copy of the same type as the original.
	cpy := reflect.New(original.Type()).Elem()

	// Recursively copy the original.
	copyRecursive(original, cpy)

	// Return theb copy as an interface.
	return cpy.Interface()
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original, cpy reflect.Value) {
	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()

		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return
		}

		cpy.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, cpy.Elem())

	case reflect.Interface:
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()

		// Get the value by calling Elem().
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		// Go through each field of the struct and copy it.
		for i := 0; i < original.NumField(); i++ {
			copyRecursive(original.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			copyRecursive(original.Index(i), cpy.Index(i))
		}

	case reflect.Map:
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			cpy.SetMapIndex(key, copyValue)
		}

	// Set the actual values from here on.
	case reflect.String:
		cpy.SetString(original.Interface().(string))

	case reflect.Int:
		cpy.SetInt(int64(original.Interface().(int)))

	case reflect.Bool:
		cpy.SetBool(original.Interface().(bool))

	case reflect.Float64:
		cpy.SetFloat(original.Interface().(float64))

	default:
		cpy.Set(original)
	}
}
