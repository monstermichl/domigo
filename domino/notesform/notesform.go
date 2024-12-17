/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFORM_CLASS.html */
package notesform

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesForm struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesForm {
	return NotesForm{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_FORM.html */
func (f NotesForm) Aliases() ([]domino.String, error) {
	vals, err := f.Com().GetArrayProperty("Aliases")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIELDS_PROPERTY.html */
func (f NotesForm) Fields() ([]domino.String, error) {
	vals, err := f.Com().GetArrayProperty("Fields")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMUSERS_PROPERTY.html */
func (f NotesForm) FormUsers() ([]domino.String, error) {
	vals, err := f.Com().GetArrayProperty("FormUsers")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMUSERS_PROPERTY.html */
func (f NotesForm) SetFormUsers(v []domino.String) error {
	return f.Com().PutProperty("FormUsers", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_FORM_COM.html */
func (f NotesForm) HttpURL() (domino.String, error) {
	val, err := f.Com().GetProperty("HttpURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUBFORM_PROPERTY.html */
func (f NotesForm) IsSubForm() (domino.Boolean, error) {
	val, err := f.Com().GetProperty("IsSubForm")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_FORM.html */
func (f NotesForm) LockHolders() ([]domino.String, error) {
	vals, err := f.Com().GetArrayProperty("LockHolders")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_FORM.html */
func (f NotesForm) Name() (domino.String, error) {
	val, err := f.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_FORM_COM.html */
func (f NotesForm) NotesURL() (domino.String, error) {
	val, err := f.Com().GetProperty("NotesURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY.html */
func (f NotesForm) ProtectReaders() (domino.Boolean, error) {
	val, err := f.Com().GetProperty("ProtectReaders")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY.html */
func (f NotesForm) SetProtectReaders(v domino.Boolean) error {
	return f.Com().PutProperty("ProtectReaders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTUSERS_PROPERTY.html */
func (f NotesForm) ProtectUsers() (domino.Boolean, error) {
	val, err := f.Com().GetProperty("ProtectUsers")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTUSERS_PROPERTY.html */
func (f NotesForm) SetProtectUsers(v domino.Boolean) error {
	return f.Com().PutProperty("ProtectUsers", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY.html */
func (f NotesForm) Readers() ([]domino.String, error) {
	vals, err := f.Com().GetArrayProperty("Readers")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY.html */
func (f NotesForm) SetReaders(v []domino.String) error {
	return f.Com().PutProperty("Readers", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIELDTYPE_METHOD_FORM.html */
func (f NotesForm) GetFieldType(name domino.String) (domino.Long, error) {
	val, err := f.Com().CallMethod("GetFieldType", name)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_FORM.html */
type lockParams struct {
	name          *[]domino.String
	provisionalOK *domino.Boolean
}

type lockParam func(*lockParams)

func WithLockName(name []domino.String) lockParam {
	return func(c *lockParams) {
		c.name = &name
	}
}

func WithLockProvisionalOK(provisionalOK domino.Boolean) lockParam {
	return func(c *lockParams) {
		c.provisionalOK = &provisionalOK
	}
}

func (f NotesForm) Lock(params ...lockParam) (domino.Boolean, error) {
	paramsStruct := &lockParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
		if paramsStruct.provisionalOK != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.provisionalOK)
		}
	}
	val, err := f.Com().CallMethod("Lock", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_FORM.html */
type lockProvisionalParams struct {
	name *[]domino.String
}

type lockProvisionalParam func(*lockProvisionalParams)

func WithLockProvisionalName(name []domino.String) lockProvisionalParam {
	return func(c *lockProvisionalParams) {
		c.name = &name
	}
}

func (f NotesForm) LockProvisional(params ...lockProvisionalParam) (domino.Boolean, error) {
	paramsStruct := &lockProvisionalParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	val, err := f.Com().CallMethod("LockProvisional", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_FORM.html */
func (f NotesForm) Remove() error {
	_, err := f.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_FORM.html */
func (f NotesForm) UnLock() error {
	_, err := f.Com().CallMethod("UnLock")
	return err
}
