package funcs

import "fmt"

func RunVariadic() {
	numbers := []int{1, 10, 15}
	sum := sumUp(1, 10, 15, 40, -5)
	anotherSum := sumUp(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumUp(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}
