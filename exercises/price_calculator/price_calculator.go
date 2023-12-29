package price_calculator

import (
	"fmt"

	"github.com/mheob/udemy-go/exercises/price_calculator/io"
	"github.com/mheob/udemy-go/exercises/price_calculator/prices"
)

const inputFile = "exercises/price_calculator/prices.txt"
const outputPath = "exercises/price_calculator/data"

func Run() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := io.NewFileManager(inputFile, fmt.Sprintf("%v/result_%.0f.json", outputPath, taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
