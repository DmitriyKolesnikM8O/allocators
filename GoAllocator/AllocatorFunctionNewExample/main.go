package main


func CreatePointer() *int {
	value2 := new(int)
	return value2
}
/*
go build -gcflags '-l -m' 
-l = disable inlining
-m = print optimization decisions
*/
func main() {
	value1 := new(int) // stack
	_ = value1

	value2 := CreatePointer() // heap
	_ = value2

}
