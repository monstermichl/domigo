/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

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
	val, err := p.com().GetProperty("Description")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_ISINPUT_PROPERTY.html */
func (p NotesProperty) IsInput() (Boolean, error) {
	val, err := p.com().GetProperty("IsInput")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAME_PROPERTY.html */
func (p NotesProperty) Name() (String, error) {
	val, err := p.com().GetProperty("Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAMESPACE_PROPERTY.html */
func (p NotesProperty) NameSpace() (String, error) {
	val, err := p.com().GetProperty("NameSpace")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TITLE_PROPERTY.html */
func (p NotesProperty) Title() (String, error) {
	val, err := p.com().GetProperty("Title")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TYPENAME_PROPERTY.html */
func (p NotesProperty) Typename() (String, error) {
	val, err := p.com().GetProperty("Typename")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func (p NotesProperty) Values() ([]Integer, error) {
	vals, err := p.com().GetArrayProperty("Values")
	return helpers.CastSlice[Integer](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func (p NotesProperty) SetValues(v []Integer) error {
	return p.com().PutProperty("Values", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLEAR_METHOD.html */
func (p NotesProperty) Clear() error {
	_, err := p.com().CallMethod("Clear")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_PUBLISH_METHOD.html */
func (p NotesProperty) Publish() error {
	_, err := p.com().CallMethod("Publish")
	return err
}
