//libadd.go
package main

import "C"

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
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
}
