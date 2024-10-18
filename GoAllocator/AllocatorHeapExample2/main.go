package main

import (
	"fmt"
)

func printValue(v interface{}) {
	fmt.Println(v)
	//	_, _ = v.(int)
}
/*
go build -gcflags '-l -m' 
-l = disable inlining
-m = print optimization decisions
*/
func main() {
	var num1 int = 10
	var str1 string = "Hello World"

	printValue(num1)
	printValue(str1)
	
	var num2 int = 20
	var str2 string = "Goodbye World"
	var i interface{}
	i = num2
	i = str2
	_ = i

}
