package main




func Examplefunc(number *int) int {
	result := *number + 200
	return result
}
/*
go build -gcflags '-l -m' 
-l = disable inlining
-m = print optimization decisions
*/
func main() {
	number := 100
	_ = Examplefunc(&number)
}
