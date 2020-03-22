package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	mem := make(map[string][]int)
	return diffWay(input, mem)
}

// diffway compute the solution recursively
func diffWay(input string, mem map[string][]int) []int {
	// Check if this input already solved, return the answer
	v, ok := mem[input]
	if ok {
		return v
	}

	// create an answer slice to store answers
	ans := make([]int, 0)

	// break the expression
	for i := 0; i < len(input); i++ {
		op := rune(input[i])
		if op == '+' || op == '-' || op == '*' {

			left := input[0:i]
			right := input[i+1:]

			// get the solution recursively

			leftAns := diffWay(left, mem)
			rightAns := diffWay(right, mem)

			var fn operation

			switch op {
			case '+':
				fn = add
				break
			case '-':
				fn = minus
				break
			default:
				fn = multiply
			}

			// combine the solution to product results
			for _, a := range leftAns {
				for _, b := range rightAns {
					ans = append(ans, fn(a, b))
				}
			}

		}
	}

	// Single number
	if len(ans) == 0 {
		n, _ := strconv.Atoi(input)
		ans = append(ans, n)
	}

	// memorize the answer for input
	mem[input] = ans
	return mem[input]

}

// define operation type
type operation func(a, b int) int

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	input := "2-1-1"
	output := diffWaysToCompute(input)
	fmt.Println(output)
}
