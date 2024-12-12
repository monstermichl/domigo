/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTYBROKER_CLASS.html */
package notespropertybroker

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesproperty"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesPropertyBroker struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesPropertyBroker {
	return NotesPropertyBroker{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_INPUTPROPERTYCONTEXT.html */
func (p NotesPropertyBroker) InputPropertyContext() ([]notesproperty.NotesProperty, error) {
	return com.GetObjectArrayProperty(p.Com(), notesproperty.New, "InputPropertyContext")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_CLEARPROPERTY_METHOD.html */
func (p NotesPropertyBroker) ClearProperty(name domino.String) error {
	_, err := p.Com().CallMethod("ClearProperty", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTY_METHOD.html */
func (p NotesPropertyBroker) GetProperty(name domino.String) (notesproperty.NotesProperty, error) {
	dispatchPtr, err := p.Com().CallObjectMethod("GetProperty", name)
	return notesproperty.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_GETPROPERTYVALUE_METHOD.html */
func (p NotesPropertyBroker) GetPropertyValue(name domino.String) (any, error) {
	return p.Com().CallMethod("GetPropertyValue", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_HASPROPERTY_METHOD.html */
func (p NotesPropertyBroker) HasProperty(Name domino.String) (domino.Boolean, error) {
	val, err := p.Com().CallMethod("HasProperty", Name)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_PUBLISH_METHOD.html */
func (p NotesPropertyBroker) Publish() error {
	_, err := p.Com().CallMethod("Publish")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROPERTYBROKER_SETPROPERTYVALUE_METHOD.html */
func (p NotesPropertyBroker) SetPropertyValue(name domino.String, value domino.String) (notesproperty.NotesProperty, error) {
	dispatchPtr, err := p.Com().CallObjectMethod("SetPropertyValue", name, value)
	return notesproperty.New(dispatchPtr), err
}
