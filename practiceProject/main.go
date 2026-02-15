package main

import (
	"fmt"

	"example.org/price-calculator/filemanager"
	"example.org/price-calculator/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}
	errorChannels1 := make(chan error, 2)

	for er := range errorChannels1 {
		fmt.Println(er)
	}
	doneChannels := make([]chan bool, len(taxRate))
	errorChannels := make([]chan error, len(taxRate))
	for index, taxValue := range taxRate {
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxValue*100))
		//cmdManager := cmdmanager.New()
		price := prices.NewTaxIncludedPriceJob(taxValue, fileManager)
		go price.Process(doneChannels[index], errorChannels[index])
		// if err != nil {
		// 	fmt.Println("Could not process the job")
		// 	fmt.Print(err)
		// }
	}

	for index := range taxRate {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[index]:
			fmt.Println("Done!")
		}
	}

	// for _, errorChan := range errorChannels {
	// 	<-errorChan
	// }
	// for _, doneChan := range doneChannels {
	// 	<-doneChan
	// }

}
