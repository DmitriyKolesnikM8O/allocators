package main

import (
	"fmt"
	"unsafe"
)

var (
	data1 [32769]byte
	data2 [32769]byte
)

/*
Explanation: for memory blocks larger than 32768
each of them is always composed of multiply memory pages.
The memory page size used by the official Golang standart runtime (1.22) is
8192 bytes.
That is, we lose 8 kilobytes even in such a small allocation of memory.
*/
func main() {
	data1Pointer := unsafe.Pointer(&data1)
	data2Pointer := unsafe.Pointer(&data2)

	fmt.Println("Address:", data1Pointer, data2Pointer)
	fmt.Println("Size:", unsafe.Sizeof(data1), unsafe.Sizeof(data2))

	distance := uintptr(data2Pointer) - uintptr(data1Pointer)
	fmt.Println("waste:", unsafe.Sizeof(distance))

}
