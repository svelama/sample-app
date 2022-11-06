package main

import "C"

import (
	"runtime"
)

func main() {}

var X string = runtime.GOOS

//export GetX
func GetX() *C.char {
	return C.CString(X)
}

//export SetX
func SetX(os *C.char) {
	X = C.GoString(os)
}
