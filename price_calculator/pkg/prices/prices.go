package prices

import (
	"fmt"

	"example.com/price_calculator/pkg/conversion"
	"example.com/price_calculator/pkg/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		// fmt.Println(err)
		return err
	}
	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process() error {

	err := job.LoadPrices()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)

	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 50, 100},
		IOManager:   iom,
	}
}
