/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesProperty struct {
	NotesStruct
}

func NewNotesProperty(dispatchPtr *ole.IDispatch) NotesProperty {
	return NotesProperty{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_DESCRIPTION_PROPERTY.html */
func (p NotesProperty) Description() (String, error) {
	return getComProperty[String](p, "Description")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_ISINPUT_PROPERTY.html */
func (p NotesProperty) IsInput() (Boolean, error) {
	return getComProperty[Boolean](p, "IsInput")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAME_PROPERTY.html */
func (p NotesProperty) Name() (String, error) {
	return getComProperty[String](p, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAMESPACE_PROPERTY.html */
func (p NotesProperty) NameSpace() (String, error) {
	return getComProperty[String](p, "NameSpace")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TITLE_PROPERTY.html */
func (p NotesProperty) Title() (String, error) {
	return getComProperty[String](p, "Title")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TYPENAME_PROPERTY.html */
func (p NotesProperty) Typename() (String, error) {
	return getComProperty[String](p, "Typename")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func (p NotesProperty) Values() ([]Integer, error) {
	return getComArrayProperty[Integer](p, "Values")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func (p NotesProperty) SetValues(v []Integer) error {
	return putComProperty(p, "Values", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLEAR_METHOD.html */
func (p NotesProperty) Clear() error {
	return callComVoidMethod(p, "Clear")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_PUBLISH_METHOD.html */
func (p NotesProperty) Publish() error {
	return callComVoidMethod(p, "Publish")
}
