/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTYBROKER_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"

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
	return callComVoidMethod(p, "ClearProperty", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTY_METHOD.html */
func (p NotesPropertyBroker) GetProperty(name String) (NotesProperty, error) {
	return callComObjectMethod(p, NewNotesProperty, "GetProperty", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTYVALUE_METHOD.html */
func (p NotesPropertyBroker) GetPropertyValue(name String) (any, error) {
	return callComMethod(p, "GetPropertyValue", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_HASPROPERTY_METHOD.html */
func (p NotesPropertyBroker) HasProperty(Name String) (Boolean, error) {
	return callComMethod[Boolean](p, "HasProperty", Name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_PUBLISH_METHOD.html */
func (p NotesPropertyBroker) Publish() error {
	return callComVoidMethod(p, "Publish")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_SETPROPERTYVALUE_METHOD.html */
func (p NotesPropertyBroker) SetPropertyValue(name String, value String) (NotesProperty, error) {
	return callComObjectMethod(p, NewNotesProperty, "SetPropertyValue", name, value)
}
