//libadd.go
package main

import "C"
import "fmt"

//export add
func add(left, right int) int {
	return left + right
}

//export subtract
func subtract(left, right int) int {
	return left - right
}

//export multiply
func multiply(left, right int) int {
	return left * right
}

//export divide
func divide(left, right int) int {
	return left / right
}

//export average
func average(C.array(int)) {
	return C.array()
}

func main() {
}
