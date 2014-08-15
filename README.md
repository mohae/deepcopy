deepCopy
========

Deep Copy is useful when one wants more than the pointer.

Functions are implemented as needed.

## Currently Implemeted  
    InterfaceToSliceInts(v interface{}) []int
    InterfaceToSliceStrings(v interface{}) []string

    MapStringInterface(m map[string]interface{}) map[string]interface{}

    SliceInts(i []int) []int
    SliceStrings(s []string) []string

## Notes:
    __MapStringInterface()__ supports interfaces of:
	* []int
	* []string
	* map[string]interface{}
