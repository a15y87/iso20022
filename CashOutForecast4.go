package iso20022

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast4 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

}


func (c *CashOutForecast4) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast4) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast4) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast4) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

