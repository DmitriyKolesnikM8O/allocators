package main

import (
	"fmt"
)

/*
go build -gcflags '-l -m' 
-l = disable inlining
-m = print optimization decisions
*/

func Example(index int) byte {
	var data[1 << 20]byte
	return data[index]
}

func main() {
	var index int = 100

	pointer := &index
	fmt.Println(pointer)
	Example(index)
	fmt.Println(pointer)

}
