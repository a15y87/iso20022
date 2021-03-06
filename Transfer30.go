package iso20022

// Parameters applied to the settlement of a security transfer.
type Transfer30 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
	Quantity *Quantity13Choice `xml:"Qty"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit6 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`

}


func (t *Transfer30) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer30) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer30) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer30) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer30) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer30) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer30) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer30) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer30) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer30) AddQuantity() *Quantity13Choice {
	t.Quantity = new(Quantity13Choice)
	return t.Quantity
}

func (t *Transfer30) AddUnitsDetails() *Unit6 {
	newValue := new (Unit6)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer30) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer30) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer30) SetTransferCurrency(value string) {
	t.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (t *Transfer30) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer30) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer30) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer30) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}

func (t *Transfer30) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}

