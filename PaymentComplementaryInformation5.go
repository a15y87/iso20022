package iso20022

// Provides further additional details on the underlying payment instruction that cannot be transferred in a regular statement message.
type PaymentComplementaryInformation5 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	// 
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	// 
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	// 
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain. 
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. 
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. 
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType4Choice `xml:"Amt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementInformation *SettlementInstruction4 `xml:"SttlmInf,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	// 
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntermediaryAgent1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	// 
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntermediaryAgent2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	// 
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntermediaryAgent3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntermediaryAgent3Account *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Agent immediately prior to the instructing agent.
	PreviousInstructingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"PrvsInstgAgt,omitempty"`

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PreviousInstructingAgentAccount *CashAccount24 `xml:"PrvsInstgAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent. 
	// 
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`

}


func (p *PaymentComplementaryInformation5) SetInstructionIdentification(value string) {
	p.InstructionIdentification = (*Max35Text)(&value)
}

func (p *PaymentComplementaryInformation5) SetEndToEndIdentification(value string) {
	p.EndToEndIdentification = (*Max35Text)(&value)
}

func (p *PaymentComplementaryInformation5) SetTransactionIdentification(value string) {
	p.TransactionIdentification = (*Max35Text)(&value)
}

func (p *PaymentComplementaryInformation5) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	p.PaymentTypeInformation = new(PaymentTypeInformation25)
	return p.PaymentTypeInformation
}

func (p *PaymentComplementaryInformation5) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentComplementaryInformation5) SetRequestedCollectionDate(value string) {
	p.RequestedCollectionDate = (*ISODate)(&value)
}

func (p *PaymentComplementaryInformation5) SetInterbankSettlementDate(value string) {
	p.InterbankSettlementDate = (*ISODate)(&value)
}

func (p *PaymentComplementaryInformation5) AddAmount() *AmountType4Choice {
	p.Amount = new(AmountType4Choice)
	return p.Amount
}

func (p *PaymentComplementaryInformation5) SetInterbankSettlementAmount(value, currency string) {
	p.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentComplementaryInformation5) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentComplementaryInformation5) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentComplementaryInformation5) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentComplementaryInformation5) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentComplementaryInformation5) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentComplementaryInformation5) AddDebtorAgentAccount() *CashAccount24 {
	p.DebtorAgentAccount = new(CashAccount24)
	return p.DebtorAgentAccount
}

func (p *PaymentComplementaryInformation5) AddSettlementInformation() *SettlementInstruction4 {
	p.SettlementInformation = new(SettlementInstruction4)
	return p.SettlementInformation
}

func (p *PaymentComplementaryInformation5) AddIntermediaryAgent1() *BranchAndFinancialInstitutionIdentification5 {
	p.IntermediaryAgent1 = new(BranchAndFinancialInstitutionIdentification5)
	return p.IntermediaryAgent1
}

func (p *PaymentComplementaryInformation5) AddIntermediaryAgent1Account() *CashAccount24 {
	p.IntermediaryAgent1Account = new(CashAccount24)
	return p.IntermediaryAgent1Account
}

func (p *PaymentComplementaryInformation5) AddIntermediaryAgent2() *BranchAndFinancialInstitutionIdentification5 {
	p.IntermediaryAgent2 = new(BranchAndFinancialInstitutionIdentification5)
	return p.IntermediaryAgent2
}

func (p *PaymentComplementaryInformation5) AddIntermediaryAgent2Account() *CashAccount24 {
	p.IntermediaryAgent2Account = new(CashAccount24)
	return p.IntermediaryAgent2Account
}

func (p *PaymentComplementaryInformation5) AddIntermediaryAgent3() *BranchAndFinancialInstitutionIdentification5 {
	p.IntermediaryAgent3 = new(BranchAndFinancialInstitutionIdentification5)
	return p.IntermediaryAgent3
}

func (p *PaymentComplementaryInformation5) AddIntermediaryAgent3Account() *CashAccount24 {
	p.IntermediaryAgent3Account = new(CashAccount24)
	return p.IntermediaryAgent3Account
}

func (p *PaymentComplementaryInformation5) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.CreditorAgent
}

func (p *PaymentComplementaryInformation5) AddCreditorAgentAccount() *CashAccount24 {
	p.CreditorAgentAccount = new(CashAccount24)
	return p.CreditorAgentAccount
}

func (p *PaymentComplementaryInformation5) AddCreditor() *PartyIdentification43 {
	p.Creditor = new(PartyIdentification43)
	return p.Creditor
}

func (p *PaymentComplementaryInformation5) AddCreditorAccount() *CashAccount24 {
	p.CreditorAccount = new(CashAccount24)
	return p.CreditorAccount
}

func (p *PaymentComplementaryInformation5) AddUltimateCreditor() *PartyIdentification43 {
	p.UltimateCreditor = new(PartyIdentification43)
	return p.UltimateCreditor
}

func (p *PaymentComplementaryInformation5) AddPurpose() *Purpose2Choice {
	p.Purpose = new(Purpose2Choice)
	return p.Purpose
}

func (p *PaymentComplementaryInformation5) SetInstructionForDebtorAgent(value string) {
	p.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (p *PaymentComplementaryInformation5) AddPreviousInstructingAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.PreviousInstructingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.PreviousInstructingAgent
}

func (p *PaymentComplementaryInformation5) AddPreviousInstructingAgentAccount() *CashAccount24 {
	p.PreviousInstructingAgentAccount = new(CashAccount24)
	return p.PreviousInstructingAgentAccount
}

func (p *PaymentComplementaryInformation5) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new (InstructionForNextAgent1)
	p.InstructionForNextAgent = append(p.InstructionForNextAgent, newValue)
	return newValue
}

func (p *PaymentComplementaryInformation5) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new (InstructionForCreditorAgent1)
	p.InstructionForCreditorAgent = append(p.InstructionForCreditorAgent, newValue)
	return newValue
}

func (p *PaymentComplementaryInformation5) AddRemittanceInformation() *RemittanceInformation11 {
	p.RemittanceInformation = new(RemittanceInformation11)
	return p.RemittanceInformation
}

