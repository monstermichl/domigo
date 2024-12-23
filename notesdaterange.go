/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATERANGE_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesDateRange struct {
	NotesStruct
}

func NewNotesDateRange(dispatchPtr *ole.IDispatch) NotesDateRange {
	return NotesDateRange{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENDDATETIME_PROPERTY.html */
func (d NotesDateRange) EndDateTime() (NotesDateTime, error) {
	return getComObjectProperty(d, NewNotesDateTime, "EndDateTime")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENDDATETIME_PROPERTY.html */
func (d NotesDateRange) SetEndDateTime(v NotesDateTime) error {
	return putComProperty(d, "EndDateTime", v.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DATERANGE_COM.html */
func (d NotesDateRange) Parent() (NotesSession, error) {
	return getComObjectProperty(d, NewNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func (d NotesDateRange) StartDateTime() (NotesDateTime, error) {
	return getComObjectProperty(d, NewNotesDateTime, "StartDateTime")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func (d NotesDateRange) SetStartDateTime(v NotesDateTime) error {
	return putComProperty(d, "StartDateTime", v.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func (d NotesDateRange) Text() (String, error) {
	return getComProperty[String](d, "Text")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func (d NotesDateRange) SetText(v String) error {
	return putComProperty(d, "Text", v)
}

/* --------------------------------- Methods ------------------------------------ */
