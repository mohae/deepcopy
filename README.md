deepCopy
========

Deep Copy is useful when one wants more than the pointer.

Functions are implemented as needed.

## Currently Implemeted  
    Iface(iface interface{}) interface{}

    InterfaceToSliceInts(v interface{}) []int
    InterfaceToSliceStrings(v interface{}) []string

    MapStringInterface(m map[string]interface{}) map[string]interface{}

    SliceInts(i []int) []int
    SliceStrings(s []string) []string

## Notes:
    __Iface__ will take almost any interface{} and return a deep copy of it as an interface{}. This handles structs too. Please check the `copyRecursive` function to see details of what is handled.
    __MapStringInterface()__ supports interfaces of:
	* []int
	* []string
	* map[string]interface{}
