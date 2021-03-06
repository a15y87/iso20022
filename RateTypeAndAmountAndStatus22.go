package iso20022

// Specifies the value expressed as a rate and an amount.
type RateTypeAndAmountAndStatus22 struct {

	// Value expressed as a rate type.
	RateType *RateType40Choice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus3Choice `xml:"RateSts,omitempty"`

}


func (r *RateTypeAndAmountAndStatus22) AddRateType() *RateType40Choice {
	r.RateType = new(RateType40Choice)
	return r.RateType
}

func (r *RateTypeAndAmountAndStatus22) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateTypeAndAmountAndStatus22) AddRateStatus() *RateStatus3Choice {
	r.RateStatus = new(RateStatus3Choice)
	return r.RateStatus
}

