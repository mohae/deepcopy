deepCopy
========
[![GoDoc](https://godoc.org/github.com/mohae/deepcopy?status.svg)](https://godoc.org/github.com/mohae/deepcopy)[![Build Status](https://travis-ci.org/mohae/deepcopy.png)](https://travis-ci.org/mohae/deepcopy)

DeepCopy makes deep copies of things: unexported field values are not copied.

## Usage
    cpy := deepcopy.Copy(orig)

## Tags

The following struct tags is supported:

```
type A struct {
    Field1 SomeType  `deepcopy:"-"` // skip, will have zero-value in copy
    Field2 *SomeType `deepcopy:"="` // treat like with "=" assignment operator
}
```

Specifically the `=` tag is usable when you want to copy a struct containing
something like `*sync.Mutex` or `*os.File` and don't want it to be deeply copied
but simply assigned.