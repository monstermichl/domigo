/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATERANGE_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

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
	dispatchPtr, err := d.com().GetObjectProperty("EndDateTime")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENDDATETIME_PROPERTY.html */
func (d NotesDateRange) SetEndDateTime(v NotesDateTime) error {
	return d.com().PutProperty("EndDateTime", v.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DATERANGE_COM.html */
func (d NotesDateRange) Parent() (NotesSession, error) {
	dispatchPtr, err := d.com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func (d NotesDateRange) StartDateTime() (NotesDateTime, error) {
	dispatchPtr, err := d.com().GetObjectProperty("StartDateTime")
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func (d NotesDateRange) SetStartDateTime(v NotesDateTime) error {
	return d.com().PutProperty("StartDateTime", v.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func (d NotesDateRange) Text() (String, error) {
	val, err := d.com().GetProperty("Text")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func (d NotesDateRange) SetText(v String) error {
	return d.com().PutProperty("Text", v)
}

/* --------------------------------- Methods ------------------------------------ */
