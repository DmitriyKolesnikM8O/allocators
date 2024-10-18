package main

func ExampleFunc() *int {
	result := 1000
	return &result
}

/*
go build -gcflags '-l -m' 
-l = disable inlining
-m = print optimization decisions
*/
func main() {
	_ = ExampleFunc()
}
