deepCopy
========
## Use github.com/mohae/utilitybelt/deepcopy instead
As of spring 2015, please do not use this repo, instead use the code at https://github.com/mohae/utilitybelt/tree/master/deepcopy as this is no longer actively maintained and the code within has been moved to `github.com/mohae/utilitybelt/deepcopy`.  Most of the slice copies in that version have been elided as they were written before I realized that there was a builtin `copy` func to copy slices.

Use:

    import github.com/mohae/utilitybelt/deepcopy

and:

    go get github.com/mohae/utilitybelt/deepcopy


## Legacy readme:
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
