/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESMIMEHEADER_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesMIMEHeader struct {
	NotesStruct
}

func NewNotesMIMEHeader(dispatchPtr *ole.IDispatch) NotesMIMEHeader {
	return NotesMIMEHeader{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERNAME_PROPERTY_MIMEHEADER.html */
func (m NotesMIMEHeader) HeaderName() (String, error) {
	val, err := getComProperty(m, "HeaderName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_MIMEHEADER.html */
func (m NotesMIMEHeader) Parent() (NotesMIMEEntity, error) {
	dispatchPtr, err := getComObjectProperty(m, "Parent")
	return NewNotesMIMEEntity(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDVALTEXT_METHOD_MIMEHEADER.html */
type notesMIMEHeaderAddValTextParams struct {
	charSet *String
}

type notesMIMEHeaderAddValTextParam func(*notesMIMEHeaderAddValTextParams)

func WithNotesMIMEHeaderAddValTextCharSet(charSet String) notesMIMEHeaderAddValTextParam {
	return func(c *notesMIMEHeaderAddValTextParams) {
		c.charSet = &charSet
	}
}

func (m NotesMIMEHeader) AddValText(valueText String, params ...notesMIMEHeaderAddValTextParam) (Boolean, error) {
	paramsStruct := &notesMIMEHeaderAddValTextParams{}
	paramsOrdered := []interface{}{valueText}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.charSet != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.charSet)
	}
	val, err := callComMethod(m, "AddValText", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETHEADERVAL_METHOD_MIMEHEADER.html */
type notesMIMEHeaderGetHeaderValParams struct {
	folded  *Boolean
	decoded *Boolean
}

type notesMIMEHeaderGetHeaderValParam func(*notesMIMEHeaderGetHeaderValParams)

func WithNotesMIMEHeaderGetHeaderValFolded(folded Boolean) notesMIMEHeaderGetHeaderValParam {
	return func(c *notesMIMEHeaderGetHeaderValParams) {
		c.folded = &folded
	}
}

func WithNotesMIMEHeaderGetHeaderValDecoded(decoded Boolean) notesMIMEHeaderGetHeaderValParam {
	return func(c *notesMIMEHeaderGetHeaderValParams) {
		c.decoded = &decoded
	}
}

func (m NotesMIMEHeader) GetHeaderVal(params ...notesMIMEHeaderGetHeaderValParam) (String, error) {
	paramsStruct := &notesMIMEHeaderGetHeaderValParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.folded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.folded)
		if paramsStruct.decoded != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
		}
	}
	val, err := callComMethod(m, "GetHeaderVal", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETHEADERVALANDPARAMS_METHOD_MIMEHEADER.html */
type notesMIMEHeaderGetHeaderValAndParamsParams struct {
	folded  *Boolean
	decoded *Boolean
}

type notesMIMEHeaderGetHeaderValAndParamsParam func(*notesMIMEHeaderGetHeaderValAndParamsParams)

func WithNotesMIMEHeaderGetHeaderValAndParamsFolded(folded Boolean) notesMIMEHeaderGetHeaderValAndParamsParam {
	return func(c *notesMIMEHeaderGetHeaderValAndParamsParams) {
		c.folded = &folded
	}
}

func WithNotesMIMEHeaderGetHeaderValAndParamsDecoded(decoded Boolean) notesMIMEHeaderGetHeaderValAndParamsParam {
	return func(c *notesMIMEHeaderGetHeaderValAndParamsParams) {
		c.decoded = &decoded
	}
}

func (m NotesMIMEHeader) GetHeaderValAndParams(params ...notesMIMEHeaderGetHeaderValAndParamsParam) (String, error) {
	paramsStruct := &notesMIMEHeaderGetHeaderValAndParamsParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.folded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.folded)
		if paramsStruct.decoded != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
		}
	}
	val, err := callComMethod(m, "GetHeaderValAndParams", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARAMVAL_METHOD_MIMEHEADER.html */
type notesMIMEHeaderGetParamValParams struct {
	folded *Boolean
}

type notesMIMEHeaderGetParamValParam func(*notesMIMEHeaderGetParamValParams)

func WithNotesMIMEHeaderGetParamValFolded(folded Boolean) notesMIMEHeaderGetParamValParam {
	return func(c *notesMIMEHeaderGetParamValParams) {
		c.folded = &folded
	}
}

func (m NotesMIMEHeader) GetParamVal(paramName String, params ...notesMIMEHeaderGetParamValParam) (String, error) {
	paramsStruct := &notesMIMEHeaderGetParamValParams{}
	paramsOrdered := []interface{}{paramName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.folded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.folded)
	}
	val, err := callComMethod(m, "GetParamVal", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) Remove() error {
	_, err := callComMethod(m, "Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHEADERVAL_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) SetHeaderVal(headerValue String) (Boolean, error) {
	val, err := callComMethod(m, "SetHeaderVal", headerValue)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHEADERVALANDPARAMS_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) SetHeaderValAndParams(headerParamValue String) (Boolean, error) {
	val, err := callComMethod(m, "SetHeaderValAndParams", headerParamValue)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPARAMVAL_METHOD_MIMEHEADER.html */
func (m NotesMIMEHeader) SetParamVal(parameterName String, parameterValue String) (Boolean, error) {
	val, err := callComMethod(m, "SetParamVal", parameterName, parameterValue)
	return helpers.CastValue[Boolean](val), err
}
