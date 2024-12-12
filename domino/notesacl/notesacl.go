/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACL_CLASS.html */
package notesacl

import (
	"domigo/domino"
	"domigo/domino/notesaclentry"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesACL struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesACL {
	return NotesACL{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_ACLENTRY.html */
/* Moved from NotesACLEntry. */
func NotesACLEntryParent(a notesaclentry.NotesACLEntry) (NotesACL, error) {
	dispatchPtr, err := a.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func (a NotesACL) AdministrationServer() (domino.String, error) {
	val, err := a.Com().GetProperty("AdministrationServer")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMINISTRATIONSERVER_PROPERTY_ACL.html */
func (a NotesACL) SetAdministrationServer(v domino.String) error {
	return a.Com().PutProperty("AdministrationServer", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) InternetLevel() (domino.Long, error) {
	val, err := a.Com().GetProperty("InternetLevel")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNETLEVEL_PROPERTY_1219.html */
func (a NotesACL) SetInternetLevel(v domino.Long) error {
	return a.Com().PutProperty("InternetLevel", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) IsAdminNames() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsAdminNames")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINNAMES_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminNames(v domino.Boolean) error {
	return a.Com().PutProperty("IsAdminNames", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) IsAdminReaderAuthor() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsAdminReaderAuthor")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISADMINREADERAUTHOR_PROPERTY_ACL.html */
func (a NotesACL) SetIsAdminReaderAuthor(v domino.Boolean) error {
	return a.Com().PutProperty("IsAdminReaderAuthor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) IsExtendedAccess() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsExtendedAccess")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEXTENDEDACCESS_PROPERTY_ACL.html */
func (a NotesACL) SetIsExtendedAccess(v domino.Boolean) error {
	return a.Com().PutProperty("IsExtendedAccess", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ROLES_PROPERTY_ACL.html */
func (a NotesACL) Roles() ([]domino.String, error) {
	vals, err := a.Com().GetArrayProperty("Roles")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) UniformAccess() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("UniformAccess")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIFORMACCESS_PROPERTY.html */
func (a NotesACL) SetUniformAccess(v domino.Boolean) error {
	return a.Com().PutProperty("UniformAccess", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDROLE_METHOD.html */
func (a NotesACL) AddRole(name domino.String) error {
	_, err := a.Com().CallMethod("AddRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEACLENTRY_METHOD.html */
func (a NotesACL) CreateACLEntry(name domino.String, level domino.Long) (notesaclentry.NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("CreateACLEntry", name, level)
	return notesaclentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEROLE_METHOD.html */
func (a NotesACL) DeleteRole(name domino.String) error {
	_, err := a.Com().CallMethod("DeleteRole", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD.html */
func (a NotesACL) GetEntry(name domino.String) (notesaclentry.NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("GetEntry", name)
	return notesaclentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD.html */
func (a NotesACL) GetFirstEntry() (notesaclentry.NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("GetFirstEntry")
	return notesaclentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_METHOD.html */
func (a NotesACL) GetNextEntry(entry notesaclentry.NotesACLEntry) (notesaclentry.NotesACLEntry, error) {
	dispatchPtr, err := a.Com().CallObjectMethod("GetNextEntry", entry.Com().Dispatch())
	return notesaclentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEACLENTRY_METHOD_ACL_COM.html */
func (a NotesACL) RemoveACLEntry(name domino.String) error {
	_, err := a.Com().CallMethod("RemoveACLEntry", name)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RENAMEROLE_METHOD.html */
func (a NotesACL) RenameRole(oldName domino.String, newName domino.String) error {
	_, err := a.Com().CallMethod("RenameRole", oldName, newName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_ACL.html */
func (a NotesACL) Save() error {
	_, err := a.Com().CallMethod("Save")
	return err
}
