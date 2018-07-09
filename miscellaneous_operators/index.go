package main

import "fmt"

func main() {
	var a = 4
	var ptr *int

	ptr = &a /* 'ptr' now contains the address of 'a'*/
	fmt.Printf("value of a is %d\n", a)
	fmt.Printf("*ptr is %d.\n", *ptr)
}
