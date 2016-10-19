deepCopy
========
[![GoDoc](https://godoc.org/github.com/mohae/deepcopy?status.svg)](https://godoc.org/github.com/mohae/deepcopy)[![Build Status](https://travis-ci.org/mohae/deepcopy.png)](https://travis-ci.org/mohae/deepcopy)

DeepCopy makes deep copies of things: unexported field values are not copied.

### Note
Copy of time will not copy all of the information because of https://github.com/golang/go/issues/15716.  The test to test deep copying time will fail until that issue is fixed and the version it is part of gets released.
