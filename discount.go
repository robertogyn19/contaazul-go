package contaazul

type DiscountMeasureUnit string

const (
	Percent DiscountMeasureUnit = "PERCENT"
	Value   DiscountMeasureUnit = "VALUE"
)

type Discount struct {
	MeasureUnit DiscountMeasureUnit `json:"measure_unit"`
	Rate        float64             `json:"rate"`
}
