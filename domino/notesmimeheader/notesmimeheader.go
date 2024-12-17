package notesmimeheader

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesMIMEHeader struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesMIMEHeader {
	return NotesMIMEHeader{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERNAME_PROPERTY_MIMEHEADER.html */
func (m NotesMIMEHeader) HeaderName() (domino.String, error) {
	val, err := m.Com().GetProperty("HeaderName")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDVALTEXT_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) AddValText(valueText domino.String, charSet domino.String) (domino.Boolean, error) {
	val, err := m.Com().CallMethod("AddValText", valueText, charSet)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETHEADERVAL_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) GetHeaderVal(folded domino.Boolean, decoded domino.Boolean) (domino.String, error) {
	val, err := m.Com().CallMethod("GetHeaderVal", folded, decoded)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETHEADERVALANDPARAMS_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) GetHeaderValAndParams(folded domino.Boolean, decoded domino.Boolean) (domino.String, error) {
	val, err := m.Com().CallMethod("GetHeaderValAndParams", folded, decoded)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARAMVAL_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) GetParamVal(paramName domino.String, folded domino.Boolean) (domino.String, error) {
	val, err := m.Com().CallMethod("GetParamVal", paramName, folded)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) Remove() error {
	_, err := m.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHEADERVAL_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) SetHeaderVal(headerValue domino.String) (domino.Boolean, error) {
	val, err := m.Com().CallMethod("SetHeaderVal", headerValue)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHEADERVALANDPARAMS_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) SetHeaderValAndParams(headerParamValue domino.String) (domino.Boolean, error) {
	val, err := m.Com().CallMethod("SetHeaderValAndParams", headerParamValue)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPARAMVAL_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) SetParamVal(parameterName domino.String, parameterValue domino.String) (domino.Boolean, error) {
	val, err := m.Com().CallMethod("SetParamVal", parameterName, parameterValue)
	return helpers.CastValue[domino.Boolean](val), err
}
