/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLASS.html */
package notesproperty

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesProperty struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesProperty {
	return NotesProperty{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_DESCRIPTION_PROPERTY.html */
func (p NotesProperty) Description() (domino.String, error) {
	val, err := p.Com().GetProperty("Description")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_ISINPUT_PROPERTY.html */
func (p NotesProperty) IsInput() (domino.Boolean, error) {
	val, err := p.Com().GetProperty("IsInput")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAME_PROPERTY.html */
func (p NotesProperty) Name() (domino.String, error) {
	val, err := p.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_NAMESPACE_PROPERTY.html */
func (p NotesProperty) NameSpace() (domino.String, error) {
	val, err := p.Com().GetProperty("NameSpace")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TITLE_PROPERTY.html */
func (p NotesProperty) Title() (domino.String, error) {
	val, err := p.Com().GetProperty("Title")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_TYPENAME_PROPERTY.html */
func (p NotesProperty) Typename() (domino.String, error) {
	val, err := p.Com().GetProperty("Typename")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func (p NotesProperty) Values() ([]domino.Integer, error) {
	vals, err := p.Com().GetArrayProperty("Values")
	return helpers.CastSlice[domino.Integer](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_VALUES_PROPERTY.html */
func (p NotesProperty) SetValues(v []domino.Integer) error {
	return p.Com().PutProperty("Values", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_CLEAR_METHOD.html */
func (p NotesProperty) Clear() error {
	_, err := p.Com().CallMethod("Clear")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESPROPERTY_PUBLISH_METHOD.html */
func (p NotesProperty) Publish() error {
	_, err := p.Com().CallMethod("Publish")
	return err
}
