package maps

import "fmt"

func Run() {
	websites := map[string]string{
		"udemy":  "https://www.udemy.com",
		"google": "https://www.google.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["udemy"])

	websites["youtube"] = "https://www.youtube.com"
	fmt.Println(websites)

	delete(websites, "udemy")
	fmt.Println(websites)
}
