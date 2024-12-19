/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTYBROKER_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesPropertyBroker struct {
	NotesStruct
}

func NewNotesPropertyBroker(dispatchPtr *ole.IDispatch) NotesPropertyBroker {
	return NotesPropertyBroker{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_INPUTPROPERTYCONTEXT.html */
func (p NotesPropertyBroker) InputPropertyContext() ([]NotesProperty, error) {
	return com.GetObjectArrayProperty(p.com(), NewNotesProperty, "InputPropertyContext")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_CLEARPROPERTY_METHOD.html */
func (p NotesPropertyBroker) ClearProperty(name String) error {
	_, err := p.com().CallMethod("ClearProperty", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTY_METHOD.html */
func (p NotesPropertyBroker) GetProperty(name String) (NotesProperty, error) {
	dispatchPtr, err := p.com().CallObjectMethod("GetProperty", name)
	return NewNotesProperty(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTYVALUE_METHOD.html */
func (p NotesPropertyBroker) GetPropertyValue(name String) (any, error) {
	return p.com().CallMethod("GetPropertyValue", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_HASPROPERTY_METHOD.html */
func (p NotesPropertyBroker) HasProperty(Name String) (Boolean, error) {
	val, err := p.com().CallMethod("HasProperty", Name)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_PUBLISH_METHOD.html */
func (p NotesPropertyBroker) Publish() error {
	_, err := p.com().CallMethod("Publish")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_SETPROPERTYVALUE_METHOD.html */
func (p NotesPropertyBroker) SetPropertyValue(name String, value String) (NotesProperty, error) {
	dispatchPtr, err := p.com().CallObjectMethod("SetPropertyValue", name, value)
	return NewNotesProperty(dispatchPtr), err
}
