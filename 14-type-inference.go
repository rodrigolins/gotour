package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("Value %v Type %T\n", v, v)
	w := 42 + 10i
	fmt.Printf("Value %v Type %T\n", w, w)
	x := 42.0
	fmt.Printf("Value %v Type %T\n", x, x)
}
