package main

import "fmt"

func main() {
	var a int     // 0
	var b float64 // 0
	var c bool    // false
	var d string  // vazio
	var e *int    // nil

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
