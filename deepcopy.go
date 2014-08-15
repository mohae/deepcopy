// deepcopy deep copies maps, slices, etc. A standard copy will copy the
// pointers: deep copy copies the values pointed to.
// 
// Only what is needed has been implemented. Could make more dynamic, at the 
// cost of reflection. Either adjust as needed or create a new function.
package deepcopy

import (
	"reflect"
)

// InterfaceToSliceString takes an interface that is a slice of strings
// and returns a deep copy of it as a slice of strings. An error is returned if
// the interface is not a slice of strings
func InterfaceToSliceString(v interface{}) []string {
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

// SliceString deep copies a slice of strings
func SliceString(s []string) []string{
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

// MapStringInterface makes a deep copy of a map[string]interface{} and
// returns the copy of the map[string]interface{}
//
// notes: This assumes that the interface{} is a []string or another 
//	map[string]interface{}.
//	Adjust as needed.
func MapStringInterface(m map[string]interface{}) map[string]interface{} {
	c := map[string]interface{}{}
	var tmp interface{}

	for k, v := range m {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			tmp = InterfaceToSliceString(v)
		case reflect.Map:
			tmp = MapStringInterface(v.(map[string]interface{}))
		}
		c[k] = tmp
	}

	return c
}


