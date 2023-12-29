package funcs

import "fmt"

func RunAnonymous() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
