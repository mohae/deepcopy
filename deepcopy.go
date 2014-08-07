// deepcopy deep copies maps, slices, etc. A standard copy will copy the
// pointers: deep copy copies the values pointed to.
package deepcopy

import (
	"reflect"
)

// deepCopyInterfaceToSliceString takes an interface that is a slice of strings
// and returns a deep copy of it as a slice of strings. An error is returned if
// the interface is not a slice of strings
func deepCopyInterfaceToSliceString(v interface{}) []string {
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

// deepCopyMapStringInterface makes a deep copy of a map[string]interface{} and
// returns the copy of the map[string]interface{}
//
// notes: This assumes that the interface{} is a []string. Adjust as needed.
func deepCopyMapStringInterface(m map[string]interface{}) map[string]interface{} {
	c := map[string]interface{}{}
	var tmp  []string
	for k, v := range m {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			tmp = deepCopyInterfaceToSliceString(v)
		}
		c[k] = tmp
	}

	return c
}


