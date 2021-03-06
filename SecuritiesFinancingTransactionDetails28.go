package iso20022

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails28 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator3Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework3Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat4Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection21 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection21 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection21 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection21 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`

}


func (s *SecuritiesFinancingTransactionDetails28) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails28) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails28) AddRevaluation() *RevaluationIndicator3Choice {
	s.Revaluation = new(RevaluationIndicator3Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails28) AddLegalFramework() *LegalFramework3Choice {
	s.LegalFramework = new(LegalFramework3Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails28) AddInterestComputationMethod() *InterestComputationMethodFormat4Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat4Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails28) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails28) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails28) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails28) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails28) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails28) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails28) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails28) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails28) AddDealAmount() *AmountAndDirection21 {
	s.DealAmount = new(AmountAndDirection21)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddAccruedInterestAmount() *AmountAndDirection21 {
	s.AccruedInterestAmount = new(AmountAndDirection21)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddForfeitAmount() *AmountAndDirection21 {
	s.ForfeitAmount = new(AmountAndDirection21)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddPremiumAmount() *AmountAndDirection21 {
	s.PremiumAmount = new(AmountAndDirection21)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails28) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection21 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection21)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails28) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails28) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}

