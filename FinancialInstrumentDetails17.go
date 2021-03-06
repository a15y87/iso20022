package iso20022

// Reporting per financial instrument.
type FinancialInstrumentDetails17 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation6 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance1 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance1 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction36 `xml:"Tx"`

}


func (f *FinancialInstrumentDetails17) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails17) AddPriceDetails() *PriceInformation6 {
	f.PriceDetails = new(PriceInformation6)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails17) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	f.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails17) AddOpeningBalance() *OpeningBalance1 {
	f.OpeningBalance = new(OpeningBalance1)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails17) AddClosingBalance() *ClosingBalance1 {
	f.ClosingBalance = new(ClosingBalance1)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails17) AddTransaction() *Transaction36 {
	newValue := new (Transaction36)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

