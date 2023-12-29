package lists

import "fmt"

func RunMake() {
	runSlices()
	fmt.Println("--------------------------------")
	runMaps()
}

func runSlices() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Joseph"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)
}

func runMaps() {
	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	fmt.Println(courseRatings)
}
