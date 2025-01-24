/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESINTERNATIONAL_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesInternational struct {
	NotesStruct
}

func newNotesInternational(dispatchPtr *ole.IDispatch) NotesInternational {
	return NotesInternational{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_AMSTRING_PROPERTY.html */
func (i NotesInternational) AMString() (String, error) {
	return getComProperty[String](i, "AMString")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENCYDIGITS_PROPERTY.html */
func (i NotesInternational) CurrencyDigits() (Long, error) {
	return getComProperty[Long](i, "CurrencyDigits")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENCYSYMBOL_PROPERTY.html */
func (i NotesInternational) CurrencySymbol() (String, error) {
	return getComProperty[String](i, "CurrencySymbol")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATESEP_PROPERTY.html */
func (i NotesInternational) DateSep() (String, error) {
	return getComProperty[String](i, "DateSep")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECIMALSEP_PROPERTY.html */
func (i NotesInternational) DecimalSep() (String, error) {
	return getComProperty[String](i, "DecimalSep")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENCYSPACE_PROPERTY.html */
func (i NotesInternational) IsCurrencySpace() (Boolean, error) {
	return getComProperty[Boolean](i, "IsCurrencySpace")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENCYSUFFIX_PROPERTY.html */
func (i NotesInternational) IsCurrencySuffix() (Boolean, error) {
	return getComProperty[Boolean](i, "IsCurrencySuffix")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCURRENCYZERO_PROPERTY.html */
func (i NotesInternational) IsCurrencyZero() (Boolean, error) {
	return getComProperty[Boolean](i, "IsCurrencyZero")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDATEDMY_PROPERTY.html */
func (i NotesInternational) IsDateDMY() (Boolean, error) {
	return getComProperty[Boolean](i, "IsDateDMY")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDATEMDY_PROPERTY.html */
func (i NotesInternational) IsDateMDY() (Boolean, error) {
	return getComProperty[Boolean](i, "IsDateMDY")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDATEYMD_PROPERTY.html */
func (i NotesInternational) IsDateYMD() (Boolean, error) {
	return getComProperty[Boolean](i, "IsDateYMD")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDST_PROPERTY_INTL.html */
func (i NotesInternational) IsDST() (Boolean, error) {
	return getComProperty[Boolean](i, "IsDST")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISTIME24HOUR_PROPERTY.html */
func (i NotesInternational) IsTime24Hour() (Boolean, error) {
	return getComProperty[Boolean](i, "IsTime24Hour")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_INTERNATIONAL_COM.html */
func (i NotesInternational) Parent() (NotesSession, error) {
	return getComObjectProperty(i, newNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PMSTRING_PROPERTY.html */
func (i NotesInternational) PMString() (String, error) {
	return getComProperty[String](i, "PMString")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_THOUSANDSSEP_PROPERTY.html */
func (i NotesInternational) ThousandsSep() (String, error) {
	return getComProperty[String](i, "ThousandsSep")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMESEP_PROPERTY.html */
func (i NotesInternational) TimeSep() (String, error) {
	return getComProperty[String](i, "TimeSep")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONE_PROPERTY_INTL.html */
func (i NotesInternational) TimeZone() (Long, error) {
	return getComProperty[Long](i, "TimeZone")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TODAY_PROPERTY.html */
func (i NotesInternational) Today() (String, error) {
	return getComProperty[String](i, "Today")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TOMORROW_PROPERTY.html */
func (i NotesInternational) Tomorrow() (String, error) {
	return getComProperty[String](i, "Tomorrow")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_YESTERDAY_PROPERTY.html */
func (i NotesInternational) Yesterday() (String, error) {
	return getComProperty[String](i, "Yesterday")
}

/* --------------------------------- Methods ------------------------------------ */
