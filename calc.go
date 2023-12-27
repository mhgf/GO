package main

import "strconv"

func main() {
	sumR := sum(1, 2, 3)
	subR := sub(1, 2, 3)
	multR := mult(1, 2, 3)
	divR := div(10)

	println(sumR, subR, multR, divR)
}
func textFormat(test, corret int) string {
	return "O valor esperado: " + strconv.Itoa(corret) + ", Valor retornado: " + strconv.Itoa(test)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func sub(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = num - total
	}
	return total
}

func mult(nums ...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func div(nums ...int) int {
	total := 1
	for _, num := range nums {
		total = num / total
	}
	return total
}
