deepCopy
========
[![GoDoc](https://godoc.org/github.com/mohae/deepcopy?status.svg)](https://godoc.org/github.com/mohae/deepcopy)[![Build Status](https://travis-ci.org/mohae/deepcopy.png)](https://travis-ci.org/mohae/deepcopy)

DeepCopy makes deep copies of things: unexported field values are not copied.

## Usage

Use the `Copy()` function and assert the returned `interface{}` to the correct type.  Any other functions in this package exist for legacy reasons and should not be used.
