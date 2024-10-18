package main

import "fmt"

// go run main.go
// GOGC=off go run main.GOGC

var data []byte

func main() {
	count := 0

	for {
		data = make([]byte, 1<<30) //GB
		//for idx := 0; idx < 1<<30; idx += 4096 {
		//	data[idx] = 100
		//}
		fmt.Println("Allocated GB:", count)
		count++
	}
}
