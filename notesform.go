/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFORM_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesForm struct {
	NotesStruct
}

func newNotesForm(dispatchPtr *ole.IDispatch) NotesForm {
	return NotesForm{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIASES_PROPERTY_FORM.html */
func (f NotesForm) Aliases() ([]String, error) {
	return getComArrayProperty[String](f, "Aliases")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIELDS_PROPERTY.html */
func (f NotesForm) Fields() ([]String, error) {
	return getComArrayProperty[String](f, "Fields")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMUSERS_PROPERTY.html */
func (f NotesForm) FormUsers() ([]String, error) {
	return getComArrayProperty[String](f, "FormUsers")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMUSERS_PROPERTY.html */
func (f NotesForm) SetFormUsers(v []String) error {
	return putComProperty(f, "FormUsers", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_FORM_COM.html */
func (f NotesForm) HttpURL() (String, error) {
	return getComProperty[String](f, "HttpURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUBFORM_PROPERTY.html */
func (f NotesForm) IsSubForm() (Boolean, error) {
	return getComProperty[Boolean](f, "IsSubForm")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_FORM.html */
func (f NotesForm) LockHolders() ([]String, error) {
	return getComArrayProperty[String](f, "LockHolders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_FORM.html */
func (f NotesForm) Name() (String, error) {
	return getComProperty[String](f, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_FORM_COM.html */
func (f NotesForm) NotesURL() (String, error) {
	return getComProperty[String](f, "NotesURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_FORM_COM.html */
func (f NotesForm) Parent() (NotesDatabase, error) {
	return getComObjectProperty(f, newNotesDatabase, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY.html */
func (f NotesForm) ProtectReaders() (Boolean, error) {
	return getComProperty[Boolean](f, "ProtectReaders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTREADERS_PROPERTY.html */
func (f NotesForm) SetProtectReaders(v Boolean) error {
	return putComProperty(f, "ProtectReaders", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTUSERS_PROPERTY.html */
func (f NotesForm) ProtectUsers() (Boolean, error) {
	return getComProperty[Boolean](f, "ProtectUsers")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROTECTUSERS_PROPERTY.html */
func (f NotesForm) SetProtectUsers(v Boolean) error {
	return putComProperty(f, "ProtectUsers", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY.html */
func (f NotesForm) Readers() ([]String, error) {
	return getComArrayProperty[String](f, "Readers")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READERS_PROPERTY.html */
func (f NotesForm) SetReaders(v []String) error {
	return putComProperty(f, "Readers", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIELDTYPE_METHOD_FORM.html */
func (f NotesForm) GetFieldType(name String) (NotesItemDataType, error) {
	return callComMethod[NotesItemDataType](f, "GetFieldType", name)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_FORM.html */
type notesFormLockParams struct {
	name          *[]String
	provisionalOK *Boolean
}

type notesFormLockParam func(*notesFormLockParams)

func WithNotesFormLockName(name []String) notesFormLockParam {
	return func(c *notesFormLockParams) {
		c.name = &name
	}
}

func WithNotesFormLockProvisionalOK(provisionalOK Boolean) notesFormLockParam {
	return func(c *notesFormLockParams) {
		c.provisionalOK = &provisionalOK
	}
}

func (f NotesForm) Lock(params ...notesFormLockParam) (Boolean, error) {
	paramsStruct := &notesFormLockParams{}
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
	return callComMethod[Boolean](f, "Lock", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_FORM.html */
type notesFormLockProvisionalParams struct {
	name *[]String
}

type notesFormLockProvisionalParam func(*notesFormLockProvisionalParams)

func WithNotesFormLockProvisionalName(name []String) notesFormLockProvisionalParam {
	return func(c *notesFormLockProvisionalParams) {
		c.name = &name
	}
}

func (f NotesForm) LockProvisional(params ...notesFormLockProvisionalParam) (Boolean, error) {
	paramsStruct := &notesFormLockProvisionalParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	return callComMethod[Boolean](f, "LockProvisional", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_FORM.html */
func (f NotesForm) Remove() error {
	return callComVoidMethod(f, "Remove")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_FORM.html */
func (f NotesForm) UnLock() error {
	return callComVoidMethod(f, "UnLock")
}
