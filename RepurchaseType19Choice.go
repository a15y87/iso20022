package iso20022

// Choice of format for the repurchase transaction type information.
type RepurchaseType19Choice struct {

	// Type of securities financing transaction process expressed as an ISO 20022 code.
	Code *RepurchaseType3Code `xml:"Cd"`

	// Type of securities financing transaction process expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`

}


func (r *RepurchaseType19Choice) SetCode(value string) {
	r.Code = (*RepurchaseType3Code)(&value)
}

func (r *RepurchaseType19Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}

