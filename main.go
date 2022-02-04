package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func main() {}

//export goRVExtensionVersion
func goRVExtensionVersion(output *C.char, outputsize C.size_t) {
	result := C.CString("arma_unix 1.0")
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func returnMyData(input string, errors error) string {
	switch input {
	case "arma_unix":
		return fmt.Sprintf("%d", time.Now().Unix())
	default:
		return "undefined command"
	}
	return ""
}

//export goRVExtension
func goRVExtension(output *C.char, outputsize C.size_t, input *C.char) {
	id := returnMyData(C.GoString(input), nil)
	printInArma(output, outputsize, id)
}

func printInArma(output *C.char, outputsize C.size_t, input string) {
	result := C.CString(input)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}
