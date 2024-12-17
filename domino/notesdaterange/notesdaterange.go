/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATERANGE_CLASS.html */
package notesdaterange

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdatetime"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDateRange struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDateRange {
	return NotesDateRange{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENDDATETIME_PROPERTY.html */
func (d NotesDateRange) EndDateTime() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("EndDateTime")
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENDDATETIME_PROPERTY.html */
func (d NotesDateRange) SetEndDateTime(v notesdatetime.NotesDateTime) error {
	return d.Com().PutProperty("EndDateTime", v.Com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func (d NotesDateRange) StartDateTime() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("StartDateTime")
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STARTDATETIME_PROPERTY.html */
func (d NotesDateRange) SetStartDateTime(v notesdatetime.NotesDateTime) error {
	return d.Com().PutProperty("StartDateTime", v.Com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func (d NotesDateRange) Text() (domino.String, error) {
	val, err := d.Com().GetProperty("Text")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY_RANG.html */
func (d NotesDateRange) SetText(v domino.String) error {
	return d.Com().PutProperty("Text", v)
}

/* --------------------------------- Methods ------------------------------------ */
