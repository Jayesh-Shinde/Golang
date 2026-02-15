package prices

import (
	"errors"
	"fmt"
	"time"

	"example.org/price-calculator/converter"
	"example.org/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IoManager         iomanager.IOManager `json:"-"`
}

func NewTaxIncludedPriceJob(taxRate float64, ioManager iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{},
		IoManager:   ioManager,
	}
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()
	errorChan <- errors.New("an error occured!")
	if err != nil {
		errorChan <- err
		return
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	time.Sleep(3 * time.Second)
	job.IoManager.WriteJSON(job)
	doneChan <- true
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IoManager.Readlines()
	if err != nil {
		return err
	}
	job.InputPrices, err = converter.StringsToFloats(lines)
	if err != nil {
		return err
	}
	return nil
}
